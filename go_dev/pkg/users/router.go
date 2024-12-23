package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRoute(router *gin.RouterGroup) {
	router.GET("/", UsersRetrive)
}
