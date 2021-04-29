package schema

import (
	"press/core/schema/field"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	RealmID primitive.ObjectID `json:"realmId"`
	Name    string             `json:"name"`
	Fields  []*field.Entity    `json:"fields"`
}
