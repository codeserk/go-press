package field

import (
	pressPrimitive "press/core/primitive"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID          primitive.ObjectID     `json:"id" bson:"_id"`
	SchemaID    primitive.ObjectID     `json:"schemaId"`
	Key         string                 `json:"key"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Primitive   pressPrimitive.Type    `json:"primitive"`
	Config      map[string]interface{} `json:"config" bson:"config"`
}
