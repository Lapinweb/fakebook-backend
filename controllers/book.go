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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Erreur récupération du paramètre")
		return
	}

	var book models.Book
	err = book.GetBookById(repository.Db, uint(id))
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
		Image:  "https://via.assets.so/img.jpg?w=250&h=150&tc=white&bg=pink",
	}
	err := newBook.UpdateOrCreateBook(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur création du livre")
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": newBook})
}

func (repository *BookRepo) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Erreur récupération du paramètre")
		return
	}

	var bookInput models.Book
	err = bookInput.GetBookById(repository.Db, uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Le livre n'existe pas")
		return
	}

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération du JSON")
		return
	}

	updatedBook := models.Book{
		ID:     uint(id),
		Title:  bookInput.Title,
		Author: bookInput.Author,
		Image:  "https://via.assets.so/img.jpg?w=250&h=150&tc=white&bg=pink",
	}
	err = updatedBook.UpdateOrCreateBook(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur création du livre")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedBook})
}

func (repository *BookRepo) DeleteBooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Erreur récupération du paramètre")
		return
	}

	var bookFind models.Book
	err = bookFind.GetBookById(repository.Db, uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Le livre n'existe pas")
		return
	}

	err = bookFind.DeleteBook(repository.Db, uint(bookFind.ID))
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur suppression du livre")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted!"})
}
