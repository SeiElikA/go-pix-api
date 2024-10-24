package post

import "go-pix-api/src/models/comment"

type PostCommentResponse struct {
	Post     *PostResponse              `json:"post"`
	Comments []*comment.CommentResponse `json:"comments"`
}
