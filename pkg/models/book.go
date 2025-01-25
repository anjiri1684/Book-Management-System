package models

import (
	"github.com/anjiri1684/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB
type Book struct {
	gorm.Model
	Name 					string 	`gorm:"json:name"`
	Author 				string 	`json:"author"`
	Publication 	string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book)CreateBook() *Book {
	db.Create(b)
	return b
}

func GetAllBooks()[]Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64)(*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}