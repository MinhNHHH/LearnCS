package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter interface {
	Allow() bool
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
func (fw *FixedWindow) Allow() bool {
	fw.mux.Lock()
	defer fw.mux.Unlock()

	now := time.Now()

	// If the current time is past the window size, reset and create new window.
	if now.After(fw.windowEnd) {
		fw.counter = 0
		fw.windowEnd = now.Add(fw.windowStart)
	}

	// Check if the reqeust is allowed within the current window
	if fw.counter < fw.limit {
		fw.counter++
		return true
	}

	// Deny if the limi is reached.
	return false
}

type SlidingWindow struct {
	counter int
	limit   int
}

type LeakingBucket struct {
	bucket     []int
	capacity   int
	rate       int
	lastFilled time.Time
	mux        sync.Mutex
}

func NewRateLimiterLeakingBucket(rate, bucketSize int) *LeakingBucket {
	return &LeakingBucket{
		rate:     rate,
		capacity: bucketSize,
		bucket:   make([]int, bucketSize),
	}
}

// It is usually implemented with FIFO queue.
// Leaking bucket algorithm takes the following two parameters:
// Bucket size: It's equal to the queue size. The queue holds the requests to be processed at a fixed rate
// Outflow rate: It defines how many request can be processed at a fixed rate.
func (rl *LeakingBucket) Allow() bool {
	rl.mux.Lock()
	defer rl.mux.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastFilled).Seconds()
	leaked := int(elapsed * float64(rl.rate))

	// Remove old requests
	if leaked > 0 {
		if leaked >= len(rl.bucket) {
			rl.bucket = rl.bucket[:0]
		} else {
			leaked = 0
			rl.bucket = rl.bucket[leaked:]
		}
	}
	if len(rl.bucket) < rl.capacity {
		rl.bucket = append(rl.bucket, 1)
		rl.lastFilled = now
		return true
	}
	return false
}

type TokenBucket struct {
	rate       int       // tokens added per second
	capacity   int       // max tokens in the bucket
	tokens     int       // current number of tokens
	lastFilled time.Time // last time the bucket was filled
	mux        sync.Mutex
}

func NewTokenBucket(rate, capacity int) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastFilled: time.Now(),
	}
}

// The token bucket algorithm work as follows:
// A token bucket is a contain that has re-defined capacity. Token are put in the bucket at preset rate
// Once the bucket is full, no more tokens are added.
// Each request consumes on token. When a request arrives, we check if there are enoguh tokens in the bucket
// + If there are enoguh tokens, we take on token out for each request, and the request goes through
// + If there are not enoguh tokens, the request is dropped.
func (rl *TokenBucket) Allow() bool {
	rl.mux.Lock()
	defer rl.mux.Unlock()
	now := time.Now()
	elapsed := now.Sub(rl.lastFilled).Seconds()
	rl.tokens += int(elapsed * float64(rl.rate))
	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}
	rl.lastFilled = now

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

func RateLimiterMiddleWare(limiter RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			c.JSON(429, gin.H{
				"message": "Too many requests",
			})
			c.Abort()
		}
	}
}
