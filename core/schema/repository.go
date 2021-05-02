package schema

import "press/core/node"

type InsertOneParams struct {
	RealmID  string
	AuthorID string
	Type     node.Type
	Name     string
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)

	FindOneByID(id string) (*Entity, error)
	FindInRealm(realmID string) ([]*Entity, error)
}
