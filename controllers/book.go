package controllers

import (
	"fakebook-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookRepo struct {
	Db *gorm.DB
}

func (repository *BookRepo) FindBooks(c *gin.Context) {
	var bookModel models.Book
	books, err := bookModel.GetBooks(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Books not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (repository *BookRepo) FindBookById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var book models.Book
	err := book.GetBookById(repository.Db, uint(id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Book not found!")
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": &book})
}

type BookCreate struct {
	Title  string `json:"title"`
	Author string `json:"author_name"`
	Image  string `json:"cover_i"`
}

func (repository *BookRepo) CreateBook(c *gin.Context) {
	var bookInput BookCreate
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération du JSON")
		return
	}

	newBook := models.Book{
		Title:  bookInput.Title,
		Author: bookInput.Author,
	}
	err := newBook.UpdateOrCreateBook(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur création du livre")
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": newBook})
}

func (repository *BookRepo) UpdateBook(c *gin.Context) {
	var updatedBook models.Book

	c.JSON(http.StatusOK, gin.H{"data": updatedBook})
}

func (repository *BookRepo) DeleteBooks(c *gin.Context) {

	//id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted!"})
}
