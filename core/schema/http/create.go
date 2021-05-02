package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"press/common"
	"press/common/util"
	"press/core/node"
	"press/core/schema"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type createSchemaRequest struct {
	Type node.Type `json:"type" enums:"scene,nested"`
	Name string    `json:"name"`
} // @name CreateSchemaRequest

// @Tags Schema
// Creates a new schema
// @Summary Creates a new schema
// @Description Creates a new schema
// @ID create-schema
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param realmId path string string "Realm ID"
// @Param body body createSchemaRequest createSchemaRequest "Schema parameters"
// @Success 200 {object} schema.Entity
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

		var input createSchemaRequest
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

		result, err := s.Create(schema.CreateParams{
			RealmID:  realmID,
			AuthorID: currentUser.ID.Hex(),
			Type:     input.Type,
			Name:     input.Name,
		})
		log.Print(result, err)
		if err != nil {
			util.InternalError(w, err)
			return
		}

		util.SendJSON(w, result)
	})
}
