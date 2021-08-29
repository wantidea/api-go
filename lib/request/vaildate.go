package request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// mobile 验证手机号
func mobile(fl validator.FieldLevel) bool {
	regular := regexp.MustCompile(`1[345678][0-9]{9}$`)
	return regular.MatchString(fl.Field().String())
}
