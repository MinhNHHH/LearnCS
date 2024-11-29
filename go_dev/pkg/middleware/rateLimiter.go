package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter interface {
	Allow() bool
}

type TokenBucket struct {
	rate       int
	capacity   int
	tokens     int
	lastFilled time.Time
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
	return func(ctx *gin.Context) {
		if limiter.Allow() {
			ctx.Next()
		} else {
			ctx.JSON(429, gin.H{
				"message": "Too many requests",
			})
		}
	}
}
