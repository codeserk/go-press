package schema

import (
	"press/core/field"
	"press/core/node"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	RealmID  primitive.ObjectID `json:"realmId"`
	AuthorID primitive.ObjectID `json:"authorId"`
	Type     node.Type          `json:"type" enums:"scene,nested"`
	Name     string             `json:"name"`
	Fields   []*field.Entity    `json:"fields"`
}
