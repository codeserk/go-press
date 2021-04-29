package field

import "press/core/primitive"

type InsertOneParams struct {
	SchemaID  string
	Name      string
	Primitive primitive.Type
	Data      interface{}
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)

	FindOneByID(id string) (*Entity, error)
	FindBySchema(schemeID string) ([]*Entity, error)
}
