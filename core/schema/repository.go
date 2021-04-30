package schema

type InsertOneParams struct {
	RealmID  string
	AuthorID string
	Name     string
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)

	FindOneByID(id string) (*Entity, error)
	FindInRealm(realmID string) ([]*Entity, error)
}
