package schema

type CreateParams struct {
	RealmID  string
	AuthorID string
	Name     string
}

type Service interface {
	Create(params CreateParams) (*Entity, error)

	GetInRealm(realmID string) ([]*Entity, error)
}
