package field

import (
	presPrimitive "press/core/primitive"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	Primitive presPrimitive.Type `json:"primitive"`
	Data      interface{}        `json:"data"`
}
