package controllers

import (
	. "../../services"
	"../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CalculatorController struct {
}

func (b CalculatorController) Bind(router *mux.Router) {
	router.HandleFunc("/api/calc", b.Calc).Methods("POST")
}

func (b CalculatorController) Calc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var expression models.ExpressionModel
	_ = json.NewDecoder(request.Body).Decode(&expression)

	calculator := Calculator{}

	result := calculator.Calculate(expression.Content)

	json.NewEncoder(writer).Encode(result)
}
