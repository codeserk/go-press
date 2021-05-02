package node

type CreateParams struct {
	RealmID  string
	SchemaID string
	Slug     string
	Name     string
}

type Service interface {
	Create(params CreateParams) (*Entity, error)

	GetInRealm(realmID string) ([]*Entity, error)
}
