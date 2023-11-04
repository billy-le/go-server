package models

import (
	"github.com/billy-le/go-server/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (book *Book) CreateBook() *Book {
	db := config.GetDB()
	db.Model(book)
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	db := config.GetDB()
	var books []Book
	db.Find(&books)
	return books
}

func GetBook(Id int64) (*Book, *gorm.DB) {
	db := config.GetDB()
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) Book {
	db := config.GetDB()
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
