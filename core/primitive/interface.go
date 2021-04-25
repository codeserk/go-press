package primitive

type Interface interface {
	Validate(data interface{}) error
}
