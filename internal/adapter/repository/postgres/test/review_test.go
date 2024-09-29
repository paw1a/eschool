package test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/adapter/repository/postgres/entity"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"testing"
)

type ReviewSuite struct {
	suite.Suite
}

func NewReviewRepository() (port.IReviewRepository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	conn := sqlx.NewDb(db, "pgx")
	repo := repository.NewReviewRepo(conn)
	return repo, mock
}

// FindAll Suite
type ReviewFindAllSuite struct {
	ReviewSuite
}

func (s *ReviewFindAllSuite) ReviewFindAllSuccessRepositoryMock(mock sqlmock.Sqlmock, review domain.Review) {
	pgReview := entity.NewPgReview(review)
	expectedRows := sqlmock.NewRows(EntityColumns(pgReview))
	expectedRows.AddRow(EntityValues(pgReview)...)
	mock.ExpectQuery(repository.ReviewFindAllQuery).WillReturnRows(expectedRows)
}

func (s *ReviewFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	s.ReviewFindAllSuccessRepositoryMock(mock, review)
	reviews, err := repo.FindAll(context.Background())
	t.Assert().Nil(err)
	t.Assert().Equal(reviews[0].Text, review.Text)
}

func (s *ReviewFindAllSuite) ReviewFindAllFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.ReviewFindAllQuery).WillReturnError(sql.ErrNoRows)
}

func (s *ReviewFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewReviewRepository()
	s.ReviewFindAllFailureRepositoryMock(mock)
	_, err := repo.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindAll", new(ReviewFindAllSuite))
}

type ReviewFindByIDSuite struct {
	ReviewSuite
}

func (s *ReviewFindByIDSuite) ReviewFindByIDSuccessRepositoryMock(mock sqlmock.Sqlmock, review domain.Review) {
	pgReview := entity.NewPgReview(review)
	expectedRows := sqlmock.NewRows(EntityColumns(pgReview)).
		AddRow(EntityValues(pgReview)...)
	mock.ExpectQuery(repository.ReviewFindByIDQuery).WithArgs(review.ID).WillReturnRows(expectedRows)
}

func (s *ReviewFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find by ID success")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	s.ReviewFindByIDSuccessRepositoryMock(mock, review)
	foundReview, err := repo.FindByID(context.Background(), review.ID)
	t.Assert().Nil(err)
	t.Assert().Equal(foundReview.ID, review.ID)
}

func (s *ReviewFindByIDSuite) ReviewFindByIDFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.ReviewFindByIDQuery).WillReturnError(sql.ErrNoRows)
}

func (s *ReviewFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find by ID failure")
	repo, mock := NewReviewRepository()
	s.ReviewFindByIDFailureRepositoryMock(mock)
	_, err := repo.FindByID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindByID", new(ReviewFindByIDSuite))
}

type ReviewFindUserReviewsSuite struct {
	ReviewSuite
}

func (s *ReviewFindUserReviewsSuite) ReviewFindUserReviewsSuccessRepositoryMock(mock sqlmock.Sqlmock,
	review domain.Review, userID domain.ID) {
	pgReview := entity.NewPgReview(review)
	expectedRows := sqlmock.NewRows(EntityColumns(pgReview))
	expectedRows.AddRow(EntityValues(pgReview)...)
	mock.ExpectQuery(repository.ReviewFindUserReviewsQuery).WithArgs(userID).WillReturnRows(expectedRows)
}

func (s *ReviewFindUserReviewsSuite) TestFindUserReviews_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	userID := domain.NewID()
	s.ReviewFindUserReviewsSuccessRepositoryMock(mock, review, userID)
	reviews, err := repo.FindUserReviews(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(reviews[0].Text, review.Text)
}

func (s *ReviewFindUserReviewsSuite) ReviewFindUserReviewsFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.ReviewFindUserReviewsQuery).WillReturnError(sql.ErrNoRows)
}

func (s *ReviewFindUserReviewsSuite) TestFindUserReviews_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewReviewRepository()
	userID := domain.NewID()
	s.ReviewFindUserReviewsFailureRepositoryMock(mock)
	_, err := repo.FindUserReviews(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindUserReviewsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindUserReviews", new(ReviewFindUserReviewsSuite))
}

