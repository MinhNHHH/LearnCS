package app

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	UserName string `json:"name"`
	jwt.RegisteredClaims
}

var jwtTokenExpiry = time.Minute * 15
var refreshTokenExpiry = time.Hour * 24

func (app *Application) generateTokenPair(user *User) (TokenPairs, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = "localhost"
	claims["iss"] = "localhost"

	claims["exp"] = time.Now().Add(jwtTokenExpiry).Unix()

	signedAccessToken, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		return TokenPairs{}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)

	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)

	refreshTokenClaims["exp"] = time.Now().Add(refreshTokenExpiry).Unix()

	signedRefreshToken, err := refreshToken.SignedString([]byte(app.JWTSecret))
	if err != nil {
		return TokenPairs{}, nil
	}

	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokenPairs, nil
}
