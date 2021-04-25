package main

import (
	"log"
	"net/http"
	"strconv"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"

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

// @host localhost:5050
// @BasePath /
func main() {
	log.SetPrefix("[GoPress] ")

	router := mux.NewRouter()

	conf, err := config.Parse()
	if err != nil {
		log.Fatalf("There was an error found while trying to load the configuration: %v", err)
	}

	client, err := mongo.Connect(conf.MongoDB.Host)
	if err != nil {
		log.Fatalf("There was an error found while connect to mongodb: %v", err)
	}

	userModule.Bootstrap(client, router)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	http.Handle("/", router)

	log.Printf("Server started at http://localhost:%v/swagger/", conf.Api.Port)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(conf.Api.Port), nil))
}
