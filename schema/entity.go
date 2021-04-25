package schema

import (
	"press/schema/field"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID     primitive.ObjectID  `bson:"_id"`
	Realm  string              `json:"realm"`
	Name   string              `json:"name"`
	Fields []field.SavedEntity `json:"fields"`
}
