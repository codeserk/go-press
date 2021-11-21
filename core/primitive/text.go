package primitive

type textConfig struct {
	Required bool  `json:"required" default:"false" validate:"required"`
	MinChars *int8 `json:"minChars" bson:"minChars" validate:"number"`
	MaxChars *int8 `json:"maxChars" bson:"maxChars" validate:"number"`
}

type textResolver struct{}

func (r textResolver) DefaultConfig() interface{} {
	return textConfig{
		Required: false,
		MinChars: nil,
		MaxChars: nil,
	}
}

func (r textResolver) TransformConfig(config Config) (Config, error) {
	return config, nil
}

func (r textResolver) DefaultValue(config Config) (interface{}, error) {
	value := ""

	return value, nil
}

var text textResolver = textResolver{}
