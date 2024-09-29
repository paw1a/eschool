package test

import (
	"github.com/paw1a/eschool/internal/core/domain"
)

type ReviewBuilder struct {
	review domain.Review
}

func NewReviewBuilder() *ReviewBuilder {
	return &ReviewBuilder{
		review: domain.Review{
			ID:       domain.NewID(),
			UserID:   domain.NewID(),
			CourseID: domain.NewID(),
			Text:     "text",
		},
	}
}

func (b *ReviewBuilder) WithID(id domain.ID) *ReviewBuilder {
	b.review.ID = id
	return b
}

func (b *ReviewBuilder) WithUserID(userID domain.ID) *ReviewBuilder {
	b.review.UserID = userID
	return b
}

func (b *ReviewBuilder) WithCourseID(courseID domain.ID) *ReviewBuilder {
	b.review.CourseID = courseID
	return b
}

func (b *ReviewBuilder) WithText(text string) *ReviewBuilder {
	b.review.Text = text
	return b
}

func (b *ReviewBuilder) Build() domain.Review {
	return b.review
}
