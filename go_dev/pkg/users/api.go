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

func DeleteUser(ctx *gin.Context) {
	var payload interface{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := DeleteRecord(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Done"})
}

func CreateUser(ctx *gin.Context) {
	var payload Users
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
