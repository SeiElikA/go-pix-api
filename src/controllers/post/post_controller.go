package post

import (
	"github.com/gin-gonic/gin"
	"go-pix-api/src/controllers"
	"go-pix-api/src/exception"
	"go-pix-api/src/models/comment"
	post2 "go-pix-api/src/models/post"
	user2 "go-pix-api/src/models/user"
	"go-pix-api/src/services/post"
	"go-pix-api/src/utils"
)

type PostController struct {
	controllers.BaseController
	service         *post.PostService
	commentService  *post.CommentService
	favoriteService *post.FavoriteService
}

func NewPostController() *PostController {
	instance := new(PostController)
	instance.service = post.NewPostService()
	instance.commentService = post.NewCommentService()
	instance.favoriteService = post.NewFavoriteService()
	return instance
}

/*
*
public method
*/
func (controller *PostController) CreatePost(c *gin.Context) {
	request := utils.GetRequestFromContext[post2.PostCreateRequest](c)
	user := utils.GetUserFromContext(c)

	var post, err = controller.service.CreatePost(user, request)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, post, "")
}

func (controller *PostController) GetPublicPost(c *gin.Context) {
	request := utils.GetRequestFromContext[post2.PostQueryRequest](c)
	var resList, err = controller.service.FindPublicPost(request)

	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, post2.PostListResponse{
		TotalCount: len(resList),
		Posts:      resList,
	}, "")
}

func (controller *PostController) GetPost(c *gin.Context) {
	postId := utils.StringToInt64(c.Param("post_id"))
	postWithComment, err := controller.service.FindPostWithComment(postId)

	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, postWithComment, "")
}

func (controller *PostController) EditPost(c *gin.Context) {
	request := utils.GetRequestFromContext[post2.PostEditRequest](c)
	postId := utils.StringToInt64(c.Param("post_id"))

	post, err := controller.service.EditPost(postId, request)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, post, "")
}

func (controller *PostController) DeletePost(c *gin.Context) {
	postId := utils.StringToInt64(c.Param("post_id"))

	err := controller.service.DeletePostById(postId)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, "", "")
}

func (controller *PostController) PostComment(c *gin.Context) {
	request := utils.GetRequestFromContext[comment.CommentPostRequest](c)
	user := utils.GetUserFromContext(c)
	postId := utils.StringToInt64(c.Param("post_id"))

	comment, err := controller.commentService.CreateComment(request, user, postId)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, comment, "")
}

func (controller *PostController) EditComment(c *gin.Context) {
	request := utils.GetRequestFromContext[comment.CommentPostRequest](c)
	commentId := utils.StringToInt64(c.Param("comment_id"))

	comment, err := controller.commentService.EditComment(request, commentId)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, comment, "")
}

func (controller *PostController) DeleteComment(c *gin.Context) {
	commentId := utils.StringToInt64(c.Param("comment_id"))

	err := controller.commentService.DeleteComment(commentId)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, "", "")
}

func (controller *PostController) GetFavoritePost(c *gin.Context) {
	request := utils.GetRequestFromContext[user2.UserQueryFollowRequest](c)
	user := utils.GetUserFromContext(c)

	postList, err := controller.favoriteService.FindFavoritePosts(user.ID, request)

	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, post2.PostListResponse{
		TotalCount: len(postList),
		Posts:      postList,
	}, "")
}

func (controller *PostController) AddRemoveFavoritePost(c *gin.Context) {
	postId := utils.StringToInt64(c.Param("post_id"))
	user := utils.GetUserFromContext(c)

	isExist, err := controller.favoriteService.IsExistFavorite(user.ID, postId)
	if isExist {
		err = controller.favoriteService.RemoveFavoritePost(user.ID, postId)
	} else {
		err = controller.favoriteService.CreateFavoritePost(user.ID, postId)
	}
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, "", "")
}
