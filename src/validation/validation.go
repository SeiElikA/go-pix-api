package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go-pix-api/src/exception"
	"go-pix-api/src/models"
)

type ValidationData struct {
	Name     string
	Function validator.Func
	Error    *models.AppError
}

func RegisterValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		for _, item := range GetValidations() {
			v.RegisterValidation(item.Name, item.Function)
		}
	}
}

func GetValidations() []ValidationData {
	return []ValidationData{
		{
			Name:     "post-type",
			Function: ValidPostType,
			Error:    exception.WrongDataTypeError(),
		},
		{
			Name:     "blank",
			Function: ValidBlankContent,
			Error:    exception.WrongDataTypeError(),
		},
		{
			Name:     "image-type",
			Function: ValidImageType,
			Error:    exception.ImageCanNotProcessError(),
		},
		{
			Name:     "images-type",
			Function: ValidImagesType,
			Error:    exception.ImageCanNotProcessError(),
		},
		{
			Name:     "password-rule",
			Function: ValidPassword,
			Error:    exception.PasswordNotSecureError(),
		},
		{
			Name:     "self-post",
			Function: ValidSelfPost,
			Error:    exception.PermissionDenyError(),
		},
		{
			Name:     "nickname",
			Function: ValidNickname,
			Error:    exception.WrongDataTypeError(),
		},
	}
}