type ReviewFindCourseReviewsSuite struct {
	ReviewSuite
}

func (s *ReviewFindCourseReviewsSuite) ReviewFindCourseReviewsSuccessRepositoryMock(mock sqlmock.Sqlmock,
	review domain.Review, courseID domain.ID) {
	pgReview := entity.NewPgReview(review)
	expectedRows := sqlmock.NewRows(EntityColumns(pgReview))
	expectedRows.AddRow(EntityValues(pgReview)...)
	mock.ExpectQuery(repository.ReviewFindCourseReviewsQuery).WithArgs(courseID).WillReturnRows(expectedRows)
}

func (s *ReviewFindCourseReviewsSuite) TestFindCourseReviews_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	courseID := domain.NewID()
	s.ReviewFindCourseReviewsSuccessRepositoryMock(mock, review, courseID)
	reviews, err := repo.FindCourseReviews(context.Background(), courseID)
	t.Assert().Nil(err)
	t.Assert().Equal(reviews[0].Text, review.Text)
}

func (s *ReviewFindCourseReviewsSuite) ReviewFindCourseReviewsFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.ReviewFindCourseReviewsQuery).WillReturnError(sql.ErrNoRows)
}

func (s *ReviewFindCourseReviewsSuite) TestFindCourseReviews_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewReviewRepository()
	courseID := domain.NewID()
	s.ReviewFindCourseReviewsFailureRepositoryMock(mock)
	_, err := repo.FindCourseReviews(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestReviewFindCourseReviewsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindCourseReviews", new(ReviewFindCourseReviewsSuite))
}

type ReviewCreateSuite struct {
	ReviewSuite
}

func (s *ReviewCreateSuite) ReviewCreateSuccessRepositoryMock(mock sqlmock.Sqlmock, review domain.Review) {
	pgReview := entity.NewPgReview(review)
	queryString := InsertQueryString(pgReview, "review")
	mock.ExpectExec(queryString).
		WithArgs(EntityValues(pgReview)...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedRows := sqlmock.NewRows(EntityColumns(pgReview)).
		AddRow(EntityValues(pgReview)...)
	mock.ExpectQuery(repository.ReviewFindByIDQuery).WithArgs(pgReview.ID).WillReturnRows(expectedRows)
}

func (s *ReviewCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository create review success")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	s.ReviewCreateSuccessRepositoryMock(mock, review)
	createdReview, err := repo.Create(context.Background(), review)
	t.Assert().Nil(err)
	t.Assert().Equal(createdReview.Text, review.Text)
}

func (s *ReviewCreateSuite) ReviewCreateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	queryString := InsertQueryString(entity.PgReview{}, "review")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *ReviewCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository create review failure")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	s.ReviewCreateFailureRepositoryMock(mock)
	_, err := repo.Create(context.Background(), review)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestReviewCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Create", new(ReviewCreateSuite))
}

type ReviewDeleteSuite struct {
	ReviewSuite
}

func (s *ReviewDeleteSuite) ReviewDeleteSuccessRepositoryMock(mock sqlmock.Sqlmock, reviewID domain.ID) {
	mock.ExpectExec(repository.ReviewDeleteQuery).WithArgs(reviewID).WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *ReviewDeleteSuite) TestDelete_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository delete review success")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	s.ReviewDeleteSuccessRepositoryMock(mock, review.ID)
	err := repo.Delete(context.Background(), review.ID)
	t.Assert().Nil(err)
}

func (s *ReviewDeleteSuite) ReviewDeleteFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.ReviewDeleteQuery).WillReturnError(sql.ErrConnDone)
}

func (s *ReviewDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository delete review failure")
	repo, mock := NewReviewRepository()
	review := NewReviewBuilder().Build()
	s.ReviewDeleteFailureRepositoryMock(mock)
	err := repo.Delete(context.Background(), review.ID)
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestReviewDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Delete", new(ReviewDeleteSuite))
}
