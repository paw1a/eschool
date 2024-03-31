package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

type ReviewService struct {
	repo port.IReviewRepository
}

func NewReviewService(repo port.IReviewRepository) *ReviewService {
	return &ReviewService{
		repo: repo,
	}
}

func (r *ReviewService) FindAll(ctx context.Context) ([]domain.Review, error) {
	return r.repo.FindAll(ctx)
}

func (r *ReviewService) FindByID(ctx context.Context, reviewID int64) (domain.Review, error) {
	return r.repo.FindByID(ctx, reviewID)
}

func (r *ReviewService) FindUserReviews(ctx context.Context, userID int64) ([]domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewService) FindCourseReviews(ctx context.Context, courseID int64) ([]domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewService) CreateCourseReview(ctx context.Context, courseID, userID int64,
	reviewDTO dto.CreateReviewDTO) (domain.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReviewService) Delete(ctx context.Context, reviewID int64) error {
	//TODO implement me
	panic("implement me")
}
