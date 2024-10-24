package post

type PostEditRequest struct {
	Type    string `form:"type" binding:"required,post-type"`
	Tags    string `form:"tags" binding:"-"`
	Content string `form:"content" binding:"required,blank"`
}
