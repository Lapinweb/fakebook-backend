package database

import (
	"fmt"

	"fakebook-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "fakebook"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func ConnectDb() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

func InitDatabase(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		book := models.Book{
			Title:  fmt.Sprintf("Book %v", i+1),
			Author: fmt.Sprintf("Author %v", i+1),
			Image:  "https://via.assets.so/img.jpg?w=250&h=150&tc=white&bg=pink",
		}

		db.Save(&book)
	}
}
