package repository

import (
	"fmt"
	"press/core/node"
	"press/core/schema"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type insertOneQuery struct {
	RealmID  primitive.ObjectID `bson:"realmId"`
	AuthorID primitive.ObjectID `bson:"authorId"`
	Type     node.Type          `json:"type"`
	Name     string             `json:"name"`
}

// CreateOne Creates one schema from the given parameters
func (r *mongoRepository) InsertOne(params schema.InsertOneParams) (*schema.Entity, error) {
	realmID, err := primitive.ObjectIDFromHex(params.RealmID)
	if err != nil {
		return nil, fmt.Errorf("invalid realm id: %v", err)
	}
	authorID, err := primitive.ObjectIDFromHex(params.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("invalid author id: %v", err)
	}

	query := insertOneQuery{realmID, authorID, params.Type, params.Name}

	result, err := mongo.Schemas.InsertOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.FindOneByID(objectID.Hex())
	}

	return nil, fmt.Errorf("error found while trying to convert the InsertedID into an ObjectID")
}
