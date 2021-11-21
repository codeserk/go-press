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
	Key         *string
	Name        *string
	Description *string
	Primitive   *primitive.Type
	Config      *primitive.Config
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)
	PatchOne(fieldID string, params PatchOneParams) (*Entity, error)
	DeleteOne(fieldID string) error

	FindOneByID(id string) (*Entity, error)
	FindBySchema(schemeID string) ([]*Entity, error)
}
