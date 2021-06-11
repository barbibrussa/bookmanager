package main

import (
	"github.com/barbibrussa/bookmanager/pkg/models"
	"github.com/barbibrussa/bookmanager/pkg/server"
	"github.com/go-chi/chi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("book_manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database: ", err)
	}

	err = db.AutoMigrate(&models.Book{}, &models.Checkout{}, &models.Review{})
	if err != nil {
		log.Fatal("Error while auto-migrating models: ", err)
	}

	r := chi.NewRouter()
	s := server.NewServer(db)

	r.Get("/books", s.ListBooks)
	r.Post("/books", s.CreateBook)
	r.Delete("/books/{id}", s.DeleteBook)
	r.Get("/books/{id}", s.GetBook)
	r.Post("/books/{id}/borrow", s.BorrowBook)
	r.Put("/books/{id}/return", s.ReturnBook)
	r.Get("/checkouts", s.ListCheckouts)
	r.Get("/books/available", s.ListBooksAvailable)
	r.Post("/books/{id}/review", s.CreateReview)
	r.Get("/books/{id}/reviews", s.ListBookReviews)
	r.Delete("/review/{id}", s.DeleteReview)

	err = http.ListenAndServe(":3030", r)
	if err != nil {
		log.Fatal("Error serving HTPP on port :3030 ", err)
	}
}
