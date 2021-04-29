package main

import (
	"log"
	"net/http"
	"strconv"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"

	"press/common/jwt"
	fieldModule "press/core/field/module"
	"press/core/http/middleware"
	realmModule "press/core/realm/module"
	schemaModule "press/core/schema/module"
	userModule "press/core/user/module"
	"press/mongo"

	"press/config"
	_ "press/docs"
)

// @title GoPress
// @version 1.0
// @description This is a sample server Petstore server.
// @contact.name API Support
// @license.name Apache 2.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:5050
// @BasePath /
// main Main function.
func main() {
	log.SetPrefix("[GoPress] ")

	router := mux.NewRouter()

	conf, err := config.Parse()
	if err != nil {
		log.Fatalf("error found while trying to load the configuration: %v", err)
	}

	client, err := mongo.Connect(conf.MongoDB.Host)
	if err != nil {
		log.Fatalf("error found while connect to mongodb: %v", err)
	}

	jwtService := jwt.New(conf)

	router.Use(middleware.CorsMiddleware)

	userService := userModule.Bootstrap(client, jwtService, router)
	router.Use(userService.CreateAuthMiddleware)
	realmModule.Bootstrap(client, router)
	schemaModule.Bootstrap(client, router)
	fieldModule.Bootstrap(client, router)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	http.Handle("/", router)

	log.Printf("server started at http://localhost:%v/swagger/", conf.API.Port)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(conf.API.Port), nil))
}
