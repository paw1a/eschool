package test

import (
	"context"
	"errors"
	"github.com/guregu/null"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

type LessonSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *LessonSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// FindAll Suite
type LessonFindAllSuite struct {
	LessonSuite
}

func LessonFindAllSuccessRepositoryMock(repository *mocks.LessonRepository) {
	repository.
		On("FindAll", context.Background()).
		Return([]domain.Lesson{NewLessonBuilder().Build()}, nil)
}

func (s *LessonFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Find all lessons success")
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindAllSuccessRepositoryMock(lessonRepository)
	_, err := lessonService.FindAll(context.Background())
	t.Assert().Nil(err)
}

func LessonFindAllFailureRepositoryMock(repository *mocks.LessonRepository) {
	repository.
		On("FindAll", context.Background()).
		Return(nil, errs.ErrNotExist)
}

func (s *LessonFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find all lessons failure")
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindAllFailureRepositoryMock(lessonRepository)
	_, err := lessonService.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindAll", new(LessonFindAllSuite))
}

// FindByID Suite
type LessonFindByIDSuite struct {
	LessonSuite
}

func LessonFindByIDSuccessRepositoryMock(repository *mocks.LessonRepository, lessonID domain.ID) {
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().WithID(lessonID).Build(), nil)
}

func (s *LessonFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Title("Find lesson by id success")
	lessonID := domain.NewID()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindByIDSuccessRepositoryMock(lessonRepository, lessonID)
	lesson, err := lessonService.FindByID(context.Background(), lessonID)
	t.Assert().Nil(err)
	t.Assert().Equal(lessonID, lesson.ID)
}

func LessonFindByIDFailureRepositoryMock(repository *mocks.LessonRepository, lessonID domain.ID) {
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(domain.Lesson{}, errs.ErrNotExist)
}

func (s *LessonFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Title("Find lesson by id failure")
	lessonID := domain.NewID()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindByIDFailureRepositoryMock(lessonRepository, lessonID)
	_, err := lessonService.FindByID(context.Background(), lessonID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindByID", new(LessonFindByIDSuite))
}

// FindUserLessons Suite
type LessonFindCourseLessonsSuite struct {
	LessonSuite
}

func LessonFindCourseLessonsSuccessRepositoryMock(repository *mocks.LessonRepository, userID domain.ID) {
	repository.
		On("FindCourseLessons", context.Background(), userID).
		Return([]domain.Lesson{NewLessonBuilder().Build()}, nil)
}

func (s *LessonFindCourseLessonsSuite) TestFindUserLessons_Success(t provider.T) {
	t.Title("Find course lessons success")
	courseID := domain.NewID()
	lesson := NewLessonBuilder().Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindCourseLessonsSuccessRepositoryMock(lessonRepository, courseID)
	lessons, err := lessonService.FindCourseLessons(context.Background(), courseID)
	t.Assert().Nil(err)
	t.Assert().Equal(lessons[0].Title, lesson.Title)
}

func LessonFindCourseLessonsFailureRepositoryMock(repository *mocks.LessonRepository, courseID domain.ID) {
	repository.
		On("FindCourseLessons", context.Background(), courseID).
		Return([]domain.Lesson{{}}, errs.ErrNotExist)
}

func (s *LessonFindCourseLessonsSuite) TestFindUserLessons_Failure(t provider.T) {
	t.Title("Find course lessons failure")
	courseID := domain.NewID()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindCourseLessonsFailureRepositoryMock(lessonRepository, courseID)
	_, err := lessonService.FindCourseLessons(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindCourseLessonsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindCourseLessons", new(LessonFindCourseLessonsSuite))
}

// CreateTheoryLesson Suite
type LessonCreateTheoryLessonSuite struct {
	LessonSuite
}

func LessonCreateTheoryLessonSuccessRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage, title string) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewLessonBuilder().WithTitle(title).Build(), nil)
	objectStorage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url("url"), nil)
}

func (s *LessonCreateTheoryLessonSuite) TestCreateTheoryLesson_Success(t provider.T) {
	t.Title("Create theory lesson success")
	courseID := domain.NewID()
	title := "lesson name"
	param := NewCreateTheoryParamBuilder().WithTitle(title).Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonCreateTheoryLessonSuccessRepositoryMock(lessonRepository, objectStorage, title)
	lesson, err := lessonService.CreateTheoryLesson(context.Background(), courseID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Title, lesson.Title)
}

func LessonCreateTheoryLessonFailureRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.Lesson{}, errors.New("error"))
	objectStorage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url("url"), nil)
}

