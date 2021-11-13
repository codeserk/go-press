package primitive

type Type string

const (
	Text    Type = "text"
	Boolean Type = "boolean"
	Number  Type = "number"
	Date    Type = "date"
	Options Type = "options"
	Node    Type = "node"
)

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name" example:"Ryuk"`
}

type Resolver interface {
	DefaultConfig() interface{}
	TransformConfig(config interface{}) (interface{}, error)
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
	case Options:
		return options
	case Node:
		return node
	default:
		return text
	}
}

func (t Type) DefaultConfig() interface{} {
	return t.Resolver().DefaultConfig()
}

func (t Type) TransformConfig(config interface{}) (interface{}, error) {
	return t.Resolver().TransformConfig(config)
}

func (t Type) DefaultValue(config interface{}) (interface{}, error) {
	return t.Resolver().DefaultValue(config)
}
