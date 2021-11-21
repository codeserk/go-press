package primitive

type nodeConfig struct {
	Required bool      `json:"required" bson:"required" default:"false" validate:"required"`
	Min      *int8     `json:"min" bson:"max" validate:"number"`
	Max      *int8     `json:"min" bson:"max" validate:"number"`
	Schemas  *[]string `json:"schemas" bson:"schemas" validate:"number"`
}

type nodeResolver struct{}

func (r nodeResolver) DefaultConfig() interface{} {
	return nodeConfig{
		Required: false,
		Min:      nil,
		Max:      nil,
		Schemas:  nil,
	}
}

func (r nodeResolver) TransformConfig(config Config) (Config, error) {
	schemas, exists := config["schemas"]
	if exists {
		if a, ok := schemas.([]interface{}); ok && len(a) == 0 {
			delete(config, "schemas")
		}
	}

	return config, nil
}

func (r nodeResolver) DefaultValue(config Config) (interface{}, error) {
	value := ""

	return value, nil
}

var node nodeResolver = nodeResolver{}
