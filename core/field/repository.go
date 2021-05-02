package field

import "press/core/primitive"

type InsertOneParams struct {
	SchemaID  string
	Key       string
	Name      string
	Primitive primitive.Type
	Config    interface{}
}

type PatchOneParams struct {
	Key       *string
	Name      *string
	Primitive *primitive.Type
	Config    *interface{}
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)

	PatchOne(fieldID string, params PatchOneParams) (*Entity, error)

	FindOneByID(id string) (*Entity, error)
	FindBySchema(schemeID string) ([]*Entity, error)
}
