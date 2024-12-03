package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		// sanity check
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no auth header"})
			ctx.Abort()
			return
		}

		// split the header on spaces
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid off header"})
			ctx.Abort()
			return
		}

		// check to see if we have the word "Bearer"
		if headerParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: no Bearer"})
			ctx.Abort()
			return
		}

		token := headerParts[1]

		// declare an empty Claims
		claims := &Claims{}
		// parse the token with our claims
		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			// validate the signing algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(""), nil
		})

		// check for an error; note that this catches expired tokens as well
		if err != nil {
			if strings.HasPrefix(err.Error(), "token is exprired by") {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "expried token"})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
