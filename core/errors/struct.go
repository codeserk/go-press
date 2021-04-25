package errors

import "fmt"

type PressError struct {
	internalMessage string
	publicMessage   string

	code string
}

var DEFAULT_ERROR_MESSAGE = "Something went wrong"
var DEFAULT_CODE = "internal_error"

func New(message string) Interface {
	return &PressError{
		internalMessage: message,
		publicMessage:   DEFAULT_ERROR_MESSAGE,
		code:            DEFAULT_CODE,
	}
}

func Newf(format string, params ...interface{}) Interface {
	return New(fmt.Sprintf(format, params))

}

func Internal(message string) Interface {
	return New(message)
}

func Internalf(format string, params ...interface{}) Interface {
	return Newf(format, params)
}

func Public(message string) Interface {
	return &PressError{
		internalMessage: DEFAULT_ERROR_MESSAGE,
		publicMessage:   message,
		code:            DEFAULT_CODE,
	}
}

func Publicf(format string, params ...interface{}) Interface {
	return Public(fmt.Sprintf(format, params))
}

func (e *PressError) Error() string {
	return e.internalMessage
}

func (e *PressError) Internal() string {
	return e.internalMessage
}

func (e *PressError) Public() string {
	return e.publicMessage
}

func (e *PressError) WithPublic(message string) Interface {
	e.publicMessage = message

	return e
}
func (e *PressError) WithInternal(message string) Interface {
	e.internalMessage = message

	return e
}
func (e *PressError) WithCode(code string) Interface {
	e.code = code

	return e
}
