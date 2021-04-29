package bootstrap

import (
	"press/core/field/http"
	"press/core/field/repository"
	"press/core/field/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, router *mux.Router) {
	fieldRepository := repository.New(mongoClient)
	fieldService := service.New(fieldRepository)

	http.MakeHandlers(router, fieldService)
}
