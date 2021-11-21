package primitive

type booleanConfig struct {
	Required bool `json:"required" default:"false" validate:"required"`
}

type booleanResolver struct{}

func (r booleanResolver) DefaultConfig() interface{} {
	return booleanConfig{
		Required: false,
	}
}

func (r booleanResolver) TransformConfig(config Config) (Config, error) {
	return config, nil
}

func (r booleanResolver) DefaultValue(config Config) (interface{}, error) {
	value := false

	return value, nil
}

var boolean booleanResolver = booleanResolver{}
