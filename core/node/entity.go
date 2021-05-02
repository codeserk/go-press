package node

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	ID       primitive.ObjectID     `json:"id" bson:"_id"`
	RealmID  primitive.ObjectID     `json:"realmId"`
	SchemaID primitive.ObjectID     `json:"schemaId"`
	Slug     string                 `json:"slug"`
	Name     string                 `json:"name"`
	Data     map[string]interface{} `json:"data" bson:"data"`
}