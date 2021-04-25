package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Schemas *mongo.Collection
var Scenes *mongo.Collection

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

	return client, nil
}

func CreateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
