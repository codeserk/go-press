package node

type InsertOneParams struct {
	RealmID  string
	SchemaID string
	Type     Type
	Slug     string
	Name     string
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)

	FindInRealm(realmID string) ([]*Entity, error)
}
