package entity

type Follower struct {
	ID        int64  `bson:"_id" json:"id"`
	UserId    int64  `bson:"user_id" json:"user_id"`
	FollowId  int64  `bson:"follow_id" json:"follow_id"`
	CreatedAt string `bson:"created_at" json:"created_at"`
}
