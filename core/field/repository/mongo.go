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
	Key       string              `json:"key"`
	Name      string              `json:"name"`
	Primitive pressPrimitive.Type `json:"primitive"`
	Config    interface{}         `json:"config"`
}

// CreateOne Creates one schema from the given parameters
func (r *mongoRepository) InsertOne(params field.InsertOneParams) (*field.Entity, error) {
	schemaID, err := primitive.ObjectIDFromHex(params.SchemaID)
	if err != nil {
		return nil, fmt.Errorf("invalid schema id: %v", err)
	}

	query := insertOneQuery{schemaID, params.Key, params.Name, params.Primitive, params.Config}

	result, err := mongo.Fields.InsertOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.FindOneByID(objectID.Hex())
	}

	return nil, fmt.Errorf("error found while trying to convert the InsertedID into an ObjectID")
}

type patchOneQuery struct {
	Key       *string              `json:"key"`
	Name      *string              `json:"name"`
	Primitive *pressPrimitive.Type `json:"primitive"`
	Config    *interface{}         `bson:"config"`
}

func (r *mongoRepository) PatchOne(fieldID string, params field.PatchOneParams) (*field.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(fieldID)
	if err != nil {
		return nil, fmt.Errorf("invalid field id: %v", err)
	}

	query := bson.M{}
	if params.Key != nil {
		query["key"] = params.Key
	}
	if params.Name != nil {
		query["name"] = params.Name
	}
	if params.Primitive != nil {
		query["primitive"] = params.Primitive
	}
	if params.Config != nil {
		query["config"] = params.Config
	}

	_, err = mongo.Fields.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": query})
	if err != nil {
		return nil, err
	}

	return r.FindOneByID(objectID.Hex())
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

	result := make([]*field.Entity, 0)
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
