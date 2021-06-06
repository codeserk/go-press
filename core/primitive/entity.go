package primitive

type Type string

const (
	Text    Type = "text"
	Boolean Type = "boolean"
	Number  Type = "number"
	Date    Type = "date"
)

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name" example:"Ryuk"`
}

type Resolver interface {
	DefaultConfig() interface{}
	DefaultValue(config interface{}) (interface{}, error)
	// ValidateConfig(config interface{}) error
	// ValidateData(data interface{}) error
}

func (t Type) Resolver() Resolver {
	switch t {
	case Text:
		return text
	case Boolean:
		return boolean
	case Number:
		return number
	case Date:
		return date
	default:
		return text
	}
}

func (t Type) DefaultConfig() interface{} {
	return t.Resolver().DefaultConfig()
}

func (t Type) DefaultValue(config interface{}) (interface{}, error) {
	return t.Resolver().DefaultValue(config)
}
