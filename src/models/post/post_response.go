package post

import (
	"go-pix-api/src/models/image"
	"go-pix-api/src/models/user"
)

type PostResponse struct {
	ID           int64                  `bson:"_id" json:"id"`
	Author       user.UserResponse      `bson:"author" json:"author"`
	Images       []*image.ImageResponse `bson:"images" json:"images"`
	LikeCount    int                    `bson:"like_count" json:"like_count"`
	Content      string                 `bson:"content" json:"content"`
	Type         string                 `json:"type"`
	Tags         []string               `json:"tags"`
	LocationName *string                `json:"location_name"`
	Liked        bool                   `json:"liked"`
	UpdatedAt    string                 `json:"updated_at"`
	CreatedAt    string                 `json:"created_at"`
}
