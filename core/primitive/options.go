package primitive

type option struct {
	Value string `json:"value" validate:"required"`
	Label string `json:"label" validate:"required"`
}

type optionsConfig struct {
	Required bool     `json:"required" default:"false" validate:"required"`
	Multiple bool     `json:"multiple" default:"false" validate:"required"`
	Options  []option `json:"options" default:"[]" validate:"required"`
}

type optionsResolver struct{}

func (r optionsResolver) DefaultConfig() interface{} {
	return optionsConfig{
		Required: false,
		Multiple: false,
		Options:  make([]option, 0),
	}
}

func (r optionsResolver) TransformConfig(config interface{}) (interface{}, error) {
	return config, nil
}

func (r optionsResolver) DefaultValue(config interface{}) (interface{}, error) {
	value := ""

	return value, nil
}

var options optionsResolver = optionsResolver{}
