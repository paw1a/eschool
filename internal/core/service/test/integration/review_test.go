package integration

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/zap"
	"testing"
)

var courseID = domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027ca")
var userID = domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027ca")

var reviews = []domain.Review{
	domain.Review{
		ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7021ca"),
		UserID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027ca"),
		CourseID: domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027ca"),
		Text:     "review1 text",
	},
	domain.Review{
		ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7021cb"),
		UserID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cb"),
		CourseID: domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027ca"),
		Text:     "review2 text",
	},
	domain.Review{
		ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7021cc"),
		UserID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027ca"),
		CourseID: domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027cb"),
		Text:     "review3 text",
	},
}

var createdReview = domain.Review{
	ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7021cd"),
	UserID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc"),
	CourseID: domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027cb"),
	Text:     "review3 text",
}

type ReviewSuite struct {
	suite.Suite
	logger    *zap.Logger
	container *postgres.PostgresContainer
	db        *sqlx.DB
}

func (s *ReviewSuite) BeforeAll(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

func (s *ReviewSuite) BeforeEach(t provider.T) {
	var err error
	s.container, err = newPostgresContainer(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	url, err := s.container.ConnectionString(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	s.db, err = newPostgresDB(url)
	if err != nil {
		t.Fatal(err)
	}
}

func (s *ReviewSuite) AfterAll(t provider.T) {
	if err := s.container.Terminate(context.Background()); err != nil {
		t.Fatalf("failed to terminate container: %s", err)
	}
}

func (s *ReviewSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *ReviewSuite) TestUserService_FindAll(t provider.T) {
	repo := repository.NewReviewRepo(s.db)
	reviewService := service.NewReviewService(repo, s.logger)
	found, err := reviewService.FindAll(context.Background())
	if err != nil {
		t.Errorf("failed to find all reviews: %v", err)
	}
	t.Assert().Equal(len(found), len(reviews))
	for i := range reviews {
		t.Assert().Equal(reviews[i], found[i])
	}
}

func (s *ReviewSuite) TestUserService_FindByID(t provider.T) {
	repo := repository.NewReviewRepo(s.db)
	reviewService := service.NewReviewService(repo, s.logger)
	review, err := reviewService.FindByID(context.Background(), reviews[0].ID)
	if err != nil {
		t.Errorf("failed to find review with id: %v", err)
	}
	t.Assert().Equal(review, reviews[0])
}

func (s *ReviewSuite) TestUserService_FindUserReviews(t provider.T) {
	repo := repository.NewReviewRepo(s.db)
	reviewService := service.NewReviewService(repo, s.logger)
	found, err := reviewService.FindUserReviews(context.Background(), userID)
	if err != nil {
		t.Errorf("failed to find user reviews: %v", err)
	}

	t.Assert().Equal(len(found), 2)
	t.Assert().Equal(reviews[0], found[0])
	t.Assert().Equal(reviews[2], found[1])
}

func (s *ReviewSuite) TestUserService_FindCourseReviews(t provider.T) {
	repo := repository.NewReviewRepo(s.db)
	reviewService := service.NewReviewService(repo, s.logger)
	found, err := reviewService.FindCourseReviews(context.Background(), courseID)
	if err != nil {
		t.Errorf("failed to find course reviews: %v", err)
	}

	t.Assert().Equal(len(found), 2)
	t.Assert().Equal(reviews[0], found[0])
	t.Assert().Equal(reviews[1], found[1])
}

func (s *ReviewSuite) TestUserService_Delete(t provider.T) {
	repo := repository.NewReviewRepo(s.db)
	reviewService := service.NewReviewService(repo, s.logger)
	err := reviewService.Delete(context.Background(), reviews[0].ID)
	if err != nil {
		t.Errorf("failed to delete review: %v", err)
	}
}

func TestReviewSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Review service suite", new(ReviewSuite))
}
