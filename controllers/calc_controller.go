package controllers

import (
	. "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CalcController struct {
}

func (b CalcController) Bind(router *mux.Router) {
	router.HandleFunc("/api/calc", b.Calc).Methods("POST")
}

func (b CalcController) Calc(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var expression Expression
	_ = json.NewDecoder(request.Body).Decode(&expression)

	json.NewEncoder(writer).Encode(expression)
}
