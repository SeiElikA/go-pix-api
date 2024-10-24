package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-pix-api/src/exception"
	"go-pix-api/src/utils"
	"go-pix-api/src/validation"
	"strings"
)

func ValidationMiddleware[T any](request T) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.ShouldBind(&request)
		if err != nil {
			defer c.Abort()
			if strings.Contains(err.Error(), "required") {
				utils.ErrorResponse(c, exception.MissingFieldError())
				return
			}

			for _, valid := range validation.GetValidations() {
				if strings.Contains(err.Error(), valid.Name) {
					utils.ErrorResponse(c, valid.Error)
					return
				}
			}

			utils.ErrorResponse(c, exception.WrongDataTypeError())
			return
		}
		c.Set("request", request)
		c.Next()
	}
}
