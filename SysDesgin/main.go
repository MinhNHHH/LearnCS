package main

import (
	mdw "github.com/MinhNHHH/sys-design/middleware"
	rdb "github.com/MinhNHHH/sys-design/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	redis := rdb.NewRedisClient()
	r := gin.Default()
	r.Use(mdw.RateLimiterMiddleWare(mdw.NewTokenBucket(5, 10, *redis)))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
