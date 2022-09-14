package router

import (
	"caesar/controller"
	"caesar/global"
	"caesar/internal"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	g := e.Group("/api")
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! Welcome to Caesar Password")
	})
	g.POST("/login", controller.Login)
	g.POST("/register", controller.Register)
	g.GET("/captcha", controller.GetCaptcha)
	g.PUT("/email", controller.AgainSendEmail)
	g.POST("/activation", controller.Activation)
	g.POST("/password", controller.ResetPassword)
	a := g.Group("/auth")
	jwtConfig := middleware.JWTConfig{
		Claims:     &internal.JwtCustomClaims{},
		SigningKey: []byte(global.Setting.SecretKey),
	}
	a.Use(middleware.JWTWithConfig(jwtConfig))
	a.POST("/logout", controller.Logout)
	a.GET("/me", controller.Me)
	a.PUT("/me", controller.UpdateMe)
}
