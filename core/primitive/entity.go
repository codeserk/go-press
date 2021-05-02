package primitive

type Type string

const (
	Text Type = "text"
)

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name" example:"Ryuk"`
}

type Resolver interface {
	DefaultConfig() interface{}
	// ValidateConfig(config interface{}) error
	// ValidateData(data interface{}) error
}

func (t Type) Resolver() Resolver {
	switch t {
	case Text:
		return text
	default:
		return text
	}
}

func (t Type) DefaultConfig() interface{} {
	return t.Resolver().DefaultConfig()
}
