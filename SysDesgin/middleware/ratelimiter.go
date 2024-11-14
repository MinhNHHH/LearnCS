package middleware

import (
	"log"
	"strconv"
	"sync"
	"time"

	rdb "github.com/MinhNHHH/sys-design/redis"
	"github.com/gin-gonic/gin"
)

type RateLimiter interface {
	Allow(c *gin.Context) bool
}

type FixedWindow struct {
	limit       int           // limit request
	windowEnd   time.Time     // end time of the current window
	windowStart time.Duration // size of the window (eg: 1 minute)
	counter     int           // number of request in the current window
	mux         sync.Mutex
}

func NewRateLimiterFixedWindow(limit int, window time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:       limit,
		windowStart: window,
		windowEnd:   time.Now().Add(window),
		counter:     0,
	}
}

// This limitation of the Fixed Window Counter comes from its inability to handle burst traffic
// Burstiness at Window Boundaries
// The Fixed window algorithm divides time into strict intervals. A user can make the maximum allowed number of request at the end od one window and immediately make more request at the start of the next window.
// Ex: A user could make 5 requests at 12:00:59 and 5 more at 12:01:00, allowing 10 requests in a span of 2 seconds.

// The rigid time windows problem comes from the user is penalized havily for making a request right before the window resets.
// The user makes 4 requests at 12:00:01, 12:00:10, 12:00:20, and 12:00:30. So far the counter for the window 12:00:00 to 12:00:59 is 4/5 requests made out of 5 allowed.
// The user makes a 5th request at 12:00:55. Fis fills up the request limit for current window. Now the counter is 5/5, and no more requests allowed for this window until it resets.
func (fw *FixedWindow) Allow(c *gin.Context) bool {
	fw.mux.Lock()
	defer fw.mux.Unlock()
	now := time.Now()
	if now.After(fw.windowEnd) {
		fw.counter = 0
		fw.windowEnd = now.Add(fw.windowStart)
	}
	if fw.counter < fw.limit {
		fw.counter++
		return true
	}
	return false
}

type LeakingBucket struct {
	capacity int
	rate     int
	rdb      rdb.Cache
	mux      sync.Mutex
}

func NewRateLimiterLeakingBucket(rate, bucketSize int, rdb rdb.Cache) *LeakingBucket {
	return &LeakingBucket{
		rate:     rate,
		capacity: bucketSize,
		rdb:      rdb,
	}
}

// It is usually implemented with FIFO queue.
// Leaking bucket algorithm takes the following two parameters:
// Bucket size: It's equal to the queue size. The queue holds the requests to be processed at a fixed rate
// Outflow rate: It defines how many request can be processed at a fixed rate.
func (rl *LeakingBucket) Allow(ctx *gin.Context) bool {
	rl.mux.Lock()
	defer rl.mux.Unlock()

	now := time.Now()
	rdb := rl.rdb.Client

	lastFilledCmd, err := rdb.Get(ctx, "lastFilled").Result()
	if err != nil {
		log.Println("Error fetching lastFilled:", err)
		return false
	}

	lastFilled, err := ParseTime(lastFilledCmd)
	if err != nil {
		log.Println("Error parsing lastFilled time:", err)
		return false
	}

	// Calculate elapsed time and number of leaked tokens
	elapsed := now.Sub(lastFilled).Seconds()
	leaked := int(elapsed * float64(rl.rate))

	// Get current length of the bucket
	lenBucket, err := rdb.LLen(ctx, "buckets").Result()
	if err != nil {
		log.Println("Error fetching bucket length:", err)
		return false
	}

	// Leak tokens from the bucket
	if leaked > 0 {
		if leaked >= int(lenBucket) {
			// Remove all tokens if leaked tokens exceed or match bucket size
			rdb.LTrim(ctx, "buckets", 1, 0)
			leaked = int(lenBucket)
		} else {
			// Trim only the leaked number of tokens
			rdb.LTrim(ctx, "buckets", int64(leaked), -1)
		}
	}

	// Check if bucket has capacity for a new request
	if int(lenBucket)-leaked < rl.capacity {
		// Add a new token to the bucket
		err = rdb.LPush(ctx, "buckets", 1).Err()
		if err != nil {
			log.Println("Error pushing to bucket:", err)
			return false
		}
		// Update lastFilled timestamp
		err = rdb.Set(ctx, "lastFilled", now, 0).Err()
		if err != nil {
			log.Println("Error setting lastFilled:", err)
			return false
		}
		return true
	}
	// If bucket is full, reject the request
	return false
}

type TokenBucket struct {
	rate     int // tokens added per second
	capacity int // max tokens in the bucket
	mux      sync.Mutex
	rdb      rdb.Cache
}

func NewTokenBucket(rate, capacity int, rdb rdb.Cache) *TokenBucket {

	return &TokenBucket{
		rate:     rate,
		capacity: capacity,
		// tokens:     capacity,
		// lastFilled: time.Now(),
		rdb: rdb,
	}
}

func ParseTime(strTime string) (time.Time, error) {
	// Layout matching the format of your date string
	layout := "2006-01-02T15:04:05.999999-07:00"
	timeParse, error := time.Parse(layout, strTime)
	if error != nil {
		return time.Time{}, error
	}
	return timeParse, nil
}

// The token bucket algorithm work as follows:
// A token bucket is a contain that has re-defined capacity. Token are put in the bucket at preset rate
// Once the bucket is full, no more tokens are added.
// Each request consumes on token. When a request arrives, we check if there are enoguh tokens in the bucket
// + If there are enoguh tokens, we take on token out for each request, and the request goes through
// + If there are not enoguh tokens, the request is dropped.
func (rl *TokenBucket) Allow(ctx *gin.Context) bool {
	rl.mux.Lock()
	defer rl.mux.Unlock()
	rdb := rl.rdb.Client
	now := time.Now()

	lastFilledCmd, err := rdb.Get(ctx, "lastFilled").Result()
	if err != nil {
		log.Fatal("Get lastFilled error:", err)
	}

	lastFilled, err := ParseTime(lastFilledCmd)
	if err != nil {
		log.Fatal("Error parse time", err)
	}

	elapsed := now.Sub(lastFilled).Seconds()
	rdb.IncrBy(ctx, "tokens", int64(elapsed*float64(rl.rate)))
	tokenCmd, err := rdb.Get(ctx, "tokens").Result()

	if err != nil {
		log.Fatal("Get token error:", err)
	}

	token, _ := strconv.Atoi(tokenCmd)
	if token > rl.capacity {
		rdb.Set(ctx, "tokens", rl.capacity, 0)
	}
	rdb.Set(ctx, "lastFilled", now, 0)

	if token > 0 {
		rdb.Decr(ctx, "tokens")
		return true
	}
	return false
}

func RateLimiterMiddleWare(limiter RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow(c) {
			c.Next()
		} else {
			c.JSON(429, gin.H{
				"message": "Too many requests",
			})
			c.Abort()
		}
	}
}
