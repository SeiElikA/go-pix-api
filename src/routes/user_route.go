package routes

import (
	"github.com/gin-gonic/gin"
	user2 "go-pix-api/src/controllers/user"
	"go-pix-api/src/middlewares"
	"go-pix-api/src/models/user"
)

func UserRoute(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		var controller = user2.NewUserController()

		userGroup.POST("/login",
			middlewares.ValidationMiddleware(user.UserLoginRequest{}),
			controller.Login)
		userGroup.POST("/register",
			middlewares.ValidationMiddleware(user.UserRegisterRequest{}),
			controller.Register)
		userGroup.POST("/logout",
			middlewares.AuthMiddleware(),
			controller.Logout)
		userGroup.GET(":user_id/profile",
			middlewares.AuthMiddleware(),
			middlewares.UserMiddleware(),
			controller.GetProfile)
		userGroup.POST(":user_id/profile",
			middlewares.AuthMiddleware(),
			middlewares.SelfUserMiddleware(),
			middlewares.ValidationMiddleware(user.UserEditProfileRequest{}),
			controller.EditProfile)
		userGroup.GET(":user_id/follow",
			middlewares.AuthMiddleware(),
			middlewares.UserMiddleware(),
			middlewares.ValidationMiddleware(user.UserQueryFollowRequest{}),
			controller.GetFollows)
		userGroup.POST(":user_id/follow",
			middlewares.AuthMiddleware(),
			middlewares.UserMiddleware(),
			controller.FollowUser)
		userGroup.DELETE(":user_id/follow",
			middlewares.AuthMiddleware(),
			middlewares.UserMiddleware(),
			controller.UnFollowUser)
		userGroup.GET(":user_id/post",
			middlewares.AuthMiddleware(),
			middlewares.UserMiddleware(),
			middlewares.ValidationMiddleware(user.UserQueryFollowRequest{}),
			controller.GetUserPosts)
	}
}
