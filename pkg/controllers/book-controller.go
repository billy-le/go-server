package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/billy-le/go-server/pkg/models"
	"github.com/billy-le/go-server/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	book, _ := models.GetBook(bookId)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	book := models.DeleteBook(bookId)
	_, _ = json.Marshal(book)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}
	book, db := models.GetBook(bookId)
	if updateBook.Title != "" {
		book.Title = updateBook.Title
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(&book)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
