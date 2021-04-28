package util

import (
	"encoding/json"
	"net/http"
	"press/common/errors"
)

type HTTPError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func SendError(w http.ResponseWriter, httpCode int, err error, code string) {
	errorStruct := HTTPError{
		Message: err.Error(),
		Code:    code,
	}
	if e, ok := err.(*errors.PressError); ok {
		errorStruct.Message = e.Public()
	}

	w.WriteHeader(httpCode)
	w.Header().Add("content-type", "application/json")
	// nolint
	json.NewEncoder(w).Encode(errorStruct)
}

func SendJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Add("content-type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		InternalError(w, err)
	}
}

func InternalError(w http.ResponseWriter, err error) {
	SendError(w, http.StatusInternalServerError, err, "Internal Error")
}

func UnauthorizedError(w http.ResponseWriter) {
	SendError(w, http.StatusUnauthorized, errors.New("not authorized"), "Not authorized")
}

func ValidationError(w http.ResponseWriter, err error) {
	SendError(w, http.StatusUnprocessableEntity, err, "Validation Error")
}
