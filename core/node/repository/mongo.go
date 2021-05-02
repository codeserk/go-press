package repository

import (
	"context"
	"fmt"
	"press/core/node"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context = context.TODO()

type mongoRepository struct {
	client *mongodb.Client
}

func New(client *mongodb.Client) node.Repository {
	return &mongoRepository{
		client: client,
	}
}

type insertOneQuery struct {
	RealmID  primitive.ObjectID `bson:"realmId"`
	SchemaID primitive.ObjectID `bson:"schemaId"`
	Slug     string             `json:"slug"`
	Name     string             `json:"name"`
	Data     bson.M             `json:"data"`
}

// CreateOne Creates one node from the given parameters
func (r *mongoRepository) InsertOne(params node.InsertOneParams) (*node.Entity, error) {
	realmID, err := primitive.ObjectIDFromHex(params.RealmID)
	if err != nil {
		return nil, fmt.Errorf("invalid realm id: %v", err)
	}
	schemaID, err := primitive.ObjectIDFromHex(params.SchemaID)
	if err != nil {
		return nil, fmt.Errorf("invalid schema id: %v", err)
	}

	query := insertOneQuery{realmID, schemaID, params.Slug, params.Name, bson.M{}}

	result, err := mongo.Nodes.InsertOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.FindOneByID(objectID.Hex())
	}

	return nil, fmt.Errorf("error found while trying to convert the InsertedID into an ObjectID")
}

// FindOneByID Tries to find a schema by its id. Returns nil if the schema was not found.
func (r *mongoRepository) FindOneByID(id string) (*node.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error found while trying to convert ObjectId '%s': %v", id, err)
	}

	var result node.Entity
	err = mongo.Nodes.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongodb.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error found while tryign to retrieve the node by its id `%s`: %v", objectID, err)
	}

	return &result, nil
}

// Finds all the schemas.
func (r *mongoRepository) FindInRealm(realmID string) ([]*node.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(realmID)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId '%s': %v", realmID, err)
	}

	var result []*node.Entity
	cursor, err := mongo.Nodes.Find(ctx, bson.M{"realmId": objectID})
	if err != nil {
		return nil, fmt.Errorf("error found while trying to retrieve the node by realm `%s`: %v", objectID, err)
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
