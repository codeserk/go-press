package node

type InsertOneParams struct {
	RealmID  string
	SchemaID string
	Type     Type
	Slug     string
	Name     string
	Data     map[string]interface{}
}

type PatchOneParams struct {
	Slug *string
	Name *string
	Data *map[string]interface{}
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)
	PatchOne(id string, params PatchOneParams) (*Entity, error)

	FindInRealm(realmID string) ([]*Entity, error)
	FindBySlug(realmID string, slug string) (*Entity, error)
}
