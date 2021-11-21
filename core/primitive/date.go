package primitive

type dateConfig struct {
	Required bool   `json:"required" default:"false" validate:"required"`
	Type     string `json:"type" default:"date" validate:"required,oneof=date time dateTime"`
}

type dateResolver struct{}

func (r dateResolver) DefaultConfig() interface{} {
	return dateConfig{
		Required: false,
		Type:     "date",
	}
}

func (r dateResolver) TransformConfig(config Config) (Config, error) {
	return config, nil
}

func (r dateResolver) DefaultValue(config Config) (interface{}, error) {
	value := ""

	return value, nil
}

var date dateResolver = dateResolver{}
