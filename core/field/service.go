package field

import "press/core/primitive"

type CreateParams struct {
	SchemaID  string
	Name      string
	Primitive primitive.Type
	Data      interface{}
}

type Service interface {
	Create(params CreateParams) (*Entity, error)

	GetBySchema(schemaID string) ([]*Entity, error)
}
