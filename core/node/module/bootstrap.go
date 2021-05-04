package bootstrap

import (
	"press/core/node/http"
	"press/core/node/repository"
	"press/core/node/service"
	"press/core/schema"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, router *mux.Router, schemas schema.Service) {
	nodeRepository := repository.New(mongoClient)
	nodeService := service.New(nodeRepository, schemas)

	http.MakeHandlers(router, nodeService)
}
