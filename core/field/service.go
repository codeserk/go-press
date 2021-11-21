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
	Key         *string
	Name        *string
	Description *string
	Primitive   *primitive.Type
	Config      *primitive.Config
}

type Service interface {
	Create(params CreateParams) (*Entity, error)
	Update(fieldID string, params UpdateParams) (*Entity, error)
	Delete(fieldId string) error

	GetBySchema(schemaID string) ([]*Entity, error)
}
