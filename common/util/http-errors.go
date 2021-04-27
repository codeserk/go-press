package util

import (
	"encoding/json"
	"net/http"
	"press/core/errors"
)

type HttpError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func SendError(w http.ResponseWriter, httpCode int, err error, code string) {
	error := HttpError{
		Message: err.Error(),
		Code:    code,
	}
	if e, ok := err.(*errors.PressError); ok {
		error.Message = e.Public()
	}

	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(error)
}

func InternalError(w http.ResponseWriter, err error) {
	SendError(w, http.StatusInternalServerError, err, "Internal Error")
}

func UnauthorizedError(w http.ResponseWriter) {
	SendError(w, http.StatusUnauthorized, errors.New("Not authorized"), "Not authorized")
}

func ValidationError(w http.ResponseWriter, err error) {
	SendError(w, http.StatusUnprocessableEntity, err, "Validation Error")
}
