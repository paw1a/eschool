package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/zap"
)

type ReviewService struct {
	repo   port.IReviewRepository
	logger *zap.Logger
}

func NewReviewService(repo port.IReviewRepository, logger *zap.Logger) *ReviewService {
	return &ReviewService{
		repo:   repo,
		logger: logger,
	}
}

func (r *ReviewService) FindAll(ctx context.Context) ([]domain.Review, error) {
	return r.repo.FindAll(ctx)
}

func (r *ReviewService) FindByID(ctx context.Context, reviewID domain.ID) (domain.Review, error) {
	return r.repo.FindByID(ctx, reviewID)
}

func (r *ReviewService) FindUserReviews(ctx context.Context, userID domain.ID) ([]domain.Review, error) {
	return r.repo.FindUserReviews(ctx, userID)
}

func (r *ReviewService) FindCourseReviews(ctx context.Context, courseID domain.ID) ([]domain.Review, error) {
	return r.repo.FindCourseReviews(ctx, courseID)
}

func (r *ReviewService) CreateCourseReview(ctx context.Context, courseID, userID domain.ID,
	param port.CreateReviewParam) (domain.Review, error) {
	return r.repo.Create(ctx, domain.Review{
		ID:       domain.NewID(),
		UserID:   userID,
		CourseID: courseID,
		Text:     param.Text,
		Rating:   param.Rating,
	})
}

func (r *ReviewService) Delete(ctx context.Context, reviewID domain.ID) error {
	return r.repo.Delete(ctx, reviewID)
}
