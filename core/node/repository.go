package node

import "press/core/primitive"

type InsertOneParams struct {
	RealmID  string
	SchemaID string
	Slug     string
	Name     string
}

type PatchOneParams struct {
	Key       *string
	Name      *string
	Primitive *primitive.Type
	Config    *interface{}
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)
}
