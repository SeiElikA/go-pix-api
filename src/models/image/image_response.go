package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pix-api/src/utils"
)

type ImageResponse struct {
	ID        int64  `bson:"_id" json:"id"`
	Url       string `bson:"url" json:"url"`
	Width     int    `bson:"width" json:"width"`
	Height    int    `bson:"height" json:"height"`
	CreatedAt string `bson:"created_at" json:"created_at"`
}

func (image *ImageResponse) WithServerUrl(c *gin.Context) {
	image.Url = fmt.Sprintf("%s%s", utils.GetServerUrl(c), image.Url)
}
