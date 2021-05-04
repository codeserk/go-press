package bootstrap

import (
	"press/core/schema"
	"press/core/schema/http"
	"press/core/schema/repository"
	"press/core/schema/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, router *mux.Router) schema.Service {
	schemaRepository := repository.New(mongoClient)
	schemaService := service.New(schemaRepository)

	http.MakeHandlers(router, schemaService)

	return schemaService
}
