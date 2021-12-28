package repository

import (
	"fmt"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes one schema
func (r *mongoRepository) Delete(schemaID string) error {
	objectID, err := primitive.ObjectIDFromHex(schemaID)
	if err != nil {
		return fmt.Errorf("invalid ObjectId '%s': %v", schemaID, err)
	}

	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	_, err = mongo.Schemas.DeleteOne(ctx, filter)

	return err
}
