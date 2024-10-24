package user

import "mime/multipart"

type UserEditProfileRequest struct {
	Nickname     string                `form:"nickname" binding:"omitempty,min=4,max=16,nickname"`
	ProfileImage *multipart.FileHeader `form:"profile_image" binding:"required,image-type"`
}
