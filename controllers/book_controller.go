package controllers

import (
	. "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type BookController struct {
}

func (b BookController) Bind(router *mux.Router) {
	router.HandleFunc("/api/books", b.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", b.GetBook).Methods("GET")
	router.HandleFunc("/api/books", b.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", b.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", b.DeleteBook).Methods("DELETE")
}

func (b BookController) DeleteBook(writer http.ResponseWriter, request *http.Request) {
	var books []Book
	author := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	books = append(books, Book{Id: "1", Isbn: "1234", Title: "Test", Author: &author})
	books = append(books, Book{Id: "2", Isbn: "1266", Title: "Test2", Author: &author})

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

func (b BookController) UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var books []Book
	author := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	books = append(books, Book{Id: "1", Isbn: "1234", Title: "Test", Author: &author})
	books = append(books, Book{Id: "2", Isbn: "1266", Title: "Test2", Author: &author})

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

func (b BookController) CreateBook(writer http.ResponseWriter, request *http.Request) {
	var books []Book
	author := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	books = append(books, Book{Id: "1", Isbn: "1234", Title: "Test", Author: &author})
	books = append(books, Book{Id: "2", Isbn: "1266", Title: "Test2", Author: &author})

	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func (b BookController) GetBook(writer http.ResponseWriter, request *http.Request) {
	var books []Book
	author := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	books = append(books, Book{Id: "1", Isbn: "1234", Title: "Test", Author: &author})
	books = append(books, Book{Id: "2", Isbn: "1266", Title: "Test2", Author: &author})

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

func (c BookController) GetBooks(writer http.ResponseWriter, request *http.Request) {
	var books []Book
	author := Author{
		Firstname: "John",
		Lastname:  "Doe",
	}

	books = append(books, Book{Id: "1", Isbn: "1234", Title: "Test", Author: &author})
	books = append(books, Book{Id: "2", Isbn: "1266", Title: "Test2", Author: &author})

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}
