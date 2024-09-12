package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var books = map[string]Book{
	"1": {ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: 2015},
	"2": {ID: "2", Title: "Learning Go", Author: "Jon Bodner", Year: 2021},
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, exists := books[params["id"]]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}
