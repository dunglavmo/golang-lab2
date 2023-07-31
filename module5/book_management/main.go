package main

import (
	"fmt"

	"book/db"
	"book/models"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Add books
	book1 := models.Book{Name: "Docker", Author: "dungla"}
	book2 := models.Book{Name: "K8s", Author: "ladung"}
	book3 := models.Book{Name: "AWS", Author: "AWS Provider"}
	db.AddBook(&book1)
	db.AddBook(&book2)
	db.AddBook(&book3)

	// List all books
	books := db.ListBooks()
	fmt.Println("All Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Name: %s, Author: %s, CreatedAt: %s, UpdatedAt: %s\n",
			book.ID, book.Name, book.Author, book.CreatedAt.Format("2006-01-02 15:04:05"), book.UpdatedAt.Format("2006-01-02 15:04:05"))
	}

	// Update a book
	bookUpdate := db.GetBookByID(1)
	if bookUpdate != nil {
		bookUpdate.Name = "DockerV2.2"
		db.UpdateBook(bookUpdate)
	}

	bookUpdate2 := db.GetBookByID(3)
	if bookUpdate2 != nil {
		bookUpdate2.Author = "le anh dung"
		db.UpdateBook(bookUpdate2)
	}

	// Get a specific book
	bookID := 3
	book := db.GetBookByID(bookID)
	fmt.Println("Get specifiec book with ID")
	if book != nil {
		fmt.Printf("Book with ID %d: Name: %s, Author: %s, CreatedAt: %s, UpdatedAt: %s\n",
			book.ID, book.Name, book.Author, book.CreatedAt.Format("2006-01-02 15:04:05"), book.UpdatedAt.Format("2006-01-02 15:04:05"))
	}

	// Delete a book
	db.DeleteBook(bookID)

	// List all books after delete
	books = db.ListBooks()
	fmt.Println("All Books after delete:")
	for _, book := range books {
		fmt.Printf("ID: %d, Name: %s, Author: %s, CreatedAt: %s, UpdatedAt: %s\n",
			book.ID, book.Name, book.Author, book.CreatedAt.Format("2006-01-02 15:04:05"), book.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}
