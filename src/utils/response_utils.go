package utils

import (
	"github.com/gin-gonic/gin"
	"go-pix-api/src/entity"
	"go-pix-api/src/models"
	"net/http"
)

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, appError *models.AppError) {
	c.JSON(appError.Code, models.APIResponse{
		Success: false,
		Message: appError.Message,
		Data:    "",
	})
}

func ErrorResponseWithData(c *gin.Context, appError *models.AppError, data interface{}) {
	c.JSON(appError.Code, models.APIResponse{
		Success: false,
		Message: appError.Message,
		Data:    data,
	})
}

func GetRequestFromContext[T any](c *gin.Context) *T {
	return _GetObjFromContext[T](c, "request")
}

func GetUserFromContext(c *gin.Context) *entity.User {
	return _GetObjFromContext[entity.User](c, "user")
}

func _GetObjFromContext[T any](c *gin.Context, key string) *T {
	value, exists := c.Get(key)
	if !exists {
		return nil
	}

	var obj = value.(T)
	return &obj
}
