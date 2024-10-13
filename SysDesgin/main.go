package main

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiterTokenBucket struct {
	rate       int       // tokens added per second
	capacity   int       // max tokens in the bucket
	tokens     int       // current number of tokens
	lastFilled time.Time // last time the bucket was filled
	mux        sync.Mutex
}

func NewRateLimiterTokenBuck(rate, capacity int) *RateLimiterTokenBucket {
	return &RateLimiterTokenBucket{
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
func (rl *RateLimiterTokenBucket) Allow() bool {
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

func RateLimiterMiddleWare(limiter *RateLimiterTokenBucket) gin.HandlerFunc {
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

func RateLimiterMiddleWareLeaking(limiter *LeakingBucket) gin.HandlerFunc {
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

func main() {
	limiter := NewRateLimiterLeakingBucket(5, 10) // 5 tokens per second, max 10 tokens
	r := gin.Default()
	r.Use(RateLimiterMiddleWareLeaking(limiter))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
