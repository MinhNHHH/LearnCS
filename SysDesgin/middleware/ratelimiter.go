package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

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

func RateLimiterMiddleWare(limiter *TokenBucket) gin.HandlerFunc {
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
