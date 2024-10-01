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
		Return([]domain.Lesson{NewLessonMother("title", 10).Create()}, nil)
}

func (s *LessonFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson service find all success")
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
	t.Title("Lesson service find all failure")
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindAllFailureRepositoryMock(lessonRepository)
	_, err := lessonService.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson service find all", new(LessonFindAllSuite))
}

// FindByID Suite
type LessonFindByIDSuite struct {
	LessonSuite
}

func LessonFindByIDSuccessRepositoryMock(repository *mocks.LessonRepository,
	lesson domain.Lesson) {
	repository.
		On("FindByID", context.Background(), lesson.ID).
		Return(lesson, nil)
}

func (s *LessonFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson service find by id success")
	lesson := NewLessonMother("title", 10).Create()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindByIDSuccessRepositoryMock(lessonRepository, lesson)
	actual, err := lessonService.FindByID(context.Background(), lesson.ID)
	t.Assert().Nil(err)
	t.Assert().Equal(actual.ID, lesson.ID)
}

func LessonFindByIDFailureRepositoryMock(repository *mocks.LessonRepository, lessonID domain.ID) {
	repository.
		On("FindByID", context.Background(), lessonID).
		Return(domain.Lesson{}, errs.ErrNotExist)
}

func (s *LessonFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson service find by id failure")
	lessonID := domain.NewID()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindByIDFailureRepositoryMock(lessonRepository, lessonID)
	_, err := lessonService.FindByID(context.Background(), lessonID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson service find by id", new(LessonFindByIDSuite))
}

// FindUserLessons Suite
type LessonFindCourseLessonsSuite struct {
	LessonSuite
}

func LessonFindCourseLessonsSuccessRepositoryMock(repository *mocks.LessonRepository, userID domain.ID,
	lesson domain.Lesson) {
	repository.
		On("FindCourseLessons", context.Background(), userID).
		Return([]domain.Lesson{lesson}, nil)
}

func (s *LessonFindCourseLessonsSuite) TestFindUserLessons_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson service find course lessons success")
	courseID := domain.NewID()
	lesson := NewLessonMother("title", 10).Create()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindCourseLessonsSuccessRepositoryMock(lessonRepository, courseID, lesson)
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
	t.Parallel()
	t.Title("Lesson service find course lessons failure")
	courseID := domain.NewID()
	lessonRepository := mocks.NewLessonRepository(t)
	objectStorage := mocks.NewObjectStorage(t)
	lessonService := service.NewLessonService(lessonRepository, objectStorage, s.logger)
	LessonFindCourseLessonsFailureRepositoryMock(lessonRepository, courseID)
	_, err := lessonService.FindCourseLessons(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindCourseLessonsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson service find course lessons", new(LessonFindCourseLessonsSuite))
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
	t.Parallel()
	t.Title("Lesson service create theory lesson success")
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
	t.Parallel()
	t.Title("Lesson service create theory lesson failure")
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
	suite.RunNamedSuite(t, "Lesson service create theory lesson", new(LessonCreateTheoryLessonSuite))
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
	t.Parallel()
	t.Title("Lesson service create video lesson success")
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
	t.Parallel()
	t.Title("Lesson service create video lesson failure")
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
	suite.RunNamedSuite(t, "Lesson service create video lesson", new(LessonCreateVideoLessonSuite))
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
	t.Parallel()
	t.Title("Lesson service create practice lesson success")
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
	t.Parallel()
	t.Title("Lesson service create practice lesson failure")
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
	suite.RunNamedSuite(t, "Lesson service create practice lesson", new(LessonCreatePracticeLessonSuite))
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
	t.Parallel()
	t.Title("Lesson service update theory lesson success")
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
	t.Parallel()
	t.Title("Lesson service update theory lesson failure")
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
	suite.RunNamedSuite(t, "Lesson service update theory lesson", new(LessonUpdateTheoryLessonSuite))
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
	t.Parallel()
	t.Title("Lesson service update video lesson success")
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
	t.Parallel()
	t.Title("Lesson service update video lesson failure")
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
	suite.RunNamedSuite(t, "Lesson service update video lesson", new(LessonUpdateVideoLessonSuite))
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
	t.Parallel()
	t.Title("Lesson service update practice lesson success")
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
	t.Parallel()
	t.Title("Lesson service update practice lesson failure")
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
	suite.RunNamedSuite(t, "Lesson service update practice lesson", new(LessonUpdatePracticeLessonSuite))
}
