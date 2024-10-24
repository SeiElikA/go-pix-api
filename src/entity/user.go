package entity

type User struct {
	ID           int64  `bson:"_id" json:"id"`
	Email        string `bson:"email" json:"email"`
	NickName     string `bson:"nickname" json:"nickname"`
	Password     string `bson:"password,omitempty" json:"password,omitempty"`
	ProfileImage string `bson:"profile_image,omitempty" json:"profile_image,omitempty"`
	Type         string `bson:"type" json:"type"`
	AccessToken  string `bson:"access_token" json:"access_token"`
}
