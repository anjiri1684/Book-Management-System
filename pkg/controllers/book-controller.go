package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anjiri1684/go-bookstore/pkg/models"
	"github.com/anjiri1684/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book


func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()

	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func UpdateBook(w http.ResponseWriter, r *http.Request){}

func DeleteBook(w http.ResponseWriter, r *http.Request){}
