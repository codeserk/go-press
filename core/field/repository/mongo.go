package repository

import (
	"context"
	"fmt"
	"press/core/field"
	pressPrimitive "press/core/primitive"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context = context.TODO()

type mongoRepository struct {
	client *mongodb.Client
}

func New(client *mongodb.Client) field.Repository {
	return &mongoRepository{
		client: client,
	}
}

type insertOneQuery struct {
	SchemaID  primitive.ObjectID  `bson:"schemaId"`
	Name      string              `json:"name"`
	Primitive pressPrimitive.Type `json:"primitive"`
	Data      interface{}         `json:"data"`
}

// CreateOne Creates one schema from the given parameters
func (r *mongoRepository) InsertOne(params field.InsertOneParams) (*field.Entity, error) {
	schemaID, err := primitive.ObjectIDFromHex(params.SchemaID)
	if err != nil {
		return nil, fmt.Errorf("invalid schema id: %v", err)
	}

	query := insertOneQuery{schemaID, params.Name, params.Primitive, params.Data}

	result, err := mongo.Fields.InsertOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.FindOneByID(objectID.Hex())
	}

	return nil, fmt.Errorf("error found while trying to convert the InsertedID into an ObjectID")
}

// FindOneByID Tries to find a schema by its id. Returns nil if the schema was not found.
func (r *mongoRepository) FindOneByID(id string) (*field.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error found while trying to convert ObjectId '%s': %v", id, err)
	}

	var result field.Entity
	err = mongo.Fields.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongodb.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error found while tryign to retrieve the field by its id `%s`: %v", objectID, err)
	}

	return &result, nil
}

// FindOneByID Tries to find a schema by its id. Returns nil if the schema was not found.
func (r *mongoRepository) FindBySchema(schemaID string) ([]*field.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(schemaID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId '%s': %v", schemaID, err)
	}

	var result []*field.Entity
	cursor, err := mongo.Fields.Find(ctx, bson.M{"schemaId": objectID})
	if err != nil {
		return nil, fmt.Errorf("error found while trying to retrieve the field by schema `%s`: %v", objectID, err)
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
