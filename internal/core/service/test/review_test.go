package test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"go.uber.org/zap"
)

type ReviewSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *ReviewSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// FindAll Suite
type ReviewFindAllSuite struct {
	ReviewSuite
}

func ReviewFindAllSuccessRepositoryMock(repository *mocks.ReviewRepository) {
	repository.
		On("FindAll", context.Background()).
		Return([]domain.Review{NewReviewBuilder().Build()}, nil)
}

func (s *ReviewFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Find all reviews success")
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindAllSuccessRepositoryMock(reviewRepository)
	_, err := reviewService.FindAll(context.Background())
	t.Assert().Nil(err)
}

func ReviewFindAllFailureRepositoryMock(repository *mocks.ReviewRepository) {
	repository.
		On("FindAll", context.Background()).
		Return(nil, errs.ErrNotExist)
}

func (s *ReviewFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find all reviews failure")
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindAllFailureRepositoryMock(reviewRepository)
	_, err := reviewService.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindAll", new(ReviewFindAllSuite))
}

// FindByID Suite
type ReviewFindByIDSuite struct {
	ReviewSuite
}

func ReviewFindByIDSuccessRepositoryMock(repository *mocks.ReviewRepository, reviewID domain.ID) {
	repository.
		On("FindByID", context.Background(), reviewID).
		Return(NewReviewBuilder().WithID(reviewID).Build(), nil)
}

func (s *ReviewFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Title("Find review by id success")
	reviewID := domain.NewID()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindByIDSuccessRepositoryMock(reviewRepository, reviewID)
	review, err := reviewService.FindByID(context.Background(), reviewID)
	t.Assert().Nil(err)
	t.Assert().Equal(reviewID, review.ID)
}

func ReviewFindByIDFailureRepositoryMock(repository *mocks.ReviewRepository, reviewID domain.ID) {
	repository.
		On("FindByID", context.Background(), reviewID).
		Return(domain.Review{}, errs.ErrNotExist)
}

func (s *ReviewFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Title("Find review by id failure")
	reviewID := domain.NewID()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindByIDFailureRepositoryMock(reviewRepository, reviewID)
	_, err := reviewService.FindByID(context.Background(), reviewID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindByID", new(ReviewFindByIDSuite))
}

// Create Suite
type ReviewCreateSuite struct {
	ReviewSuite
}

func ReviewCreateSuccessRepositoryMock(repository *mocks.ReviewRepository) {
	repository.
		On("Create", context.Background(), mock.Anything, mock.Anything, mock.Anything).
		Return(NewReviewBuilder().Build(), nil)
}

func (s *ReviewCreateSuite) TestCreate_Success(t provider.T) {
	t.Title("Create review success")
	param := NewCreateReviewParamBuilder().Build()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewCreateSuccessRepositoryMock(reviewRepository)
	review, err := reviewService.CreateCourseReview(context.Background(),
		domain.NewID(), domain.NewID(), param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Text, review.Text)
}

func ReviewCreateFailureRepositoryMock(repository *mocks.ReviewRepository) {
	repository.
		On("Create", context.Background(), mock.Anything, mock.Anything, mock.Anything).
		Return(domain.Review{}, errors.New("error"))
}

func (s *ReviewCreateSuite) TestCreate_Failure(t provider.T) {
	t.Title("Create review failure")
	param := NewCreateReviewParamBuilder().Build()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewCreateFailureRepositoryMock(reviewRepository)
	_, err := reviewService.CreateCourseReview(context.Background(), domain.NewID(), domain.NewID(), param)
	t.Assert().NotNil(err)
}

func TestReviewCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Create", new(ReviewCreateSuite))
}

// Delete Suite
type ReviewDeleteSuite struct {
	ReviewSuite
}

func ReviewDeleteSuccessRepositoryMock(repository *mocks.ReviewRepository, reviewID domain.ID) {
	repository.
		On("Delete", context.Background(), reviewID).
		Return(nil)
}

