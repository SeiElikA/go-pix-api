package entity

type Image struct {
	ID        int64  `bson:"_id" json:"id"`
	Url       string `bson:"url" json:"url"`
	Width     int    `bson:"width" json:"width"`
	Height    int    `bson:"height" json:"height"`
	CreatedAt string `bson:"created_at" json:"created_at"`
}
