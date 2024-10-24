package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-pix-api/src/exception"
	user2 "go-pix-api/src/services/user"
	"go-pix-api/src/utils"
	"strconv"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("user_id")

		// check user id type is correct
		id, err := strconv.Atoi(idStr)
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.WrongDataTypeError())
			return
		}

		// check is user exist
		var service = user2.NewUserService()
		_, err = service.FindUserById(int64(id))
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.UserNotExistsError())
			return
		}

		c.Next()
	}
}

func SelfUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("user_id")

		// check user id type is correct
		id, err := strconv.Atoi(idStr)
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.WrongDataTypeError())
			return
		}

		// check is user exist
		var service = user2.NewUserService()
		findUser, err := service.FindUserById(int64(id))
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.UserNotExistsError())
			return
		}

		// check is self
		user := utils.GetUserFromContext(c)
		if user.Type == "ADMIN" {
			defer c.Next()
			return
		}

		if user.ID != findUser.ID {
			defer c.Abort()
			utils.ErrorResponse(c, exception.PermissionDenyError())
			return
		}

		c.Next()
	}
}
