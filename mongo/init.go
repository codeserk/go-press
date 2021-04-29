package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Schemas *mongo.Collection
var Fields *mongo.Collection
var Scenes *mongo.Collection
var Users *mongo.Collection
var Realms *mongo.Collection

func Connect(connectionURL string) (*mongo.Client, error) {
	cxt, cancel := CreateContext()
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(cxt, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(cxt, nil)
	if err != nil {
		return nil, err
	}

	// Init collections
	Schemas = client.Database("press").Collection("schemas")
	Fields = client.Database("press").Collection("fields")
	Scenes = client.Database("press").Collection("scenes")
	createUsersSchema(client)
	createRealmSchema(client)

	return client, nil
}

func createUsersSchema(client *mongo.Client) {
	ctx, cancel := CreateContext()
	defer cancel()

	Users = client.Database("press").Collection("users")
	_, err := Users.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		panic(err)
	}
}

func createRealmSchema(client *mongo.Client) {
	Realms = client.Database("press").Collection("realms")
}

func CreateContext() (context.Context, context.CancelFunc) {
	// nolint
	return context.WithTimeout(context.Background(), 10 * time.Second)
}
