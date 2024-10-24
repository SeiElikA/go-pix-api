package user

import "mime/multipart"

type UserRegisterRequest struct {
	Email        string                `form:"email" binding:"required,email"`
	Nickname     string                `form:"nickname" binding:"required"`
	Password     string                `form:"password" binding:"required,password-rule"`
	ProfileImage *multipart.FileHeader `form:"profile_image" binding:"required,image-type"`
}
