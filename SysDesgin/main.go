package main

import (
	"sync"
	"time"

	mdw "github.com/MinhNHHH/sys-design/middleware"
	"github.com/gin-gonic/gin"
)

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
	limiter := mdw.NewTokenBucket(5, 10) // 5 tokens per second, max 10 tokens

	r := gin.Default()
	r.Use(mdw.RateLimiterMiddleWare(limiter))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
