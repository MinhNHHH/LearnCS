package main

import (
	mdw "github.com/MinhNHHH/sys-design/middleware"
	"github.com/gin-gonic/gin"
)

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
