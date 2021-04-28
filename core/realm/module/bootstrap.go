package bootstrap

import (
	"press/core/realm/http"
	"press/core/realm/repository"
	"press/core/realm/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, router *mux.Router) service.Interface {
	realmRepository := repository.New(mongoClient)
	realmService := service.New(realmRepository)

	http.MakeHandlers(router, realmService)

	return realmService
}
