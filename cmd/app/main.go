package main

import (
	"github.com/barbibrussa/bookmanager/pkg/models"
	"github.com/barbibrussa/bookmanager/pkg/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := gorm.Open(mysql.Open(server.NewDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database: ", err)
	}

	err = db.AutoMigrate(&models.Book{}, &models.Checkout{}, &models.Review{})
	if err != nil {
		log.Fatal("Error while auto-migrating models: ", err)
	}

	r := chi.NewRouter()
	s := server.NewServer(db)

	user := os.Getenv("APP_USER")
	password := os.Getenv("APP_PASSWORD")

	r.Use(middleware.BasicAuth("librarians", map[string]string{
		user: password,
	}))

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
