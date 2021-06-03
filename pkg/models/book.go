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
