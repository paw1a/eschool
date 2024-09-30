package test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

type StatSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *StatSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// FindLessonStat Suite
type StatFindLessonStatSuite struct {
	StatSuite
}

func StatFindByIDSuccessRepositoryMock(repository *mocks.StatRepository, userID, lessonID domain.ID) {
	repository.
		On("FindLessonStat", context.Background(), userID, lessonID).
		Return(domain.LessonStat{UserID: userID, LessonID: lessonID}, nil)
}

func (s *StatFindLessonStatSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Find lesson stat success")
	userID := domain.NewID()
	lessonID := domain.NewID()
	statRepository := mocks.NewStatRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	statService := service.NewStatService(statRepository, lessonRepository, s.logger)
	StatFindByIDSuccessRepositoryMock(statRepository, userID, lessonID)
	stat, err := statService.FindLessonStat(context.Background(), userID, lessonID)
	t.Assert().Nil(err)
	t.Assert().Equal(stat.UserID, userID)
	t.Assert().Equal(stat.LessonID, lessonID)
}

func StatFindByIDFailureRepositoryMock(repository *mocks.StatRepository, userID, lessonID domain.ID) {
	repository.
		On("FindLessonStat", context.Background(), userID, lessonID).
		Return(domain.LessonStat{}, errs.ErrNotExist)
}

func (s *StatFindLessonStatSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find stat by id failure")
	userID := domain.NewID()
	lessonID := domain.NewID()
	statRepository := mocks.NewStatRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	statService := service.NewStatService(statRepository, lessonRepository, s.logger)
	StatFindByIDFailureRepositoryMock(statRepository, userID, lessonID)
	_, err := statService.FindLessonStat(context.Background(), userID, lessonID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestStatFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Find lesson stat", new(StatFindLessonStatSuite))
}

// Create Suite
type StatCreateSuite struct {
	StatSuite
}

func StatCreateSuccessRepositoryMock(repository *mocks.StatRepository,
	lessonRepository *mocks.LessonRepository, lessonID domain.ID) {
	repository.
		On("CreateLessonStat", context.Background(), mock.Anything).
		Return(nil)
	lessonRepository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().Build(), nil)
}

func (s *StatCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Create test success")
	userID := domain.NewID()
	lessonID := domain.NewID()
	statRepository := mocks.NewStatRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	statService := service.NewStatService(statRepository, lessonRepository, s.logger)
	StatCreateSuccessRepositoryMock(statRepository, lessonRepository, lessonID)
	err := statService.CreateLessonStat(context.Background(), userID, lessonID)
	t.Assert().Nil(err)
}

func StatCreateFailureRepositoryMock(repository *mocks.StatRepository,
	lessonRepository *mocks.LessonRepository, lessonID domain.ID) {
	lessonRepository.
		On("FindByID", context.Background(), lessonID).
		Return(domain.Lesson{}, errs.ErrNotExist)
}

func (s *StatCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Create test failure")
	userID := domain.NewID()
	lessonID := domain.NewID()
	statRepository := mocks.NewStatRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	statService := service.NewStatService(statRepository, lessonRepository, s.logger)
	StatCreateFailureRepositoryMock(statRepository, lessonRepository, lessonID)
	err := statService.CreateLessonStat(context.Background(), userID, lessonID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestStatCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Create test stat", new(StatCreateSuite))
}

// Update Suite
type StatUpdateSuite struct {
	StatSuite
}

func StatUpdateSuccessRepositoryMock(repository *mocks.StatRepository, userID, lessonID domain.ID) {
	repository.
		On("UpdateLessonStat", context.Background(), mock.Anything).
		Return(nil)
	repository.
		On("FindLessonStat", context.Background(), userID, lessonID).
		Return(domain.LessonStat{LessonID: lessonID}, nil)
}

func (s *StatUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Title("Update lesson stat success")
	userID := domain.NewID()
	lessonID := domain.NewID()
	statRepository := mocks.NewStatRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	statService := service.NewStatService(statRepository, lessonRepository, s.logger)
	StatUpdateSuccessRepositoryMock(statRepository, userID, lessonID)
	err := statService.UpdateLessonStat(context.Background(), userID, lessonID,
		port.UpdateLessonStatParam{})
	t.Assert().Nil(err)
}

func StatUpdateFailureRepositoryMock(repository *mocks.StatRepository, userID, lessonID domain.ID) {
	repository.
		On("FindLessonStat", context.Background(), userID, lessonID).
		Return(domain.LessonStat{}, errs.ErrNotExist)
}

func (s *StatUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Title("Update lesson stat failure")
	userID := domain.NewID()
	lessonID := domain.NewID()
	statRepository := mocks.NewStatRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	statService := service.NewStatService(statRepository, lessonRepository, s.logger)
	StatUpdateFailureRepositoryMock(statRepository, userID, lessonID)
	err := statService.UpdateLessonStat(context.Background(), userID, lessonID,
		port.UpdateLessonStatParam{})
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestStatUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Update lesson stat", new(StatUpdateSuite))
}
