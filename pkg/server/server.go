package server

import (
	"encoding/json"
	"fmt"
	"github.com/barbibrussa/bookmanager/pkg/models"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type Server struct {
	db *gorm.DB
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

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var book models.Book

	err = json.Unmarshal(body, &book)
	if err != nil {
		http.Error(w, "Failed to unmarshall request", http.StatusInternalServerError)
		return
	}

	err = s.db.Model(&models.Book{}).Save(&book).Error
	if err != nil {
		http.Error(w, "Failed to create book to database", http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(payload)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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

	_, err = w.Write([]byte(fmt.Sprintf("Book id %s has been deleted", id)))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book models.Book

	err := s.db.Model(&models.Book{}).Where("id = ?", id).First(&book).Error
	if err != nil {
		http.Error(w, "Failed to get book from database", http.StatusNotFound)
		return
	}

	body, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshall response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}
