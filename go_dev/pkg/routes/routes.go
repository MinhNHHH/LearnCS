package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/MinhNHHH/go_dev/pkg/middleware"
)

func Routes() *gin.Engine {
	r := gin.Default()
	limiter := middleware.NewTokenBucket(10, 100)

	r.Use(middleware.EnableCors())
	r.Use(middleware.RateLimiterMiddleWare(limiter))

	r.GET("/heath-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
