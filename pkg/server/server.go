package server

import (
	"encoding/json"
	"github.com/barbibrussa/bookmanager/pkg/models"
	"gorm.io/gorm"
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

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}
