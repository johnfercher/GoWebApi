package main

import (
	. "./models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct (Model)
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
	books = append(books, Book{Id: "2", Isbn: "1266", Title: "Test2", Author: &author})

	// Handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(&Book{})
}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.Id == params["id"] {
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.Id = params["id"]
			books[index] = book
			json.NewEncoder(writer).Encode(book)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Book{})
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Book{})
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}
