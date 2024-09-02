package middleware

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	mux sync.Mutex
	rdb *redis.Client
}

var capacity = 10

func (rl *RateLimiter) RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rl.mux.Lock()
		defer rl.mux.Unlock()

		val, err := rl.rdb.Get(ctx, ctx.ClientIP()).Result()
		if err != nil {
			val = "0"
			rl.rdb.Set(ctx, ctx.ClientIP(), val, 0)
		}

		counter, errAtoi := strconv.Atoi(val)
		if errAtoi != nil {
			panic(errAtoi)
		}

		if counter+1 > capacity {
			ctx.JSON(429, gin.H{
				"message": "Too many requests",
			})
			ctx.Abort()
		} else {
			counter++
			rl.rdb.Set(ctx, ctx.ClientIP(), counter, 0)
			ctx.Next()
		}
	}
}

func InitLimiter(rdb *redis.Client) *RateLimiter {
	return &RateLimiter{
		rdb: rdb,
	}
}
