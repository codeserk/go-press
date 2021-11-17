package node

import "go.mongodb.org/mongo-driver/bson/primitive"

type Type string

const (
	TypeScene Type = "scene"
	TypeModel Type = "model"
	TypeView  Type = "view"
)

type Entity struct {
	ID       primitive.ObjectID     `json:"id" bson:"_id" validate:"required"`
	RealmID  primitive.ObjectID     `json:"realmId" validate:"required"`
	SchemaID primitive.ObjectID     `json:"schemaId" validate:"required"`
	Slug     string                 `json:"slug" validate:"required"`
	Name     string                 `json:"name" validate:"required"`
	Type     Type                   `json:"type" enums:"scene,model,view" validate:"required"`
	Data     map[string]interface{} `json:"data" bson:"data" validate:"required"`
	Views    *[]View                `json:"views"`
}

type View struct {
	SchemaID primitive.ObjectID     `json:"schemaId" validate:"required"`
	Data     map[string]interface{} `json:"data" bson:"data" validate:"required"`
	Children []View                 `json:"children" validate:"required"`
}
