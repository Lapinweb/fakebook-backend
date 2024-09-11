package controllers

import (
	"fakebook-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Library []models.Book
var Counter int

func InitDatabase() {
	Counter = 1

	for i := 0; i < 10; i++ {
		book := models.Book{
			Id:     Counter,
			Title:  fmt.Sprintf("Book %v", Counter),
			Author: fmt.Sprintf("Author %v", Counter),
			Image:  "",
		}
		Library = append(Library, book)
		Counter++
	}
}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Library})
}

func CreateBooks(c *gin.Context) {
	var input models.Book

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Counter++

	book := models.Book{Id: Counter, Title: input.Title, Author: input.Author, Image: input.Image}
	Library = append(Library, book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func removeIt(book models.Book, bookSlice []models.Book) []models.Book {
	for index, bookToDelete := range bookSlice {
		if bookToDelete == book {
			return append(bookSlice[0:index], bookSlice[index+1:]...)
		}
	}
	return bookSlice
}

func DeleteBooks(c *gin.Context) {
	bookFound := false

	var bookFind models.Book
	for _, book := range Library {
		if c.Param("id") == strconv.Itoa(book.Id) {
			bookFound = true
			bookFind = book
		}
	}

	if !bookFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	Library = removeIt(bookFind, Library)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
