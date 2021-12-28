package schema

import "press/core/node"

type CreateParams struct {
	RealmID  string
	AuthorID string
	Type     node.Type
	Name     string
}

type UpdateParams struct {
	Type node.Type
	Name string
}

type Service interface {
	GetByID(schemaID string) (*Entity, error)
	GetInRealm(realmID string) ([]*Entity, error)
	Create(params CreateParams) (*Entity, error)
	Update(id string, params UpdateParams) (*Entity, error)
	Delete(id string) error
}
