package schema

type CreateFieldParams struct {
	Name      string
	Primitive int
	Data      interface{}
}

type CreateParams struct {
	RealmID  string
	AuthorID string
	Name     string
	Fields   []*CreateFieldParams
}

type Service interface {
	Create(params CreateParams) (*Entity, error)
}
