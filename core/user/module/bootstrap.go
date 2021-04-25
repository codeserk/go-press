package bootstrap

import (
	"press/core/user/http"
	"press/core/user/repository"
	"press/core/user/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, router *mux.Router) {
	repository := repository.Create(mongoClient)
	service := service.Create(repository)

	http.MakeHandlers(router, service)
}
