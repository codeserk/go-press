package repository

import (
	"context"
	"errors"
	"fmt"

	"press/core/user"
	"press/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context = context.TODO()

type Mongo struct {
	client *mongodb.Client
}

func New(client *mongodb.Client) Interface {
	return &Mongo{client: client}
}

// -- Create one

type CreateOneQuery struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Creates one user from the given params.
func (r *Mongo) CreateOne(params CreateOneParams) (*user.Entity, error) {
	query := CreateOneQuery(params)
	result, err := mongo.Users.InsertOne(ctx, query)
	if err != nil {
		return nil, errors.New("error found while trying to insert a new user: " + err.Error())
	}
	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return r.FindOneByID(objectID.Hex())
	}

	return nil, fmt.Errorf("error found while trying to convert the InsertedID into an ObjectID")
}

// Tries to find a user by its id.
func (r *Mongo) FindOneByID(id string) (*user.Entity, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error found while trying to convert ObjectId '%s': %v", id, err)
	}

	var createdUser user.Entity
	err = mongo.Users.FindOne(ctx, bson.M{"_id": objectID}).Decode(&createdUser)
	if err != nil {
		return nil, fmt.Errorf("error found while tryign to retrieve the user by its id `%s`: %v", objectID, err)
	}

	return &createdUser, nil
}

// Tries to find a user by its id.
func (r *Mongo) FindOneByEmail(email string) (*user.Entity, error) {
	query := bson.M{"email": email}

	var foundUser user.Entity
	err := mongo.Users.FindOne(ctx, query).Decode(&foundUser)
	if err != nil {
		if err == mongodb.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("error found while tryign to retrieve the user by its email '%s': %v", email, err)
	}

	return &foundUser, nil
}
