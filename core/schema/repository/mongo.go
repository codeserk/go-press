package repository

import (
	"context"
	"press/core/schema"

	mongodb "go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context = context.TODO()

type mongoRepository struct {
	client *mongodb.Client
}

func New(client *mongodb.Client) schema.Repository {
	return &mongoRepository{
		client: client,
	}
}
