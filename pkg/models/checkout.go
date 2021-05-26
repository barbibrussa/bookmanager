package models

import (
	"gorm.io/gorm"
	"time"
)

type Checkout struct {
	gorm.Model
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DNI         string     `json:"dni"`
	PhoneNumber string     `json:"phone_number"`
	BorrowedAt  time.Time  `json:"borrowed_at"`
	ReturnedAt  *time.Time `json:"returned_at"`
	Deadline    time.Time  `json:"deadline"`

	Book   Book `json:"-"`
	BookID int  `json:"book_id"`
}
