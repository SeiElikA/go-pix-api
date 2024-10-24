package validation

import (
	"github.com/go-playground/validator/v10"
	"slices"
)

func ValidPostType(fl validator.FieldLevel) bool {
	types := []string{"public", "only_follow", "only_self"}

	if slices.Contains(types, fl.Field().String()) {
		return true
	}

	return false
}
