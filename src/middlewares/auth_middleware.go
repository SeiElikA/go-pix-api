package middlewares

import (
	"context"
	"go-pix-api/src/config"
	"go-pix-api/src/entity"
	"go-pix-api/src/exception"
	"go-pix-api/src/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if token == "" {
			utils.ErrorResponse(c, exception.InvalidAccessTokenError())
			c.Abort()
			return
		}

		var user entity.User
		collection := config.DB.Collection("users")
		err := collection.FindOne(context.TODO(), bson.M{"access_token": token}).Decode(&user)

		if err != nil {
			utils.ErrorResponse(c, exception.InvalidAccessTokenError())
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
