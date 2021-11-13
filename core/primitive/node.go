package primitive

import "errors"

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

func (r nodeResolver) TransformConfig(config interface{}) (interface{}, error) {
	c, ok := config.(map[string]interface{})
	if !ok {
		return config, errors.New("invalid config")
	}

	schemas, exists := c["schemas"]
	if exists {
		if a, ok := schemas.([]interface{}); ok && len(a) == 0 {
			delete(c, "schemas")
		}
	}

	return c, nil
}

func (r nodeResolver) DefaultValue(config interface{}) (interface{}, error) {
	value := ""

	return value, nil
}

var node nodeResolver = nodeResolver{}
