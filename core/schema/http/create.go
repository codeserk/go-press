package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/schema"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type createFieldRequest struct {
	Name      string      `json:"name" validate:"required"`
	Primitive int         `json:"primitive" validate:"required"`
	Data      interface{} `json:"data"`
} // @name CreateFieldRequest

type createRequest struct {
	Name   string               `json:"name"`
	Fields []createFieldRequest `json:"fields"`
} // @name CreateSchemaRequest

// @Tags Schema
// Creates a new schema
// @Summary Creates a new schema
// @Description Creates a new schema
// @ID create-schema
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "??"
// @Param body body createRequest createRequest "Realm parameters"
// @Success 200
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /v1/realm/{realmId}/schema [post]
func create(s schema.Service) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		realmID := params["realmId"]
		if realmID == "" {
			util.ValidationError(w, errors.New("invalid request, realm is missing"))
		}

		var input createRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			util.ValidationError(w, errors.New("invalid input"))
			return
		}
		currentUser := common.GetUser(r.Context())
		if currentUser == nil {
			util.UnauthorizedError(w)
			return
		}
		err = validate.Struct(input)
		if err != nil {
			util.ValidationError(w, err)
			return
		}

		var fields []*schema.CreateFieldParams
		for _, f := range input.Fields {
			fields = append(fields, &schema.CreateFieldParams{Name: f.Name, Primitive: f.Primitive, Data: f.Data})
		}
		realms, err := s.Create(schema.CreateParams{
			RealmID:  realmID,
			AuthorID: currentUser.ID.Hex(),
			Name:     input.Name,
			Fields:   fields,
		})
		if err != nil {
			util.InternalError(w, err)
		}

		util.SendJSON(w, realms)
	})
}
