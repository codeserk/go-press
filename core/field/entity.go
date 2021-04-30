package field

import (
	pressPrimitive "press/core/primitive"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID        primitive.ObjectID     `json:"id" bson:"_id"`
	SchemaID  primitive.ObjectID     `json:"schemaId"`
	Name      string                 `json:"name"`
	Primitive pressPrimitive.Type    `json:"primitive"`
	Data      map[string]interface{} `json:"data" bson:"data,inline"`
}
