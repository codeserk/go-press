package repository

import (
	"context"
	"fmt"

	"press/common/errors"
	"press/core/realm"
	"press/core/user"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context = context.TODO()

type mongoRepository struct {
	client *mongodb.Client
}

func New(client *mongodb.Client) Interface {
	return &mongoRepository{client: client}
}

type createOneQuery struct {
	Name   string             `json:"name"`
	Author primitive.ObjectID `bson:"author"`
}

// Creates one realm from the given params.
func (r *mongoRepository) CreateOne(params CreateOneParams) (*realm.Entity, error) {
	query := createOneQuery{params.Name, params.Author.ID}
	result, err := mongo.Realms.InsertOne(ctx, query)
	if err != nil {
		return nil, errors.New("error found while trying to insert a new user: " + err.Error())
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.FindOneByID(objectID.Hex())
	}

	return nil, fmt.Errorf("error found while trying to convert the InsertedID into an ObjectID")
}

type RealmUnpopulated struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name"`
	AuthorID primitive.ObjectID `bson:"author"`
}

// FindOneByID Tries to find a realm by its id. Returns nil if the user was not found.
func (r *mongoRepository) FindOneByID(id string) (*realm.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error found while trying to convert ObjectId '%s': %v", id, err)
	}

	var result realm.Entity
	err = mongo.Realms.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongodb.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error found while tryign to retrieve the realm by its id `%s`: %v", objectID, err)
	}

	return &result, nil
}

// Tries to find a user by its id.
func (r *mongoRepository) FindByAuthor(author *user.Entity) ([]*realm.Entity, error) {
	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.M{"author": author.ID}}}

	result := make([]*realm.Entity, 0)
	cursor, err := mongo.Realms.Aggregate(ctx, mongodb.Pipeline{matchStage})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return result, nil
}
