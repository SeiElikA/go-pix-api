package comment

type CommentPostRequest struct {
	Content string `form:"content" binding:"required,blank"`
}
