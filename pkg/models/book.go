package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b Book) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

func NewBookFromJSON(b []byte) (Book, error) {
	var book Book
	err := json.Unmarshal(b, &book)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}
