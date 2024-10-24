package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-pix-api/src/entity"
	"go-pix-api/src/exception"
	"go-pix-api/src/models/comment"
	"go-pix-api/src/services/post"
	user2 "go-pix-api/src/services/user"
	"go-pix-api/src/utils"
	"slices"
	"strconv"
)

func SelfPostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("post_id")

		// check post id type is correct
		id, err := strconv.Atoi(idStr)
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.WrongDataTypeError())
			return
		}

		// check is post exist
		var service = post.NewPostService()
		post, err := service.FindPostById(int64(id))
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.PostNotExistsError())
			return
		}

		user := utils.GetUserFromContext(c)
		if user.Type == "ADMIN" {
			c.Next()
			return
		}

		// check is self post
		if user.ID != post.Author.ID {
			defer c.Abort()
			utils.ErrorResponse(c, exception.PermissionDenyError())
			return
		}

		c.Next()
	}
}

func PostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("post_id")

		// check post id type is correct
		id, err := strconv.Atoi(idStr)
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.WrongDataTypeError())
			return
		}

		// check is post exist
		var service = post.NewPostService()
		post, err := service.FindPostById(int64(id))
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.PostNotExistsError())
			return
		}

		user := utils.GetUserFromContext(c)
		if user.Type == "ADMIN" {
			c.Next()
			return
		}

		// check is only self post
		if post.Type == "only_self" && user.ID != post.Author.ID {
			defer c.Abort()
			utils.ErrorResponse(c, exception.PermissionDenyError())
			return
		}

		// check is only_follow
		if post.Type == "only_follow" {
			defer c.Abort()
			var followService = user2.NewFollowerService()
			authorFollowList, err := followService.FindFollows(post.Author.ID)
			if err != nil {
				utils.ErrorResponse(c, exception.PermissionDenyError())
				return
			}

			isFollowByAuthor := slices.ContainsFunc(authorFollowList, func(s *entity.Follower) bool {
				return s.FollowId == user.ID
			})

			if !isFollowByAuthor {
				utils.ErrorResponse(c, exception.PermissionDenyError())
				return
			}
		}

		c.Next()
	}
}

func SelfCommentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("comment_id")

		// check comment id type is correct
		id, err := strconv.Atoi(idStr)
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.WrongDataTypeError())
			return
		}

		// check comment is exist
		postId := utils.StringToInt64(c.Param("post_id"))
		postService := post.NewPostService()
		postWithComment, _ := postService.FindPostWithComment(postId)
		idx := slices.IndexFunc(postWithComment.Comments, func(e *comment.CommentResponse) bool {
			return e.ID == int64(id)
		})

		if idx == -1 {
			defer c.Abort()
			utils.ErrorResponse(c, exception.CommentNotExistsError())
			return
		}

		// check permission
		user := utils.GetUserFromContext(c)
		if user.Type == "ADMIN" {
			c.Next()
			return
		}

		if postWithComment.Comments[idx].User.ID != user.ID {
			defer c.Abort()
			utils.ErrorResponse(c, exception.PermissionDenyError())
			return
		}

		c.Next()
	}
}

func CommentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("comment_id")

		// check comment id type is correct
		id, err := strconv.Atoi(idStr)
		if err != nil {
			defer c.Abort()
			utils.ErrorResponse(c, exception.WrongDataTypeError())
			return
		}

		postId := utils.StringToInt64(c.Param("post_id"))
		postService := post.NewPostService()
		postWithComment, _ := postService.FindPostWithComment(postId)
		idx := slices.IndexFunc(postWithComment.Comments, func(e *comment.CommentResponse) bool {
			return e.ID == int64(id)
		})

		if idx == -1 {
			defer c.Abort()
			utils.ErrorResponse(c, exception.CommentNotExistsError())
			return
		}

		c.Next()
	}
}
