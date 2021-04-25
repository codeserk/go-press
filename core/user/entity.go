package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Email string             `json:"email"`
}
