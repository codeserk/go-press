package field

import "press/core/primitive"

type CreateParams struct {
	SchemaID  string
	Key       string
	Name      string
	Primitive primitive.Type
	Config    interface{}
}

type UpdateParams struct {
	Key       *string
	Name      *string
	Primitive *primitive.Type
	Config    *interface{}
}

type Service interface {
	Create(params CreateParams) (*Entity, error)

	Update(fieldID string, params UpdateParams) (*Entity, error)

	GetBySchema(schemaID string) ([]*Entity, error)
}
