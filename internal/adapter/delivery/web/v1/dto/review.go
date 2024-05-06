package dto

import "github.com/paw1a/eschool/internal/core/domain"

type CreateReviewDTO struct {
	Text string `json:"text" binding:"required"`
}

type ReviewDTO struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	CourseID string `json:"course_id"`
	Text     string `json:"text"`
}

func NewReviewDTO(review domain.Review) ReviewDTO {
	return ReviewDTO{
		ID:       review.ID.String(),
		UserID:   review.UserID.String(),
		CourseID: review.CourseID.String(),
		Text:     review.Text,
	}
}
