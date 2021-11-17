package field

import (
	pressPrimitive "press/core/primitive"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID          primitive.ObjectID     `json:"id" bson:"_id" validate:"required"`
	SchemaID    primitive.ObjectID     `json:"schemaId" validate:"required"`
	Key         string                 `json:"key" validate:"required"`
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description" validate:"required"`
	Primitive   pressPrimitive.Type    `json:"primitive" validate:"required"`
	Config      map[string]interface{} `json:"config" bson:"config" validate:"required"`
}
