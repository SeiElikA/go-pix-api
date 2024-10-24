package validation

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"strings"
)

func ValidImageType(fl validator.FieldLevel) bool {
	v := fl.Field()
	x := v.Interface()
	var file = x.(multipart.FileHeader)
	return _IsValidImage(&file)
}

func ValidImagesType(fl validator.FieldLevel) bool {
	v := fl.Field()
	x := v.Interface()
	var files = x.([]*multipart.FileHeader)
	for _, file := range files {
		if !_IsValidImage(file) {
			return false
		}
	}
	return true
}

func _IsValidImage(file *multipart.FileHeader) bool {
	ext := strings.ReplaceAll(file.Header.Get("Content-Type"), "image/", "")
	if ext != "png" && ext != "jpeg" {
		return false
	}
	return true
}
