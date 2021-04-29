package schema

type InsertOneFieldParams struct {
	Name      string
	Primitive int
	Data      interface{}
}

type InsertOneParams struct {
	RealmID  string
	AuthorID string
	Name     string
	Fields   []*InsertOneFieldParams
}

type Repository interface {
	InsertOne(params InsertOneParams) (*Entity, error)

	FindOneByID(id string) (*Entity, error)
	Find() ([]*Entity, error)
}
