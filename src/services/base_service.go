package services

import (
	"context"
	"go-pix-api/src/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BaseService struct {
	Collection *mongo.Collection
}

func (service *BaseService) GetNextID() int64 {
	opts := options.FindOne().SetSort(bson.D{{"_id", -1}})
	var lastEntity entity.User
	var err = service.Collection.FindOne(context.TODO(), bson.M{}, opts).Decode(&lastEntity)
	if err != nil {
		return 1
	}

	if lastEntity.ID == 0 {
		return 1
	}

	return lastEntity.ID + 1
}
