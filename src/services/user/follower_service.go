package user

import (
	"context"
	"go-pix-api/src/config"
	"go-pix-api/src/entity"
	"go-pix-api/src/models/user"
	"go-pix-api/src/services"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"slices"
	"time"
)

type FollowerService struct {
	services.BaseService
}

func NewFollowerService() *FollowerService {
	service := new(FollowerService)
	service.Collection = config.DB.Collection("follower")
	return service
}

/*
public method
*/
func (service *FollowerService) FindFollows(userId int64) ([]*entity.Follower, error) {
	ctx := context.Background()
	var cursor, err = service.Collection.Find(ctx, bson.M{
		"user_id": userId,
	})
	if err != nil {
		return nil, err
	}

	var follower []*entity.Follower
	for cursor.Next(ctx) {
		var data entity.Follower
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		follower = append(follower, &data)
	}

	return follower, nil
}

func (service *FollowerService) CreateFollow(user *entity.User, followUserId int64) error {
	// check is exist follow
	var follows, err = service.FindFollows(user.ID)
	if err != nil {
		return err
	}
	isExist := slices.ContainsFunc(follows, func(data *entity.Follower) bool {
		return data.FollowId == followUserId
	})
	if isExist || user.ID == followUserId {
		return nil
	}

	var id = service.GetNextID()
	_, err = service.Collection.InsertOne(context.TODO(), &entity.Follower{
		ID:        id,
		UserId:    user.ID,
		FollowId:  followUserId,
		CreatedAt: time.Now().String(),
	})
	return err
}

func (service *FollowerService) DeleteFollow(user *entity.User, followUserId int64) error {
	// check is exist follow
	var follows, err = service.FindFollows(user.ID)
	if err != nil {
		return err
	}
	isExist := slices.ContainsFunc(follows, func(data *entity.Follower) bool {
		return data.FollowId == followUserId
	})
	if !isExist {
		return nil
	}

	_, err = service.Collection.DeleteOne(context.TODO(), bson.M{
		"user_id":   user.ID,
		"follow_id": followUserId,
	})
	return err
}

func (service *FollowerService) FindFollowUsers(userId int64, query *user.UserQueryFollowRequest) ([]*user.UserResponse, error) {
	ctx := context.TODO()

	// Default value
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}
	if query.OrderBy == "" {
		query.OrderBy = "created_at"
	}
	if query.OrderType == "" {
		query.OrderType = "desc"
	}

	// Query
	filter := bson.M{
		"user_id": userId,
	}

	// Order
	sortOrder := -1
	if query.OrderType == "asc" {
		sortOrder = 1
	}
	sort := bson.D{{Key: query.OrderBy, Value: sortOrder}}

	// Pagination
	opts := options.Find()
	opts.SetSort(sort)
	opts.SetSkip((query.Page - 1) * query.PageSize)
	opts.SetLimit(query.PageSize)

	cursor, err := service.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var userService = NewUserService()
	var users = make([]*user.UserResponse, 0)
	for cursor.Next(ctx) {
		var userEntity entity.User
		if err := cursor.Decode(&userEntity); err != nil {
			return nil, err
		}

		var userResponse, err = userService.FindUserById(userEntity.ID)
		if err != nil {
			return nil, err
		}

		users = append(users, userResponse)
	}

	return users, nil
}
