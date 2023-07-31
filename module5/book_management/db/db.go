package db

import (
	"book/models"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Book{})
}

func AddBook(book *models.Book) {
	db.Create(book)
}

func ListBooks() []models.Book {
	var books []models.Book
	db.Find(&books)
	return books
}

func UpdateBook(book *models.Book) {
	db.Save(book)
}

func GetBookByID(id int) *models.Book {
	var book models.Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil
	}
	return &book
}

func DeleteBook(id int) {
	db.Delete(&models.Book{}, id)
}
