package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

func main() {
	// Router
	router := mux.NewRouter()

	// Mock Data
	author := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	books = append(books, Book{Id: "1", Isbn: "1234", Title: "Test", Author: &author})

	// Handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {

}

func updateBook(writer http.ResponseWriter, request *http.Request) {

}

func createBook(writer http.ResponseWriter, request *http.Request) {

}

func getBook(writer http.ResponseWriter, request *http.Request) {

}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}
