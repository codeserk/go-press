package schema

import "press/core/node"

type CreateParams struct {
	RealmID  string
	AuthorID string
	Type     node.Type
	Name     string
}

type Service interface {
	Create(params CreateParams) (*Entity, error)

	GetInRealm(realmID string) ([]*Entity, error)
}
