package entity

type Comment struct {
	ID        int64  `bson:"_id" json:"id"`
	UserId    int64  `bson:"user_id" json:"user"`
	PostId    int64  `bson:"post_id" json:"post_id"`
	Content   string `bson:"content" json:"content"`
	CreatedAt string `bson:"created_at" json:"created_at"`
	UpdatedAt string `bson:"updated_at" json:"updated_at"`
}
