package comment

import "go-pix-api/src/models/user"

type CommentResponse struct {
	ID        int64              `json:"id"`
	User      *user.UserResponse `json:"user"`
	Content   string             `json:"content"`
	UpdatedAt string             `json:"updated_at"`
	CreatedAt string             `json:"created_at"`
}
