package dto

import (
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
)

type CreateReviewDTO struct {
	Text   string `json:"text" binding:"required"`
	Rating int64  `json:"rating" binding:"required"`
}

type ReviewDTO struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	CourseID string `json:"course_id"`
	Text     string `json:"text"`
	Rating   int64  `json:"rating"`
}

func PrintReviewDTO(d ReviewDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("User ID: %s\n", d.UserID)
	fmt.Printf("Course ID: %s\n", d.CourseID)
	fmt.Printf("Text: %s\n", d.Text)
	fmt.Printf("Rating: %d\n", d.Rating)
}

func NewReviewDTO(review domain.Review) ReviewDTO {
	return ReviewDTO{
		ID:       review.ID.String(),
		UserID:   review.UserID.String(),
		CourseID: review.CourseID.String(),
		Text:     review.Text,
		Rating:   review.Rating,
	}
}
