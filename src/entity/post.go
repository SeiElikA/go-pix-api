package entity

type Post struct {
	ID           int64    `bson:"_id" json:"id"`
	AuthorId     int64    `bson:"author_id" json:"author"`
	ImageIds     []int64  `bson:"image_ids" json:"images"`
	LikeCount    int      `bson:"like_count" json:"like_count"`
	Content      string   `bson:"content" json:"content"`
	Type         string   `bson:"type" json:"type"`
	Tags         []string `bson:"tags" json:"tags"`
	LocationName string   `bson:"location_name" json:"location_name"`
	Liked        bool     `bson:"liked" json:"liked"`
	UpdatedAt    string   `bson:"updated_at" json:"updated_at"`
	CreatedAt    string   `bson:"created_at" json:"created_at"`
}
