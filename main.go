package main

import (
	"encoding/json"
	"log"
	"net/http"

	"press/mongo"
	"press/schema"
	"press/schema/field"
	"press/schema/repository"

	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Print("Request at home")
	json.NewEncoder(w).Encode("hello")
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.

// @contact.name API Support
// @license.name Apache 2.0

// @host localhost:8050
// @BasePath /
func main() {
	log.SetPrefix("training> ")

	router := mux.NewRouter()

	// conf, err := config.Parse()
	// if err != nil {
	// 	log.Fatalf("There was an error found while trying to load the configuration: \n%s \n", err.Error())
	// }

	// cats := cat.NewLocalService()
	// catHttp.MakeCatHandlers(router, cats)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	http.Handle("/", router)

	client, err := mongo.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	context, cancel := mongo.CreateContext()
	defer cancel()

	var field field.SavedEntity
	// field.ID = primitive.NewObjectID()
	field.Name = "test name"
	field.Primitive = 5

	var newSchema schema.Entity
	newSchema.Realm = "realm"
	newSchema.Fields = append(newSchema.Fields, field)

	_, err = mongo.Schemas.DeleteMany(context, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var schemas = repository.Create(client)
	schemas.Create(context, newSchema)

	schemasFound, err := schemas.Find(context)

	log.Print(schemasFound[0], len(schemasFound))

	// log.Printf("Server started at http://localhost:%v/swagger/", conf.Api.Port)
	// log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(conf.Api.Port), nil))
}
