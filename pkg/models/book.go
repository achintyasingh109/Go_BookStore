package model

import (
	"github.com/anmol/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	//gorm.Model is a predefined struct which includes feilds to manage database records
	/*
	   type Model struct {
	       ID        uint             `gorm:"primaryKey"`
	       CreatedAt time.Time
	       UpdatedAt time.Time
	       DeletedAt gorm.DeletedAt   `gorm:"index"` (Supports soft deletion by marking records as
	   	deleted without removing them from the database.)
	   }
	*/
	gorm.Model // So whenever, suppose we added a new book it will automatically define its id
	// which is primary key by default and all other factors also.

	Name string `gorm:"" json:"name"`
	/*

		gorm:"" it is the directives provided by gorm which are string literals enclosed in backticks ``
		it provides information about  the desired behavior of each field when interfacing with the database.
		gorm:"column:book_name"
		gorm:"type:varchar(100)"
		gorm:"primaryKey" , etc

		json:"name": This tag directs the encoding/json package to map the Name field to the JSON key "name" during serialization and deserialization.

	*/
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) //During production we use Atlas to have more control like versioning
	// control and rollback

	/*
	   Auto Migrate performs :-

	   Table Creation: If the books table doesn't exist in your database, AutoMigrate will
	   create it based on the Book struct definition.

	   Column Addition: If the books table exists but lacks some columns present in the Book
	   struct,AutoMigrate will add the missing columns.

	   Index Creation: It will create any missing indexes defined in your model's struct tags.

	   No Column Deletion: It won't remove columns from the table, even if they are no longer
	   present in your struct.

	   No Column Type Modification: It won't change the data type of existing columns to match
	   your struct.

	   No Renaming: It won't rename columns or tables.
	*/

}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
