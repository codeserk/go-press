package schema

import (
	"log"
	"press/core/field"
	"press/core/node"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id" validate:"required"`
	RealmID  primitive.ObjectID `json:"realmId" validate:"required"`
	AuthorID primitive.ObjectID `json:"authorId" validate:"required"`
	Type     node.Type          `json:"type" enums:"scene,model,view" validate:"required"`
	Name     string             `json:"name" validate:"required"`
	Fields   []*field.Entity    `json:"fields" validate:"required"`
}

// DefaultValue Gets the default value for the schema
func (e *Entity) DefaultValue() map[string]interface{} {
	result := make(map[string]interface{})
	for _, field := range e.Fields {
		value, err := field.Primitive.DefaultValue(field.Config)
		if err != nil {
			log.Print(err)
		}

		result[field.Key] = value
	}

	return result
}
