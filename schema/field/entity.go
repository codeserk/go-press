package field

import (
	presPrimitive "press/primitive"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	Name      string             `json:"name"`
	Primitive presPrimitive.Type `json:"primitive"`
	Data      interface{}        `json:"data"`
}

type SavedEntity struct {
	ID primitive.ObjectID `bson:"_id"`
	Entity
}
