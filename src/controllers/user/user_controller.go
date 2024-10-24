package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-pix-api/src/controllers"
	"go-pix-api/src/entity"
	"go-pix-api/src/exception"
	"go-pix-api/src/models"
	post2 "go-pix-api/src/models/post"
	"go-pix-api/src/models/user"
	"go-pix-api/src/services/post"
	user2 "go-pix-api/src/services/user"
	"go-pix-api/src/utils"
	"slices"
)

type UserController struct {
	controllers.BaseController
	service       *user2.UserService
	postService   *post.PostService
	followService *user2.FollowerService
}

func NewUserController() *UserController {
	instance := new(UserController)
	instance.service = user2.NewUserService()
	instance.followService = user2.NewFollowerService()
	instance.postService = post.NewPostService()
	return instance
}

/*
public method
*/
func (controller *UserController) Login(c *gin.Context) {
	var request = utils.GetRequestFromContext[user.UserLoginRequest](c)

	result, err := controller.service.FindUserByEmailPwd(request.Email, request.Password)

	if err != nil {
		var cusErr *models.AppError
		if errors.As(err, &cusErr) {
			utils.ErrorResponse(c, err.(*models.AppError))
			return
		}

		utils.ErrorResponseWithData(c, exception.InternalServerError(), err)
		return
	}

	utils.SuccessResponse(c, result, "")
}

func (controller *UserController) Register(c *gin.Context) {
	var request = utils.GetRequestFromContext[user.UserRegisterRequest](c)
	newUser, err := controller.service.CreateUser(request)
	if err != nil {
		var cusErr *models.AppError
		if errors.As(err, &cusErr) {
			utils.ErrorResponse(c, err.(*models.AppError))
			return
		}

		utils.ErrorResponseWithData(c, exception.InternalServerError(), err)
		return
	}

	utils.SuccessResponse(c, newUser, "")
}

func (controller *UserController) Logout(c *gin.Context) {
	user := utils.GetUserFromContext(c)

	controller.service.RemoveToken(user.ID)

	utils.SuccessResponse(c, "", "")
}

func (controller *UserController) GetProfile(c *gin.Context) {
	userId := utils.StringToInt64(c.Param("user_id"))
	user, err := controller.service.FindUserById(userId)

	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, user, "")
}

func (controller *UserController) EditProfile(c *gin.Context) {
	userId := utils.StringToInt64(c.Param("user_id"))
	request := utils.GetRequestFromContext[user.UserEditProfileRequest](c)

	user, err := controller.service.EditUser(userId, request)

	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, user, "")
}

func (controller *UserController) FollowUser(c *gin.Context) {
	followId := utils.StringToInt64(c.Param("user_id"))
	user := utils.GetUserFromContext(c)
	err := controller.followService.CreateFollow(user, followId)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, "", "")
}

func (controller *UserController) UnFollowUser(c *gin.Context) {
	followId := utils.StringToInt64(c.Param("user_id"))
	user := utils.GetUserFromContext(c)
	err := controller.followService.DeleteFollow(user, followId)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, "", "")
}

func (controller *UserController) GetFollows(c *gin.Context) {
	followId := utils.StringToInt64(c.Param("user_id"))
	request := utils.GetRequestFromContext[user.UserQueryFollowRequest](c)

	userList, err := controller.followService.FindFollowUsers(followId, request)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, user.UserListResponse{
		TotalCount: len(userList),
		Posts:      userList,
	}, "")
}

func (controller *UserController) GetUserPosts(c *gin.Context) {
	targetUserId := utils.StringToInt64(c.Param("user_id"))
	_user := utils.GetUserFromContext(c)
	request := utils.GetRequestFromContext[user.UserQueryFollowRequest](c)

	authorFollowList, err := controller.followService.FindFollows(targetUserId)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}
	isFollowByAuthor := slices.ContainsFunc(authorFollowList, func(s *entity.Follower) bool {
		return s.FollowId == _user.ID
	})

	posts, err := controller.postService.FindUserPosts(targetUserId, isFollowByAuthor, request)
	if err != nil {
		utils.ErrorResponseWithData(c, exception.InternalServerError(), err.Error())
		return
	}

	utils.SuccessResponse(c, post2.PostListResponse{
		TotalCount: len(posts),
		Posts:      posts,
	}, "")
}
