package schema

import "press/core/node"

type InsertOneParams struct {
	RealmID  string
	AuthorID string
	Type     node.Type
	Name     string
}

type UpdateOneParams struct {
	Type node.Type
	Name string
}

type Repository interface {
	FindOneByID(id string) (*Entity, error)
	FindInRealm(realmID string) ([]*Entity, error)
	InsertOne(params InsertOneParams) (*Entity, error)
	UpdateOne(id string, params UpdateOneParams) (*Entity, error)
	Delete(schemaID string) error
}
