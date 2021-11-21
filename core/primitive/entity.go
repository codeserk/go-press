package primitive

type Type string

const (
	Text    Type = "text"
	Boolean Type = "boolean"
	Number  Type = "number"
	Date    Type = "date"
	Options Type = "options"
	Node    Type = "node"
	View    Type = "view"
)

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name" example:"Ryuk"`
}

type Config map[string]interface{}

type Resolver interface {
	DefaultConfig() interface{}
	TransformConfig(config Config) (Config, error)
	DefaultValue(config Config) (interface{}, error)
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
	case View:
		return view
	default:
		return text
	}
}

func (t Type) DefaultConfig() interface{} {
	return t.Resolver().DefaultConfig()
}

func (t Type) TransformConfig(config Config) (Config, error) {
	return t.Resolver().TransformConfig(config)
}

func (t Type) DefaultValue(config Config) (interface{}, error) {
	return t.Resolver().DefaultValue(config)
}
