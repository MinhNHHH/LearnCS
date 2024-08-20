package main

import (
	"net/http"

	"github.com/MinhNHHH/sysDesign/db"
	"github.com/MinhNHHH/sysDesign/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	rdb := db.NewRedis()
	r := gin.Default()
	limitter := middleware.InitRateLimiter(rdb)

	r.Use(limitter.RateLimiter())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
