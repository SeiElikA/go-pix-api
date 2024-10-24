package post

import (
	"context"
	"go-pix-api/src/config"
	"go-pix-api/src/entity"
	"go-pix-api/src/models/comment"
	"go-pix-api/src/models/post"
	user2 "go-pix-api/src/models/user"
	"go-pix-api/src/services"
	"go-pix-api/src/services/image"
	"go-pix-api/src/services/user"
	"go-pix-api/src/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"strings"
	"time"
)

type PostService struct {
	services.BaseService
}

func NewPostService() *PostService {
	service := new(PostService)
	service.Collection = config.DB.Collection("posts")
	return service
}

/*
public method
*/
func (service *PostService) FindPublicPost(query *post.PostQueryRequest) ([]*post.PostResponse, error) {
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
	filter := bson.M{}
	if query.Content != "" {
		filter["content"] = bson.M{"$regex": query.Content, "$options": "i"}
	}
	if query.Tag != "" {
		filter["tags"] = bson.M{"$in": bson.A{query.Tag}}
	}
	if query.LocationName != "" {
		filter["location_name"] = query.LocationName
	}
	filter["type"] = "public"

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

	// Execute
	cursor, err := service.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var posts = make([]*post.PostResponse, 0)
	for cursor.Next(ctx) {
		var postEntity entity.Post
		if err := cursor.Decode(&postEntity); err != nil {
			return nil, err
		}

		var postResponse, err = service.FindPostById(postEntity.ID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, postResponse)
	}

	return posts, nil
}

func (service *PostService) FindPostById(id int64) (*post.PostResponse, error) {
	var entity entity.Post
	var userService = user.NewUserService()
	var imageService = image.NewImageService()

	var err = service.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	author, err := userService.FindUserById(entity.AuthorId)
	images, err := imageService.FindManyImageById(entity.ImageIds)

	var res = post.PostResponse{
		ID:           entity.ID,
		Author:       *author,
		LikeCount:    entity.LikeCount,
		Content:      entity.Content,
		Type:         entity.Type,
		Tags:         entity.Tags,
		LocationName: utils.StringOrNil(entity.LocationName),
		Images:       images,
		UpdatedAt:    entity.UpdatedAt,
		CreatedAt:    entity.CreatedAt,
	}

	return &res, nil
}

func (service *PostService) FindUserPosts(userId int64, isFollow bool, query *user2.UserQueryFollowRequest) ([]*post.PostResponse, error) {
	var ctx = context.Background()
	var userService = user.NewUserService()
	var imageService = image.NewImageService()
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

	// filter
	var filter = bson.M{"author_id": userId}
	var postType = bson.A{"public"}
	if isFollow {
		postType = append(postType, "only_follow")
	}
	filter["type"] = bson.M{
		"$in": postType,
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

	var cursor, err = service.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var posts = make([]*post.PostResponse, 0)
	for cursor.Next(ctx) {
		var postEntity entity.Post
		if err := cursor.Decode(&postEntity); err != nil {
			return nil, err
		}

		author, err := userService.FindUserById(postEntity.AuthorId)
		images, err := imageService.FindManyImageById(postEntity.ImageIds)

		var res = post.PostResponse{
			ID:           postEntity.ID,
			Author:       *author,
			LikeCount:    postEntity.LikeCount,
			Content:      postEntity.Content,
			Type:         postEntity.Type,
			Tags:         postEntity.Tags,
			LocationName: utils.StringOrNil(postEntity.LocationName),
			Images:       images,
			UpdatedAt:    postEntity.UpdatedAt,
			CreatedAt:    postEntity.CreatedAt,
		}

		if err != nil {
			return nil, err
		}

		posts = append(posts, &res)
	}

	return posts, nil
}

func (service *PostService) CreatePost(user *entity.User, request *post.PostCreateRequest) (*post.PostResponse, error) {
	var imageService = image.NewImageService()
	var imageIds, err = imageService.InsertManyImage("post", request.Images)
	if err != nil {
		return nil, err
	}

	var id = service.GetNextID()
	post := &entity.Post{
		ID:           id,
		AuthorId:     user.ID,
		Type:         request.Type,
		Tags:         strings.Split(request.Tags, " "),
		Content:      request.Content,
		LocationName: request.LocationName,
		ImageIds:     imageIds,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}

	_, err = service.Collection.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}
	return service.FindPostById(id)
}

func (service *PostService) EditPost(id int64, request *post.PostEditRequest) (*post.PostResponse, error) {
	updateValue := bson.D{
		{"type", request.Type},
		{"tags", strings.Split(request.Tags, " ")},
		{"content", request.Content},
		{"updated_at", time.Now().String()},
	}
	_, err := service.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{{"$set", updateValue}})

	if err != nil {
		return nil, err
	}

	return service.FindPostById(id)
}

func (service *PostService) DeletePostById(id int64) error {
	var _, err = service.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (service *PostService) FindPostWithComment(id int64) (*post.PostCommentResponse, error) {
	_post, err := service.FindPostById(id)
	if err != nil {
		return nil, err
	}

	commentService := NewCommentService()
	_comments, err := commentService.FindCommentsByPostId(id)
	if len(_comments) == 0 {
		_comments = make([]*comment.CommentResponse, 0)
	}

	if err != nil {
		return nil, err
	}

	return &post.PostCommentResponse{
		Post:     _post,
		Comments: _comments,
	}, err
}

/*
private method
*/
