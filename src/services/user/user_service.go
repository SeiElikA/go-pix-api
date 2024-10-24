package user

import (
	"context"
	"errors"
	"go-pix-api/src/config"
	"go-pix-api/src/entity"
	"go-pix-api/src/exception"
	"go-pix-api/src/models/user"
	"go-pix-api/src/services"
	"go-pix-api/src/services/image"
	"go-pix-api/src/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"strings"
	"time"
)

type UserService struct {
	services.BaseService
}

func NewUserService() *UserService {
	service := new(UserService)
	service.Collection = config.DB.Collection("users")
	return service
}

/*
public method
*/
func (service *UserService) FindUserByEmailPwd(email string, password string) (*user.UserResponse, error) {
	var user entity.User
	var filter = bson.M{"email": email, "password": utils.HashPasswordSHA256(password)}
	var err = service.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, exception.InvalidLoginError()
	}

	token := utils.HashEmailSHA256(user.Email)
	_, err = service.Collection.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.D{{"$set", bson.D{{"access_token", token}}}})
	if err != nil {
		return nil, exception.InvalidLoginError()
	}

	res, err := service.FindUserById(user.ID)
	res.AccessToken = token

	return res, nil
}

func (service *UserService) RemoveToken(id int64) {
	service.Collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{{"$set", bson.D{{"access_token", ""}}}})
}

func (service *UserService) CreateUser(request *user.UserRegisterRequest) (*user.UserResponse, error) {
	if service._IsEmailExist(request.Email) {
		return nil, exception.UserExistError()
	}
	var imageService = image.NewImageService()

	hashedPassword := utils.HashPasswordSHA256(request.Password)
	profileImage, _, err := imageService.SaveImage("profile_image", request.ProfileImage)

	if err != nil {
		return nil, exception.ImageCanNotProcessError()
	}

	id := service.GetNextID()
	user := &entity.User{
		ID:           id,
		Email:        request.Email,
		NickName:     request.Nickname,
		Password:     hashedPassword,
		ProfileImage: profileImage,
		Type:         strings.ToUpper("user"),
	}

	_, err = service.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return service.FindUserById(id)
}

func (service *UserService) EditUser(userId int64, request *user.UserEditProfileRequest) (*user.UserResponse, error) {
	var imageService = image.NewImageService()
	profileImage, _, err := imageService.SaveImage("profile_image", request.ProfileImage)

	if err != nil {
		return nil, exception.ImageCanNotProcessError()
	}

	updateValue := bson.M{
		"profile_image": profileImage,
		"updated_at":    time.Now().String(),
	}

	if request.Nickname != "" {
		updateValue["nickname"] = request.Nickname
	}

	_, err = service.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userId},
		bson.D{{"$set", updateValue}})

	if err != nil {
		return nil, err
	}

	return service.FindUserById(userId)
}

func (service *UserService) FindUserById(id int64) (*user.UserResponse, error) {
	var insertedUser entity.User
	err := service.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&insertedUser)
	var res = user.UserResponse{
		ID:           insertedUser.ID,
		Email:        insertedUser.Email,
		Nickname:     insertedUser.NickName,
		ProfileImage: config.AppConfig.ServerURL + insertedUser.ProfileImage,
		Type:         insertedUser.Type,
	}
	return &res, err
}

//func (service *UserService) FindUserPost(selfUserId int64, targetUserId int64) ([]*post.PostResponse, error) {
//	//postService := post2.NewPostService()
//	var followService = NewFollowerService()
//
//	authorFollowList, err := followService.FindFollows(targetUserId)
//	if err != nil {
//		return nil, err
//	}
//	_  slices.ContainsFunc(authorFollowList, func(s *entity.Follower) bool {
//		return s.FollowId == selfUserId
//	})
//	//
//	//posts, err := postService.FindUserPosts(targetUserId, isFollowByAuthor)
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	return nil, nil
//}

/*
private method
*/
func (service *UserService) _IsEmailExist(email string) bool {
	var user entity.User
	collection := config.DB.Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false
	}
	return true
}
