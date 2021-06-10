package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Score  int    `json:"score"`
	Review string `json:"review"`

	Book   Book `json:"-"`
	BookID int  `json:"book_id"`
}
