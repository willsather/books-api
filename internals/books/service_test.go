package books

import (
	"testing"
)

func TestAddBook(t *testing.T) {
	// Setup
	newBook := Book{
		Title:  "Go in Action",
		Author: "William Kennedy",
		ISBN:   "978-1617291784",
	}

	// Add the book
	addedBook := AddBook(newBook)

	// Check if the book was added correctly
	if addedBook.Title != newBook.Title || addedBook.Author != newBook.Author || addedBook.ISBN != newBook.ISBN {
		t.Errorf("Expected book to be added with Title: %s, Author: %s, ISBN: %s, but got Title: %s, Author: %s, ISBN: %s",
			newBook.Title, newBook.Author, newBook.ISBN, addedBook.Title, addedBook.Author, addedBook.ISBN)
	}

	// Ensure the book exists in the map after adding
	_, found := GetBookByID(addedBook.ID)
	if !found {
		t.Errorf("Expected to find the book with ID: %d, but it was not found", addedBook.ID)
	}
}

func TestRemoveBook(t *testing.T) {
	// Setup: adding a book to be removed
	bookToRemove := Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		ISBN:   "978-0134190440",
	}
	addedBook := AddBook(bookToRemove)

	// Remove the book
	removed := RemoveBook(addedBook.ID)

	// Check if the book was successfully removed
	if !removed {
		t.Errorf("Expected to remove the book with ID: %d, but it was not removed", addedBook.ID)
	}

	// Verify the book no longer exists
	_, found := GetBookByID(addedBook.ID)
	if found {
		t.Errorf("Expected to not find the book with ID: %d, but it was still found", addedBook.ID)
	}
}
