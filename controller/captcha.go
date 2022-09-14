package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mojocn/base64Captcha"
	"net/http"
	"time"
)

var store = base64Captcha.NewMemoryStore(10240, 3*time.Minute)

func GetCaptcha(c echo.Context) error {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		return err
	}
	type captchaR struct {
		Id    string `json:"id"`
		Image string `json:"image"`
	}
	return c.JSON(http.StatusOK, &captchaR{
		Id:    id,
		Image: b64s,
	})
}
