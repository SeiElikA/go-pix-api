package post

import (
	"context"
	"go-pix-api/src/config"
	"go-pix-api/src/entity"
	"go-pix-api/src/models/post"
	user2 "go-pix-api/src/models/user"
	"go-pix-api/src/services"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"slices"
	"time"
)

type FavoriteService struct {
	services.BaseService
}

func NewFavoriteService() *FavoriteService {
	service := new(FavoriteService)
	service.Collection = config.DB.Collection("favorite")
	return service
}

func (service *FavoriteService) FindFavorite(userId int64) ([]*entity.Favorite, error) {
	ctx := context.Background()
	var cursor, err = service.Collection.Find(ctx, bson.M{
		"user_id": userId,
	})
	if err != nil {
		return nil, err
	}

	var follower []*entity.Favorite
	for cursor.Next(ctx) {
		var data entity.Favorite
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		follower = append(follower, &data)
	}

	return follower, nil
}

func (service *FavoriteService) FindFavoritePosts(userId int64, query *user2.UserQueryFollowRequest) ([]*post.PostResponse, error) {
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

	var userService = NewPostService()
	var users = make([]*post.PostResponse, 0)
	for cursor.Next(ctx) {
		var postEntity entity.Post
		if err := cursor.Decode(&postEntity); err != nil {
			return nil, err
		}

		var postResponse, err = userService.FindPostById(postEntity.ID)
		if err != nil {
			return nil, err
		}

		users = append(users, postResponse)
	}

	return users, nil
}

func (service *FavoriteService) CreateFavoritePost(userId int64, postId int64) error {
	var id = service.GetNextID()
	_, err := service.Collection.InsertOne(context.TODO(), &entity.Favorite{
		ID:         id,
		UserId:     userId,
		FavoriteId: postId,
		CreatedAt:  time.Now().String(),
	})
	return err
}

func (service *FavoriteService) RemoveFavoritePost(userId int64, postId int64) error {
	_, err := service.Collection.DeleteOne(context.TODO(), bson.M{
		"user_id":     userId,
		"favorite_id": postId,
	})
	return err
}

func (service *FavoriteService) IsExistFavorite(userId int64, postId int64) (bool, error) {
	var follows, err = service.FindFavorite(userId)
	if err != nil {
		return false, err
	}
	isExist := slices.ContainsFunc(follows, func(data *entity.Favorite) bool {
		return data.FavoriteId == postId
	})
	return isExist, err
}
