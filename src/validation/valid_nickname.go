package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidNickname(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[a-z0-9_]+$`)
	return regex.MatchString(fl.Field().String())
}
