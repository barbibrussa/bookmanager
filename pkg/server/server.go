package server

import (
	"encoding/json"
	"fmt"
	"github.com/barbibrussa/bookmanager/pkg/models"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	db *gorm.DB
}

func (s *Server) writeResponse(w http.ResponseWriter, status int, body []byte) {
	_, err := w.Write(body)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
}

func (s *Server) ListBooks(w http.ResponseWriter, r *http.Request) {
	var list []models.Book

	err := s.db.Model(&models.Book{}).Find(&list).Error
	if err != nil {
		http.Error(w, "Failed to list books from database", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(list)
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, http.StatusOK, body)
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	book, err := models.NewBookFromJSON(body)
	if err != nil {
		http.Error(w, "Failed to unmarshall request", http.StatusInternalServerError)
		return
	}

	err = s.db.Model(&models.Book{}).Save(&book).Error
	if err != nil {
		http.Error(w, "Failed to create book to database", http.StatusInternalServerError)
		return
	}

	payload, err := book.ToJSON()
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, http.StatusCreated, payload)
}

func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book models.Book

	err := s.db.Model(&models.Book{}).Where("id = ?", id).First(&book).Error
	if err != nil {
		http.Error(w, "Failed to get book from database", http.StatusNotFound)
		return
	}

	err = s.db.Delete(&models.Book{}, id).Error
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, http.StatusOK, []byte(fmt.Sprintf("Book id %s has been deleted", id)))
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book models.Book

	err := s.db.Model(&models.Book{}).Where("id = ?", id).First(&book).Error
	if err != nil {
		http.Error(w, "Failed to get book from database", http.StatusNotFound)
		return
	}

	body, err := book.ToJSON()
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, http.StatusOK, body)
}

func (s *Server) BorrowBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book models.Book

	err := s.db.Model(&models.Book{}).Where("id = ?", id).First(&book).Error
	if err != nil {
		http.Error(w, "Failed to get book from database", http.StatusNotFound)
		return
	}

	var count int64

	err = s.db.Model(&models.Checkout{}).Where("book_id = ?", id).Where("returned_at IS NULL").Count(&count).Error
	if err != nil {
		http.Error(w, "Error while getting book", http.StatusInternalServerError)
		return
	}

	if count != 0 {
		http.Error(w, "Book not available", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	checkout, err := models.NewCheckoutFromJSON(body)
	if err != nil {
		http.Error(w, "Failed to unmarshall request", http.StatusInternalServerError)
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	checkout.BookID = i
	checkout.BorrowedAt = time.Now()
	checkout.Deadline = time.Now().Add(730 * time.Hour)

	err = s.db.Model(&models.Checkout{}).Save(&checkout).Error
	if err != nil {
		http.Error(w, "Failed to create checkout to database", http.StatusInternalServerError)
		return
	}

	payload, err := checkout.ToJSON()
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, http.StatusCreated, payload)
}

func (s *Server) ReturnBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	now := time.Now()

	var count int64

	err := s.db.Model(&models.Checkout{}).Where("book_id = ?", id).Where("returned_at IS NULL").Count(&count).Error
	if err != nil {
		http.Error(w, "Error while getting book", http.StatusInternalServerError)
		return
	}

	if count != 1 {
		http.Error(w, "This book has not been borrowed", http.StatusNotFound)
		return
	}

	q := s.db.Model(&models.Checkout{}).Where("book_id = ? AND returned_at IS NULL", id).UpdateColumn("returned_at", &now)
	if q.Error != nil && q.RowsAffected != 1 {
		http.Error(w, "Failed to return book", http.StatusNotFound)
		return
	}

	s.writeResponse(w, http.StatusOK, []byte(fmt.Sprintf("The book %s has been returned", id)))
}

func (s *Server) ListCheckouts(w http.ResponseWriter, r *http.Request) {
	var list []models.Checkout

	err := s.db.Model(&models.Checkout{}).Where("returned_at IS NULL").Find(&list).Error
	if err != nil {
		http.Error(w, "Failed to list checkouts from database", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(list)
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, http.StatusOK, body)
}

func (s *Server) ListBooksAvailable(w http.ResponseWriter, r *http.Request) {
	var list []models.Book

	err := s.db.Model(&models.Book{}).
		Joins("LEFT JOIN checkouts ON books.id = checkouts.book_id AND checkouts.returned_at IS NULL").
		Where("checkouts.id IS NULL AND books.deleted_at IS NULL").Find(&list).Error
	if err != nil {
		http.Error(w, "Failed to list checkouts from database", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(list)
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	s.writeResponse(w, http.StatusOK, body)
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}
