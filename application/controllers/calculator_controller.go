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

	var expressionRequest models.ExpressionRequest
	_ = json.NewDecoder(request.Body).Decode(&expressionRequest)

	calculator := Calculator{}

	expressionResult := calculator.Calculate(expressionRequest.Expression)

	var expressionResponse models.ExpressionResponse
	expressionResponse.Resolution = expressionResult.Resolution

	json.NewEncoder(writer).Encode(expressionResponse)
}
