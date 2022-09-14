package validate

import (
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates the fields of a struct.
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.RegisterValidation("chinaphone", chinaPhone)
	if err != nil {
		return err
	}
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func chinaPhone(sl validator.FieldLevel) bool {
	v := sl.Field().String()
	reg := regexp.MustCompile(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`)
	return reg.MatchString(v)
}

func Register(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}
