package node

import "go.mongodb.org/mongo-driver/bson/primitive"

type Type string

const (
	Scene  Type = "scene"
	Nested Type = "nested"
)

type Entity struct {
	ID       primitive.ObjectID     `json:"id" bson:"_id"`
	RealmID  primitive.ObjectID     `json:"realmId"`
	SchemaID primitive.ObjectID     `json:"schemaId"`
	Type     Type                   `json:"type" enums:"scene,nested"`
	Slug     string                 `json:"slug"`
	Name     string                 `json:"name"`
	Data     map[string]interface{} `json:"data" bson:"data"`
}