func (s *LessonCreateTheoryLessonSuite) TestCreateTheoryLesson_Failure(t provider.T) {
	t.Title("Create theory lesson failure")
	courseID := domain.NewID()
	param := NewCreateTheoryParamBuilder().Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonCreateTheoryLessonFailureRepositoryMock(lessonRepository, objectStorage)
	_, err := lessonService.CreateTheoryLesson(context.Background(), courseID, param)
	t.Assert().NotNil(err)
}

func TestLessonCreateTheoryLessonSuite(t *testing.T) {
	suite.RunNamedSuite(t, "CreateTheoryLesson", new(LessonCreateTheoryLessonSuite))
}

// CreateVideoLesson Suite
type LessonCreateVideoLessonSuite struct {
	LessonSuite
}

func LessonCreateVideoLessonSuccessRepositoryMock(repository *mocks.LessonRepository, title string) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewLessonBuilder().WithTitle(title).Build(), nil)
}

func (s *LessonCreateVideoLessonSuite) TestCreateVideoLesson_Success(t provider.T) {
	t.Title("Create video lesson success")
	courseID := domain.NewID()
	title := "lesson name"
	param := NewCreateVideoParamBuilder().WithTitle(title).Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonCreateVideoLessonSuccessRepositoryMock(lessonRepository, title)
	lesson, err := lessonService.CreateVideoLesson(context.Background(), courseID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Title, lesson.Title)
}

func LessonCreateVideoLessonFailureRepositoryMock(repository *mocks.LessonRepository) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.Lesson{}, errors.New("error"))
}

func (s *LessonCreateVideoLessonSuite) TestCreateVideoLesson_Failure(t provider.T) {
	t.Title("Create video lesson failure")
	courseID := domain.NewID()
	param := NewCreateVideoParamBuilder().Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonCreateVideoLessonFailureRepositoryMock(lessonRepository)
	_, err := lessonService.CreateVideoLesson(context.Background(), courseID, param)
	t.Assert().NotNil(err)
}

func TestLessonCreateVideoLessonSuite(t *testing.T) {
	suite.RunNamedSuite(t, "CreateVideoLesson", new(LessonCreateVideoLessonSuite))
}

// CreatePracticeLesson Suite
type LessonCreatePracticeLessonSuite struct {
	LessonSuite
}

func LessonCreatePracticeLessonSuccessRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage, title string) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewLessonBuilder().WithTitle(title).Build(), nil)
	objectStorage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url("url"), nil)
}

func (s *LessonCreatePracticeLessonSuite) TestCreatePracticeLesson_Success(t provider.T) {
	t.Title("Create practice lesson success")
	courseID := domain.NewID()
	title := "lesson name"
	param := NewCreatePracticeParamBuilder().WithTitle(title).Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonCreatePracticeLessonSuccessRepositoryMock(lessonRepository, objectStorage, title)
	lesson, err := lessonService.CreatePracticeLesson(context.Background(), courseID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Title, lesson.Title)
}

func LessonCreatePracticeLessonFailureRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.Lesson{}, errors.New("error"))
	objectStorage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url("url"), nil)
}

func (s *LessonCreatePracticeLessonSuite) TestCreatePracticeLesson_Failure(t provider.T) {
	t.Title("Create practice lesson failure")
	courseID := domain.NewID()
	param := NewCreatePracticeParamBuilder().Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonCreatePracticeLessonFailureRepositoryMock(lessonRepository, objectStorage)
	_, err := lessonService.CreatePracticeLesson(context.Background(), courseID, param)
	t.Assert().NotNil(err)
}

func TestLessonCreatePracticeLessonSuite(t *testing.T) {
	suite.RunNamedSuite(t, "CreatePracticeLesson", new(LessonCreatePracticeLessonSuite))
}

// UpdateTheoryLesson Suite
type LessonUpdateTheoryLessonSuite struct {
	LessonSuite
}

func LessonUpdateTheoryLessonSuccessRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage, title string, lessonID domain.ID) {
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewLessonBuilder().WithTitle(title).Build(), nil)
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().WithID(lessonID).Build(), nil)
}

func (s *LessonUpdateTheoryLessonSuite) TestUpdateTheoryLesson_Success(t provider.T) {
	t.Title("Update theory lesson success")
	lessonID := domain.NewID()
	title := "lesson name"
	param := NewUpdateTheoryParamBuilder().WithTitle(null.StringFrom(title)).Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonUpdateTheoryLessonSuccessRepositoryMock(lessonRepository, objectStorage, title, lessonID)
	lesson, err := lessonService.UpdateTheoryLesson(context.Background(), lessonID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Title.String, lesson.Title)
}

func LessonUpdateTheoryLessonFailureRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage, lessonID domain.ID) {
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(domain.Lesson{}, errors.New("error"))
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().WithID(lessonID).Build(), nil)
}

func (s *LessonUpdateTheoryLessonSuite) TestUpdateTheoryLesson_Failure(t provider.T) {
	t.Title("Update theory lesson failure")
	lessonID := domain.NewID()
	param := NewUpdateTheoryParamBuilder().Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonUpdateTheoryLessonFailureRepositoryMock(lessonRepository, objectStorage, lessonID)
	_, err := lessonService.UpdateTheoryLesson(context.Background(), lessonID, param)
	t.Assert().NotNil(err)
}

func TestLessonUpdateTheoryLessonSuite(t *testing.T) {
	suite.RunNamedSuite(t, "UpdateTheoryLesson", new(LessonUpdateTheoryLessonSuite))
}

// UpdateVideoLesson Suite
type LessonUpdateVideoLessonSuite struct {
	LessonSuite
}

func LessonUpdateVideoLessonSuccessRepositoryMock(repository *mocks.LessonRepository,
	title string, lessonID domain.ID) {
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewLessonBuilder().WithTitle(title).Build(), nil)
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().WithID(lessonID).Build(), nil)
}

func (s *LessonUpdateVideoLessonSuite) TestUpdateVideoLesson_Success(t provider.T) {
	t.Title("Update video lesson success")
	lessonID := domain.NewID()
	title := "lesson name"
	param := NewUpdateVideoParamBuilder().WithTitle(null.StringFrom(title)).Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonUpdateVideoLessonSuccessRepositoryMock(lessonRepository, title, lessonID)
	lesson, err := lessonService.UpdateVideoLesson(context.Background(), lessonID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Title.String, lesson.Title)
}

func LessonUpdateVideoLessonFailureRepositoryMock(repository *mocks.LessonRepository, lessonID domain.ID) {
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(domain.Lesson{}, errors.New("error"))
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().WithID(lessonID).Build(), nil)
}

func (s *LessonUpdateVideoLessonSuite) TestUpdateVideoLesson_Failure(t provider.T) {
	t.Title("Update video lesson failure")
	lessonID := domain.NewID()
	param := NewUpdateVideoParamBuilder().Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonUpdateVideoLessonFailureRepositoryMock(lessonRepository, lessonID)
	_, err := lessonService.UpdateVideoLesson(context.Background(), lessonID, param)
	t.Assert().NotNil(err)
}

func TestLessonUpdateVideoLessonSuite(t *testing.T) {
	suite.RunNamedSuite(t, "UpdateVideoLesson", new(LessonUpdateVideoLessonSuite))
}

// UpdatePracticeLesson Suite
type LessonUpdatePracticeLessonSuite struct {
	LessonSuite
}

func LessonUpdatePracticeLessonSuccessRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage, title string, lessonID domain.ID) {
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewLessonBuilder().WithTitle(title).Build(), nil)
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().WithID(lessonID).WithType(domain.PracticeLesson).Build(), nil)
	objectStorage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url("url"), nil)
}

func (s *LessonUpdatePracticeLessonSuite) TestUpdatePracticeLesson_Success(t provider.T) {
	t.Title("Update practice lesson success")
	lessonID := domain.NewID()
	title := "lesson name"
	param := NewUpdatePracticeParamBuilder().WithTitle(null.StringFrom(title)).Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonUpdatePracticeLessonSuccessRepositoryMock(lessonRepository, objectStorage, title, lessonID)
	lesson, err := lessonService.UpdatePracticeLesson(context.Background(), lessonID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Title.String, lesson.Title)
}

func LessonUpdatePracticeLessonFailureRepositoryMock(repository *mocks.LessonRepository,
	objectStorage *mocks.ObjectStorage, lessonID domain.ID) {
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(domain.Lesson{}, errors.New("error"))
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(NewLessonBuilder().WithID(lessonID).WithType(domain.PracticeLesson).Build(), nil)
	objectStorage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url("url"), nil)
}

func (s *LessonUpdatePracticeLessonSuite) TestUpdatePracticeLesson_Failure(t provider.T) {
	t.Title("Update practice lesson failure")
	lessonID := domain.NewID()
	param := NewUpdatePracticeParamBuilder().Build()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonUpdatePracticeLessonFailureRepositoryMock(lessonRepository, objectStorage, lessonID)
	_, err := lessonService.UpdatePracticeLesson(context.Background(), lessonID, param)
	t.Assert().NotNil(err)
}

func TestLessonUpdatePracticeLessonSuite(t *testing.T) {
	suite.RunNamedSuite(t, "UpdatePracticeLesson", new(LessonUpdatePracticeLessonSuite))
}