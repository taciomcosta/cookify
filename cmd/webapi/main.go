package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taciomcosta/cookify/cmd/webapi/handlers"
	"github.com/taciomcosta/cookify/cmd/webapi/middlewares"
	"github.com/taciomcosta/cookify/internal/config"
)

var serverAddress = config.GetString("SERVER_ADDRESS")
var swaggerURLPath = "/swagger"
var router *mux.Router = mux.NewRouter()

func main() {
	addSwagger(router)
	addHandlers(router)
	addMiddlewares(router)
	serve()
}

func addSwagger(router *mux.Router) {
	fileServer := http.FileServer(http.Dir("./swagger/"))
	handler := http.StripPrefix(swaggerURLPath, fileServer)
	router.PathPrefix(swaggerURLPath).Handler(handler)
}

func addHandlers(router *mux.Router) {
	router.HandleFunc("/recipes", handlers.FindRecipes).Methods("GET")
}

func addMiddlewares(r *mux.Router) {
	r.Use(middlewares.Json)
	r.Use(middlewares.Logging)
}

func serve() {
	log.Printf("Server listening on %s\n", serverAddress)
	http.Handle("/", router)
	err := http.ListenAndServe(serverAddress, nil)
	log.Fatal(err)
}
