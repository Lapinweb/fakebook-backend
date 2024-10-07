package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id     uint   `gorm:"primaryKey;autoIncrement" json:"key"`
	Title  string `json:"title"`
	Author string `json:"author_name"`
	Image  string `json:"cover_i"`
}

func (book *Book) GetBooks(db *gorm.DB) (*[]Book, error) {
	var books []Book
	err := db.Model(&Book{}).Find(&books).Error
	return &books, err
}

func (book *Book) GetBookById(db *gorm.DB, id uint) error {
	return db.Model(&Book{}).First(book, id).Error
}

func (book *Book) UpdateOrCreateBook(db *gorm.DB) error {
	return db.Save(book).Error
}

func (book *Book) DeleteBook(db *gorm.DB, id uint) error {
	err := db.Delete(book, id).Error
	return err
}
