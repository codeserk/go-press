package bootstrap

import (
	"press/common/jwt"
	"press/core/user/http"
	"press/core/user/repository"
	"press/core/user/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Bootstrap(mongoClient *mongo.Client, jwtService jwt.Interface, router *mux.Router) service.Interface {
	userRepository := repository.New(mongoClient)
	userService := service.New(userRepository, jwtService)

	http.MakeHandlers(router, userService)

	return userService
}
