package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidPassword(fl validator.FieldLevel) bool {
	var password = fl.Field().String()
	if len(password) < 8 || len(password) > 24 {
		return false
	}

	return true
}
