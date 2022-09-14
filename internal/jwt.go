package internal

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Id    int64  `json:"user_id"`
	Name  string `json:"user_name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func Claims(c echo.Context) (*JwtCustomClaims, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("无法获取用户 token")
	}
	claims, ok := user.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, errors.New("无法获取用户 claims")
	}
	return claims, nil
}
