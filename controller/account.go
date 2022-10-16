package controller

import (
	"caesar/controller/forms"
	"caesar/internal"
	"caesar/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAccount 创建平台账号
func CreateAccount(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	r := new(forms.AccountCreate)
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

// ReadAccount 查看平台账号
func ReadAccount(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	a := new(forms.AccountRead)
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

// UpdateAccount 更新平台账号
func UpdateAccount(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	a := new(forms.AccountUpdate)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	r, err := service.UpdateAccount(a, claims.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

// DeleteAccount 删除平台账号
func DeleteAccount(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	a := new(forms.AccountRead)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	accountId, err := strconv.Atoi(a.ID)
	if err != nil {
		return err
	}
	err = service.DeleteAccount(a.MainPassword, int64(accountId), claims.Id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "删除成功"})
}

// GetLog 获取日志
func GetLog(c echo.Context) error {
	accountIdParam := c.Param("id")
	accountId, err := strconv.Atoi(accountIdParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	pageSizeParam := c.QueryParam("page_size")
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		pageSize = 10
	}

	r, count := service.GetLog(accountId, page, pageSize)
	return c.JSON(http.StatusOK, List{
		Data:     r,
		Count:    count,
		Page:     page,
		PageSize: pageSize,
	})
}

// GetList 获取列表
func GetList(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}

	a := new(forms.AccountList)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	r, count := service.GetList(a, claims.Id)
	return c.JSON(http.StatusOK, List{
		Data:     r,
		Count:    count,
		Page:     a.Page,
		PageSize: a.PageSize,
	})
}
