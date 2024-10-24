package routes

import (
	"github.com/gin-gonic/gin"
	"go-pix-api/src/config"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		if config.AppConfig == nil {
			config.InitializeConfig(c.Request)
		}
		c.Next()
	})

	apiGroup := r.Group("api/")
	{
		UserRoute(apiGroup)
		PostRoute(apiGroup)
	}

	return r
}
