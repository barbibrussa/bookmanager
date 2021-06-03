package models

import (
	"encoding/json"
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

func (c Checkout) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

func NewCheckoutFromJSON(b []byte) (Checkout, error) {
	var checkout Checkout
	err := json.Unmarshal(b, &checkout)
	if err != nil {
		return Checkout{}, err
	}
	return checkout, nil
}
