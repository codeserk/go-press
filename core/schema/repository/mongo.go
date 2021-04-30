package repository

import (
	"context"
	"fmt"
	"press/core/schema"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type insertOneQuery struct {
	RealmID  primitive.ObjectID `bson:"realmId"`
	AuthorID primitive.ObjectID `bson:"authorId"`
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

	query := insertOneQuery{realmID, authorID, params.Name}

	result, err := mongo.Schemas.InsertOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.FindOneByID(objectID.Hex())
	}

	return nil, fmt.Errorf("error found while trying to convert the InsertedID into an ObjectID")
}

// FindOneByID Tries to find a schema by its id. Returns nil if the schema was not found.
func (r *mongoRepository) FindOneByID(id string) (*schema.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error found while trying to convert ObjectId '%s': %v", id, err)
	}

	var result schema.Entity
	err = mongo.Schemas.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongodb.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error found while tryign to retrieve the schema by its id `%s`: %v", objectID, err)
	}

	return &result, nil
}

// Finds all the schemas.
func (r *mongoRepository) Find() ([]*schema.Entity, error) {
	var result []*schema.Entity

	cursor, err := mongo.Schemas.Find(ctx, bson.D{{}})
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
