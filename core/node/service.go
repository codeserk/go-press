package node

type CreateParams struct {
	RealmID  string
	SchemaID string
	Type     Type
	Slug     string
	Name     string
}

type UpdateParams struct {
	Slug *string
	Name *string
	Data *interface{}
}

type Service interface {
	Create(params CreateParams) (*Entity, error)
	Update(id string, params UpdateParams) (*Entity, error)

	GetInRealm(realmID string) ([]*Entity, error)
}
