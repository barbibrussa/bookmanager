package main

import (
	"github.com/barbibrussa/bookmanager/pkg/models"
	"github.com/barbibrussa/bookmanager/pkg/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := gorm.Open(sqlite.Open("book_manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database: ", err)
		return
	}

	err = db.AutoMigrate(&models.Book{}, &models.Checkout{}, &models.Review{})
	if err != nil {
		log.Fatal("Error while auto-migrating models: ", err)
		return
	}

	r := chi.NewRouter()
	s := server.NewServer(db)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	user := os.Getenv("APP_USER")
	password := os.Getenv("APP_PASSWORD")

	if len(user) == 0 || len(password) == 0 {
		log.Fatalf("Error: empty user and password credentials")
		return
	}

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
		return
	}
}
