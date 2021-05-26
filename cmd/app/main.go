package main

import (
	"github.com/barbibrussa/bookmanager/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("book_manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database: ", err)
	}

	err = db.AutoMigrate(&models.Book{}, &models.Checkout{})
	if err != nil {
		log.Fatal("Error while auto-migrating models: ", err)
	}
}
