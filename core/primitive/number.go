package primitive

type numberConfig struct {
	Required bool  `json:"required" default:"false" validate:"required"`
	Min      *int8 `json:"min" bson:"min" validate:"number"`
	Max      *int8 `json:"max" bson:"max" validate:"number"`
}

type numberResolver struct{}

func (r numberResolver) DefaultConfig() interface{} {
	return numberConfig{
		Required: false,
		Min:      nil,
		Max:      nil,
	}
}

func (r numberResolver) DefaultValue(config interface{}) (interface{}, error) {
	value := 0

	return value, nil
}

var number numberResolver = numberResolver{}
