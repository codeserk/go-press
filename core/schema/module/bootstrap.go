package bootstrap

import (
	"press/core/schema/http"
	"press/core/schema/repository"
	"press/core/schema/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, router *mux.Router) {
	schemaRepository := repository.New(mongoClient)
	schemaService := service.New(schemaRepository)

	http.MakeHandlers(router, schemaService)
}
