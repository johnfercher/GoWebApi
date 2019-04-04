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
	calcController := controllers.CalcController{}

	// Binds
	calcController.Bind(router)

	log.Fatal(http.ListenAndServe(":5000", router))
}
