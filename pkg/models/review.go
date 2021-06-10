package models

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (r Review) Validate() error {
	if r.Score < 0 || r.Score > 10 {
		return fmt.Errorf("invalid score: %d [1-10]", r.Score)
	}
	if len(r.Review) == 0 {
		return errors.New("empty review")
	}
	return nil
}

func NewReviewFromJSON(id int, b []byte) (Review, error) {
	var review Review
	err := json.Unmarshal(b, &review)
	if err != nil {
		return Review{}, err
	}

	review.BookID = id

	if err = review.Validate(); err != nil {
		return Review{}, err
	}

	return review, nil
}
