package repository

import (
	"context"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/domain"
)

type ReviewRepository struct {
}

func (r *ReviewRepository) FindAll(ctx context.Context) ([]domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewRepository) FindByID(ctx context.Context, reviewID int64) (domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewRepository) FindUserReviews(ctx context.Context, userID int64) ([]domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewRepository) FindCourseReviews(ctx context.Context, courseID int64) ([]domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewRepository) CreateCourseReview(ctx context.Context, courseID, userID int64, reviewDTO dto.CreateReviewDTO) (domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewRepository) Delete(ctx context.Context, reviewID int64) error {
	//TODO implement me
	panic("implement me")
}
