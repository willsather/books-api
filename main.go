package main

import (
	"log"
	"net/http"

	"github.com/willsather/books-api/internals/books"
)

func main() {
	port := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", books.GetBooksHandler)
	mux.HandleFunc("GET /books/{id}", books.GetBookHandler)
	mux.HandleFunc("POST /books", books.PostBookHandler)
	mux.HandleFunc("DELETE /books/{id}", books.RemoveBookHandler)

	log.Printf("Starting server on http://localhost%s 🚀", port)
	http.ListenAndServe(port, mux)
}
