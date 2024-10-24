package validation

import (
	"github.com/go-playground/validator/v10"
	"go-pix-api/src/services/post"
)

func ValidSelfPost(fl validator.FieldLevel) bool {
	service := post.NewPostService()
	id := fl.Field().Int()
	var post, err = service.FindPostById(id)
	if err != nil {
		return false
	}

	println(post)

	return true
}
