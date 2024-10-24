package routes

import (
	"github.com/gin-gonic/gin"
	"go-pix-api/src/controllers/post"
	"go-pix-api/src/middlewares"
	"go-pix-api/src/models/comment"
	post2 "go-pix-api/src/models/post"
	user2 "go-pix-api/src/models/user"
)

func PostRoute(r *gin.RouterGroup) {
	postGroup := r.Group("/post")
	{
		var controller = post.NewPostController()

		postGroup.POST("/",
			middlewares.AuthMiddleware(),
			middlewares.ValidationMiddleware(post2.PostCreateRequest{}),
			controller.CreatePost)

		postGroup.GET("/public",
			middlewares.ValidationMiddleware(post2.PostQueryRequest{}),
			controller.GetPublicPost)

		postGroup.GET("/:post_id",
			middlewares.AuthMiddleware(),
			middlewares.PostMiddleware(),
			controller.GetPost)

		postGroup.POST("/:post_id",
			middlewares.AuthMiddleware(),
			middlewares.SelfPostMiddleware(),
			middlewares.ValidationMiddleware(post2.PostEditRequest{}),
			controller.EditPost)

		postGroup.DELETE("/:post_id",
			middlewares.AuthMiddleware(),
			middlewares.SelfPostMiddleware(),
			controller.DeletePost)

		postGroup.POST("/:post_id/comment",
			middlewares.AuthMiddleware(),
			middlewares.PostMiddleware(),
			middlewares.ValidationMiddleware(comment.CommentPostRequest{}),
			controller.PostComment)

		postGroup.POST("/:post_id/comment/:comment_id",
			middlewares.AuthMiddleware(),
			middlewares.PostMiddleware(),
			middlewares.CommentMiddleware(),
			middlewares.ValidationMiddleware(comment.CommentPostRequest{}),
			controller.EditComment)

		postGroup.DELETE("/:post_id/comment/:comment_id",
			middlewares.AuthMiddleware(),
			middlewares.PostMiddleware(),
			middlewares.SelfCommentMiddleware(),
			controller.DeleteComment)

		postGroup.GET("/favorite",
			middlewares.AuthMiddleware(),
			middlewares.ValidationMiddleware(user2.UserQueryFollowRequest{}),
			controller.GetFavoritePost)

		postGroup.POST("/:post_id/favorite",
			middlewares.AuthMiddleware(),
			middlewares.PostMiddleware(),
			controller.AddRemoveFavoritePost)
	}
}
