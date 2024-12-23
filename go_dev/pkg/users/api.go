package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRetrive(ctx *gin.Context) {
	users, err := GetAllUsers(ctx.Request.URL.Query().Get("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"user": users})
}
