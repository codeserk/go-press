package repository

import (
	"context"
	"press/mongo"
	"press/schema"

	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	client *mongodb.Client
}

func Create(client *mongodb.Client) Interface {
	return &Mongo{
		client: client,
	}
}

// Finds all the schemas.
func (r *Mongo) Create(context context.Context, schema schema.Entity) error {
	_, err := mongo.Schemas.InsertOne(context, schema)
	if err != nil {
		return err
	}

	return nil
}

// Finds all the schemas.
func (r *Mongo) Find(context context.Context) ([]*schema.Entity, error) {
	var result []*schema.Entity

	cursor, err := mongo.Schemas.Find(context, bson.D{{}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context)
	for cursor.Next(context) {
		var newEntity schema.Entity
		err := cursor.Decode(&newEntity)
		if err != nil {
			return nil, err
		}

		result = append(result, &newEntity)
	}

	return result, nil
}
