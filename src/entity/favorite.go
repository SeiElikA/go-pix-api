package entity

type Favorite struct {
	ID         int64  `bson:"_id" json:"id"`
	UserId     int64  `bson:"user_id" json:"user_id"`
	FavoriteId int64  `bson:"favorite_id" json:"favorite_id"`
	CreatedAt  string `bson:"created_at" json:"created_at"`
}