func (s *ReviewDeleteSuite) TestDelete_Success(t provider.T) {
	t.Title("Delete review success")
	reviewID := domain.NewID()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewDeleteSuccessRepositoryMock(reviewRepository, reviewID)
	err := reviewService.Delete(context.Background(), reviewID)
	t.Assert().Nil(err)
}

func ReviewDeleteFailureRepositoryMock(repository *mocks.ReviewRepository, reviewID domain.ID) {
	repository.
		On("Delete", context.Background(), reviewID).
		Return(errs.ErrDeleteFailed)
}

func (s *ReviewDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Title("Delete review failure")
	reviewID := domain.NewID()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewDeleteFailureRepositoryMock(reviewRepository, reviewID)
	err := reviewService.Delete(context.Background(), reviewID)
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestReviewDeleteSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "Delete", new(ReviewDeleteSuite))
}

// FindUserReviews Suite
type ReviewFindUserReviewsSuite struct {
	ReviewSuite
}

func ReviewFindUserReviewsSuccessRepositoryMock(repository *mocks.ReviewRepository, userID domain.ID) {
	repository.
		On("FindUserReviews", context.Background(), userID).
		Return([]domain.Review{NewReviewBuilder().Build()}, nil)
}

func (s *ReviewFindUserReviewsSuite) TestFindUserReviews_Success(t provider.T) {
	t.Title("Find user reviews success")
	userID := domain.NewID()
	review := NewReviewBuilder().Build()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindUserReviewsSuccessRepositoryMock(reviewRepository, userID)
	reviews, err := reviewService.FindUserReviews(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(reviews[0].Text, review.Text)
}

func ReviewFindUserReviewsFailureRepositoryMock(repository *mocks.ReviewRepository, userID domain.ID) {
	repository.
		On("FindUserReviews", context.Background(), userID).
		Return([]domain.Review{{}}, errs.ErrNotExist)
}

func (s *ReviewFindUserReviewsSuite) TestFindUserReviews_Failure(t provider.T) {
	t.Title("Find user reviews failure")
	userID := domain.NewID()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindUserReviewsFailureRepositoryMock(reviewRepository, userID)
	_, err := reviewService.FindUserReviews(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindUserReviewsIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindUserReviews", new(ReviewFindUserReviewsSuite))
}

// FindCourseReviews Suite
type ReviewFindCourseReviewsSuite struct {
	ReviewSuite
}

func ReviewFindCourseReviewsSuccessRepositoryMock(repository *mocks.ReviewRepository, courseID domain.ID) {
	repository.
		On("FindCourseReviews", context.Background(), courseID).
		Return([]domain.Review{NewReviewBuilder().Build()}, nil)
}

func (s *ReviewFindCourseReviewsSuite) TestFindCourseReviews_Success(t provider.T) {
	t.Title("Find course reviews success")
	courseID := domain.NewID()
	review := NewReviewBuilder().Build()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindCourseReviewsSuccessRepositoryMock(reviewRepository, courseID)
	reviews, err := reviewService.FindCourseReviews(context.Background(), courseID)
	t.Assert().Nil(err)
	t.Assert().Equal(reviews[0].Text, review.Text)
}

func ReviewFindCourseReviewsFailureRepositoryMock(repository *mocks.ReviewRepository, courseID domain.ID) {
	repository.
		On("FindCourseReviews", context.Background(), courseID).
		Return([]domain.Review{{}}, errs.ErrNotExist)
}

func (s *ReviewFindCourseReviewsSuite) TestFindCourseReviews_Failure(t provider.T) {
	t.Title("Find course reviews failure")
	courseID := domain.NewID()
	reviewRepository := mocks.NewReviewRepository(t)
	reviewService := service.NewReviewService(reviewRepository, s.logger)
	ReviewFindCourseReviewsFailureRepositoryMock(reviewRepository, courseID)
	_, err := reviewService.FindCourseReviews(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindCourseReviewsIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindCourseReviews", new(ReviewFindCourseReviewsSuite))
}
