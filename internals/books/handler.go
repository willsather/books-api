package books

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetBooksHandler handles GET /books
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := GetAllBooks()
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Error encountered", http.StatusInternalServerError)
	}
}

// GetBookHandler handles GET /books/{id}
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	book, found := GetBookByID(id)
	if !found {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "Error encountered", http.StatusInternalServerError)
	}
}

// PostBookHandler handles POST /books/create
func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid book data", http.StatusBadRequest)
		return
	}

	newBook := AddBook(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newBook); err != nil {
		http.Error(w, "Error encountered", http.StatusInternalServerError)
	}
}

// RemoveBookHandler handles DELETE /books/{id}
func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Remove the book by ID
	if !RemoveBook(id) {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content, successfully deleted
}
