package post

import (
	"context"
	"go-pix-api/src/config"
	"go-pix-api/src/entity"
	"go-pix-api/src/models/comment"
	"go-pix-api/src/services"
	"go-pix-api/src/services/user"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type CommentService struct {
	services.BaseService
}

func NewCommentService() *CommentService {
	service := new(CommentService)
	service.Collection = config.DB.Collection("comment")
	return service
}

/*
public method
*/
func (service *CommentService) FindCommentById(id int64) (*comment.CommentResponse, error) {
	var entity entity.Comment
	var userService = user.NewUserService()

	var err = service.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	user, err := userService.FindUserById(entity.UserId)

	var res = comment.CommentResponse{
		ID:        entity.ID,
		User:      user,
		Content:   entity.Content,
		UpdatedAt: entity.UpdatedAt,
		CreatedAt: entity.CreatedAt,
	}

	return &res, nil
}

func (service *CommentService) FindCommentsByPostId(postId int64) ([]*comment.CommentResponse, error) {
	userService := user.NewUserService()

	filter := bson.M{"post_id": postId}
	var ctx = context.TODO()
	cursor, err := service.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var comments []*comment.CommentResponse
	for cursor.Next(ctx) {
		var data entity.Comment
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		_user, _ := userService.FindUserById(data.UserId)
		comments = append(comments, &comment.CommentResponse{
			ID:        data.ID,
			User:      _user,
			Content:   data.Content,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	return comments, err
}

func (service *CommentService) CreateComment(comment *comment.CommentPostRequest, user *entity.User, postId int64) (*comment.CommentResponse, error) {
	var id = service.GetNextID()
	var _, err = service.Collection.InsertOne(context.TODO(), &entity.Comment{
		ID:        id,
		UserId:    user.ID,
		PostId:    postId,
		Content:   comment.Content,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	})
	if err != nil {
		return nil, err
	}
	return service.FindCommentById(id)
}

func (service *CommentService) EditComment(comment *comment.CommentPostRequest, commentId int64) (*comment.CommentResponse, error) {
	updateValue := bson.D{
		{"content", comment.Content},
		{"updated_at", time.Now().String()},
	}
	_, err := service.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": commentId},
		bson.D{{"$set", updateValue}})
	if err != nil {
		return nil, err
	}
	return service.FindCommentById(commentId)
}

func (service *CommentService) DeleteComment(commentId int64) error {
	var _, err = service.Collection.DeleteOne(context.TODO(), bson.M{"_id": commentId})
	return err
}
