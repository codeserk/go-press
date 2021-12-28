package repository

import (
	"fmt"
	"press/core/schema"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateOne Updates one schema
func (r *mongoRepository) UpdateOne(id string, params schema.UpdateOneParams) (*schema.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error found while trying to convert ObjectId '%s': %v", id, err)
	}

	query := bson.M{
		"$set": bson.M{
			"type": params.Type,
			"name": params.Name,
		},
	}

	_, err = mongo.Schemas.UpdateByID(ctx, objectID, query)
	if err != nil {
		return nil, err
	}

	return r.FindOneByID(objectID.Hex())
}
