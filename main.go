package main

import (
	"fakebook-api/controllers"
	"fakebook-api/database"
	"fakebook-api/models"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDb()
	db.AutoMigrate(&models.Book{})
	//database.InitDatabase(db)

	bookRepo := controllers.BookRepo{
		Db: db,
	}

	//controllers.InitDatabase()
	r := setupRouter()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))

	r.GET("/books", bookRepo.FindBooks)
	r.GET("/books/:id", bookRepo.FindBookById)
	r.POST("/books", bookRepo.CreateBook)
	r.PUT("/books/:id", bookRepo.UpdateBook)
	r.DELETE("/books/:id", bookRepo.DeleteBooks)

	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	return r
}
