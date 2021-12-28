package repository

import (
	"fmt"
	"press/core/schema"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindOneByID Tries to find a schema by its id. Returns nil if the schema was not found.
func (r *mongoRepository) FindOneByID(id string) (*schema.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error found while trying to convert ObjectId '%s': %v", id, err)
	}

	var result schema.Entity
	query := []bson.M{
		{"$match": bson.M{"_id": objectID}},
		{
			"$lookup": bson.M{
				"from":         "fields",
				"localField":   "_id",
				"foreignField": "schemaId",
				"as":           "fields",
			},
		},
	}

	cursor, err := mongo.Schemas.Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	// Check if there are items
	if hasItems := cursor.Next(ctx); !hasItems {
		return nil, nil
	}
	err = cursor.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Finds all the schemas.
func (r *mongoRepository) FindInRealm(realmID string) ([]*schema.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(realmID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId '%s': %v", realmID, err)
	}

	result := make([]*schema.Entity, 0)

	query := []bson.M{
		{"$match": bson.M{"realmId": objectID}},
		{
			"$lookup": bson.M{
				"from":         "fields",
				"localField":   "_id",
				"foreignField": "schemaId",
				"as":           "fields",
			},
		},
	}

	cursor, err := mongo.Schemas.Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var newEntity schema.Entity
		err := cursor.Decode(&newEntity)
		if err != nil {
			return nil, err
		}

		result = append(result, &newEntity)
	}

	return result, nil
}
