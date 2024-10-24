package validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidBlankContent(fl validator.FieldLevel) bool {
	str := strings.ReplaceAll(fl.Field().String(), " ", "")
	if len(str) == 0 {
		return false
	}
	return true
}
