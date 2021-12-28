package primitive

type viewConfig struct {
	// Whether the field is required or not.
	Required bool `json:"required" bson:"required" default:"false" validate:"required"`

	// Min views.
	Min *int8 `json:"min" bson:"max" validate:"number"`

	// Max number of views.
	Max *int8 `json:"min" bson:"max" validate:"number"`

	// View schemas that are accepted. Any schema if not present.
	Schemas *[]string `json:"schemas" bson:"schemas" validate:"number"`
}

type viewResolver struct{}

func (r viewResolver) DefaultConfig() interface{} {
	return viewConfig{
		Required: false,
		Min:      nil,
		Max:      nil,
		Schemas:  nil,
	}
}

func (r viewResolver) TransformConfig(config Config) (Config, error) {
	schemas, exists := config["schemas"]
	if exists {
		if a, ok := schemas.([]interface{}); ok && len(a) == 0 {
			delete(config, "schemas")
		}
	}

	return config, nil
}

func (r viewResolver) DefaultValue(config Config) (interface{}, error) {
	value := ""

	return value, nil
}

var view viewResolver = viewResolver{}
