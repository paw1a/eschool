package dto

import "github.com/paw1a/eschool/internal/core/domain"

type CreateReviewDTO struct {
	Text string `json:"text" binding:"required"`
}

type ReviewDTO struct {
	ID       string `json:"id" binding:"required"`
	UserID   string `json:"user_id" binding:"required"`
	CourseID string `json:"course_id" binding:"required"`
	Text     string `json:"text" binding:"required"`
}

func NewReviewDTO(review domain.Review) ReviewDTO {
	return ReviewDTO{
		ID:       review.ID.String(),
		UserID:   review.UserID.String(),
		CourseID: review.CourseID.String(),
		Text:     review.Text,
	}
}
