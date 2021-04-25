package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Schemas *mongo.Collection
var Scenes *mongo.Collection
var Users *mongo.Collection

func Connect(connectionUrl string) (*mongo.Client, error) {
	context, cancel := CreateContext()
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context, nil)
	if err != nil {
		return nil, err
	}

	// Init collections
	Schemas = client.Database("press").Collection("schemas")
	Scenes = client.Database("press").Collection("scenes")
	createUsersSchema(client)

	return client, nil
}

func createUsersSchema(client *mongo.Client) {
	ctx, cancel := CreateContext()
	defer cancel()

	Users = client.Database("press").Collection("users")
	Users.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}

func CreateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
