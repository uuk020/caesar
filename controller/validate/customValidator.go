package validate

import (
	"net/http"
	"regexp"
	"unicode"

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
	err = cv.validator.RegisterValidation("complexpassword", complexPassword)
	if err != nil {
		return err
	}
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// chinaPhone 校验国内手机
func chinaPhone(sl validator.FieldLevel) bool {
	v := sl.Field().String()
	reg := regexp.MustCompile(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`)
	return reg.MatchString(v)
}

// complexPassword 校验复杂密码
func complexPassword(sl validator.FieldLevel) bool {
	var (
		isUpper   = false
		isLower   = false
		isNumber  = false
		isSpecial = false
	)

	str := sl.Field().String()
	if len(str) < 6 || len(str) > 12 {
		return false
	}

	for _, s := range str {
		switch {
		case unicode.IsUpper(s):
			isUpper = true
		case unicode.IsLower(s):
			isLower = true
		case unicode.IsNumber(s):
			isNumber = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			isSpecial = true
		default:
		}
	}
	return isUpper && isLower && isNumber && isSpecial
}

// Register 注册校验器
func Register(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}
