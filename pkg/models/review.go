package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Score  int    `json:"score"`
	Review string `json:"review"`

	Book   Book `json:"-"`
	BookID int  `json:"book_id"`
}

func (r Review) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func NewReviewFromJSON(b []byte, id int) (Review, error) {
	var review Review
	err := json.Unmarshal(b, &review)
	if err != nil {
		return Review{}, err
	}

	review.BookID = id

	return review, nil
}
