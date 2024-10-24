package config

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"time"
)

var DB *mongo.Database

type Collection struct {
	Name  string
	Index string
}

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load .env:", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("MongoDB client connect failed:", err)
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DB = client.Database(dbName)

	_initCollections([]Collection{
		{
			Name:  "users",
			Index: "email",
		},
		{
			Name:  "posts",
			Index: "",
		},
		{
			Name:  "comment",
			Index: "",
		},
		{
			Name:  "images",
			Index: "",
		},
		{
			Name:  "comment",
			Index: "",
		},
		{
			Name:  "follower",
			Index: "",
		},
		{
			Name:  "favorite",
			Index: "",
		},
	})
}

func _initCollections(collections []Collection) {
	var err error
	for _, c := range collections {
		collection := DB.Collection(c.Name)

		if c.Index != "" {
			indexModel := mongo.IndexModel{
				Keys:    bson.D{{Key: c.Index, Value: 1}},
				Options: options.Index().SetUnique(true),
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			_, err = collection.Indexes().CreateOne(ctx, indexModel)

			cancel()
		}

	}

	defer func() {
		if err != nil {
			log.Fatal("Create index failed:", err)
		}
	}()
}
