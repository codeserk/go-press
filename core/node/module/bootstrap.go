package bootstrap

import (
	"press/core/node/http"
	"press/core/node/repository"
	"press/core/node/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, router *mux.Router) {
	nodeRepository := repository.New(mongoClient)
	nodeService := service.New(nodeRepository)

	http.MakeHandlers(router, nodeService)
}
