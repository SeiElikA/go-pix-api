package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pix-api/src/utils"
)

type UserResponse struct {
	ID           int64  `bson:"_id" json:"id"`
	Email        string `bson:"email" json:"email"`
	Nickname     string `bson:"nickname" json:"nickname"`
	ProfileImage string `bson:"profile_image" json:"profile_image"`
	Type         string `bson:"type" json:"type"`
	AccessToken  string `json:"access_token,omitempty"`
}

func (user *UserResponse) ProfileWithServerUrl(c *gin.Context) {
	user.ProfileImage = fmt.Sprintf("%s%s", utils.GetServerUrl(c), user.ProfileImage)
}
