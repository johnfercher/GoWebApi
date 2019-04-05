package main

import (
	"./application/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Router
	router := mux.NewRouter()

	// Controllers
	calculatorController := controllers.CalculatorController{}

	// Binds
	calculatorController.Bind(router)

	log.Fatal(http.ListenAndServe(":5000", router))
}
