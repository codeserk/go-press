package bootstrap

import (
	"press/common/jwt"
	"press/core/user/http"
	"press/core/user/repository"
	"press/core/user/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, jwt jwt.Interface, router *mux.Router) service.Interface {
	repository := repository.Create(mongoClient)
	service := service.New(repository, jwt)

	http.MakeHandlers(router, service)

	return service
}
