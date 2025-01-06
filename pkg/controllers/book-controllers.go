package controllers

import (
	m "github.com/anmol/go-bookstore/pkg/models"
	u "github.com/anmol/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var NewBook m.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := m.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //Always set writeheader before write
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"] // /books/{id}  id key defined in routes.
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := m.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &m.Book{}
	u.ParseBody(r, createBook)
	// This one is also correct but above one is better as dealing with a pointer throughout code enhance
	// more readability
	// createBook := m.Book{}
	// u.ParseBody(r, &createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deleteBook := m.DeleteBook(ID)
	res, _ := json.Marshal(deleteBook)
	m.DeleteBook(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Book successfully deleted"}
	json.NewEncoder(w).Encode(response)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &m.Book{}
	u.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	getBookDetails, db := m.GetBookById(ID)
	if updateBook.Name != "" {
		getBookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		getBookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		getBookDetails.Publication = updateBook.Publication
	}
	db.Save((&getBookDetails))
	res, _ := json.Marshal(getBookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
