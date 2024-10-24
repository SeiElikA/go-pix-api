package post

import (
	"mime/multipart"
)

type PostCreateRequest struct {
	Content      string                  `form:"content" binding:"required,blank"`
	LocationName string                  `form:"location_name" binding:"-"`
	Tags         string                  `form:"tags" binding:"-"`
	Type         string                  `form:"type" binding:"required,post-type"`
	Images       []*multipart.FileHeader `form:"images[]" binding:"required,images-type"`
}
