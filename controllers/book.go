package controllers

import (
	"fakebook-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var Library []models.Book
var Counter int

func InitDatabase() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "fakebook",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	/*imageUrl := []string{
		"https://via.assets.so/img.jpg?w=400&h=650&tc=white&bg=pink",
		"https://via.assets.so/img.jpg?w=220&h=250&tc=white&bg=pink",
		"https://via.assets.so/img.jpg?w=100&h=100&tc=white&bg=pink",
		"https://via.assets.so/img.jpg?w=300&h=250&tc=white&bg=pink",
	}*/

	insertedBooks := "INSERT INTO books (title, author, image) VALUES "
	for i := 0; i < 10; i++ {
		book := models.Book{
			Title:  fmt.Sprintf("Book %v", i+1),
			Author: fmt.Sprintf("Author %v", i+1),
			Image:  "https://via.assets.so/img.jpg?w=250&h=150&tc=white&bg=pink",
		}
		//Library = append(Library, book)

		insertedBooks = insertedBooks + fmt.Sprintf("('%v', '%v', '%v')", book.Title, book.Author, book.Image)
		if i != 9 {
			insertedBooks = insertedBooks + ", "
		}
	}
	fmt.Print(insertedBooks)
	insert, err := db.Query(insertedBooks + ";")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func FindBooks(c *gin.Context) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "fakebook",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}

	books := []models.Book{}
	for results.Next() {
		var book models.Book
		err = results.Scan(&book.Id, &book.Title, &book.Author, &book.Image)
		if err != nil {
			panic(err.Error())
		}

		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
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

func FindBookById(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"data": bookFind})
}

func UpdateBook(c *gin.Context) {
	var updatedBook models.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, book := range Library {
		if c.Param("id") == strconv.Itoa(book.Id) {
			if updatedBook.Title != "" {
				Library[index].Title = updatedBook.Title
			}
			if len(updatedBook.Author) != 0 {
				Library[index].Author = updatedBook.Author
			}
			c.JSON(http.StatusOK, gin.H{"data": updatedBook})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
}

func removeIt(book models.Book, bookSlice []models.Book) []models.Book {
	for index, bookToDelete := range bookSlice {
		if bookToDelete.Id == book.Id {
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
