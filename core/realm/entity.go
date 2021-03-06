package realm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID   primitive.ObjectID `json:"id" bson:"_id" validate:"required"`
	Name string             `json:"name" validate:"required"`

	// Fist iteration, a realm belongs only to one user.
	AuthorID primitive.ObjectID `json:"authorId" bson:"author" validate:"required"`
}
