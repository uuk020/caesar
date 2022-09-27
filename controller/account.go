package controller

import (
	"caesar/controller/forms"
	"caesar/internal"
	"caesar/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateAccount 创建平台账号
func CreateAccount(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	r := new(forms.Account)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	id, err := service.CreateAccount(r, claims.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "创建成功",
		"id":      id,
	})
}

// ReadAccount - 查看平台账号
func ReadAccount(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	a := new(forms.AccountParams)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	r, err := service.ReadAccount(a, claims.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}
