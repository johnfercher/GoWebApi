package main

import (
	"./controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Router
	router := mux.NewRouter()

	// Controllers
	bookController := controllers.BookController{}

	// Binds
	bookController.Bind(router)

	log.Fatal(http.ListenAndServe(":5000", router))
}
