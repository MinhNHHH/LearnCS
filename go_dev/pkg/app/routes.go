package app

import (
	"github.com/MinhNHHH/go_dev/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (app *Application) Routes() *gin.Engine {
	r := gin.Default()
	limiter := middleware.NewTokenBucket(10, 100)
	r.Use(middleware.EnableCors())
	r.Use(middleware.RateLimiterMiddleWare(limiter))

	apiGroups := r.Group("/api")
	{
		// apiGroups.Use(middleware.Authentication())

		// userGroups := apiGroups.Group("/users")
		// {
		// 	userGroups.GET("/", func(ctx *gin.Context) {
		// 		ctx.JSON(200, gin.H{
		// 			"message": "User ====",
		// 		})
		// 	})
		// }
		crawlGroups := apiGroups.Group("/crawl")
		{
			crawlGroups.POST("/", app.GetDocumentApi())
		}
	}

	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
