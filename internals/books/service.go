package books

import (
	"math/rand"
	"sync"
)

var (
	books = map[int]Book{
		1: {ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", ISBN: "978-0134190440"},
		2: {ID: 2, Title: "Clean Code", Author: "Robert C. Martin", ISBN: "978-0132350884"},
		3: {ID: 3, Title: "The Pragmatic Programmer", Author: "Andrew Hunt", ISBN: "978-0201616224"},
		4: {ID: 4, Title: "You Don't Know JS", Author: "Kyle Simpson", ISBN: "978-1491904244"},
	}

	lock sync.RWMutex
)

// GetAllBooks returns all books
func GetAllBooks() map[int]Book {
	lock.RLock()
	defer lock.RUnlock()
	return books
}

// GetBookByID fetches a book by its ID
func GetBookByID(id int) (Book, bool) {
	lock.RLock()
	defer lock.RUnlock()
	book, found := books[id]
	return book, found
}

// AddBook adds a new book to the collection
func AddBook(book Book) Book {
	book.ID = rand.Intn(1024) // Generate random ID

	lock.Lock()
	books[book.ID] = book
	lock.Unlock()

	return book
}

// RemoveBook deletes a book by its ID
func RemoveBook(id int) bool {
	lock.Lock()
	defer lock.Unlock()
	_, found := books[id]
	if found {
		delete(books, id)
	}
	return found
}
