package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	// MOCK DATA // TODO: implement DB
	books = append(books, Book{ID: "1", Isbn: "1234-1234-1234", Title: "This is Da Book", Author: &Author{Firstname: "John", Lastname: "Smith"}})
	books = append(books, Book{ID: "2", Isbn: "9876-9876-9876", Title: "How to work in a office", Author: &Author{Firstname: "Sally", Lastname: "Smith"}})
	// ROUTE HANDLERS / EndPoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//LISTERN
	log.Fatal(http.ListenAndServe(":8080", r))
}
