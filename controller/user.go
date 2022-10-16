package controller

import (
	"caesar/controller/forms"
	"caesar/internal"
	"caesar/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

// Register 注册
func Register(c echo.Context) error {
	r := new(forms.Register)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	if !store.Verify(r.CaptchaId, r.Captcha, true) {
		return echo.NewHTTPError(http.StatusBadRequest, "验证码错误")
	}
	id, err := service.Register(r)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "创建成功",
		"id":      id,
	})
}

// Activation 激活
func Activation(c echo.Context) error {
	a := new(forms.Activation)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(a); err != nil {
		return err
	}
	err := service.Activation(a)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "激活成功"})
}

// ResetPassword 重置密码
func ResetPassword(c echo.Context) error {
	r := new(forms.ResetPassword)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	err := service.ResetPassword(r)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "重置密码成功"})
}

// AgainSendEmail 重新发送激活或重置密码邮件
func AgainSendEmail(c echo.Context) error {
	a := new(forms.SendMailS)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(a); err != nil {
		return err
	}
	err := service.AgainSendEmail(a)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "发送成功, 请查看邮箱"})
}

// Login 登录
func Login(c echo.Context) error {
	l := new(forms.Login)
	if err := c.Bind(l); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(l); err != nil {
		return err
	}
	if !store.Verify(l.CaptchaId, l.Captcha, true) {
		return echo.NewHTTPError(http.StatusBadRequest, "验证码错误")
	}
	ip := c.RealIP()
	token, err := service.Login(l, ip)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

// Logout 注销
func Logout(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	err = service.Logout(claims.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "注销成功, 有缘再见~"})
}

// Me 个人资料
func Me(c echo.Context) error {
	claims, err := internal.Claims(c)

	if err != nil {
		return err
	}
	r, err := service.Me(claims.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, r)
}

// UpdateMe 更新个人资料
func UpdateMe(c echo.Context) error {
	claims, err := internal.Claims(c)
	if err != nil {
		return err
	}
	a := new(forms.CommonS)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(a); err != nil {
		return err
	}
	m := make(map[string]interface{})
	m["real_name"] = a.RealName
	m["email"] = a.Email
	m["phone"] = a.Phone
	err = service.UpdateMe(claims.Id, m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "更新资料成功"})
}
