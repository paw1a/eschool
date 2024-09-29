package test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"strings"
	"testing"
)

type MediaSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *MediaSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// SaveMediaFile Suite
type MediaSaveMediaFileSuite struct {
	MediaSuite
}

func MediaSaveMediaFileSuccessRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, nil)
}

func (s *MediaSaveMediaFileSuite) TestSaveMediaFile_Success(t provider.T) {
	t.Parallel()
	t.Title("Save media file success")
	url := domain.Url("url")
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveMediaFileSuccessRepositoryMock(objectStorage, url)
	actual, err := mediaService.SaveMediaFile(context.Background(), NewFileBuilder().Build())
	t.Assert().Nil(err)
	t.Assert().Equal(url, actual)
}

func MediaSaveMediaFileFailureRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, errs.ErrSaveFileError)
}

func (s *MediaSaveMediaFileSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Save media file failure")
	url := domain.Url("url")
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveMediaFileFailureRepositoryMock(objectStorage, url)
	_, err := mediaService.SaveMediaFile(context.Background(), NewFileBuilder().Build())
	t.Assert().ErrorIs(err, errs.ErrSaveFileError)
}

func TestMediaSaveMediaFileSuite(t *testing.T) {
	suite.RunNamedSuite(t, "SaveMediaFile", new(MediaSaveMediaFileSuite))
}

// SaveUserAvatar Suite
type MediaSaveUserAvatarSuite struct {
	MediaSuite
}

func MediaSaveUserAvatarSuccessRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, nil)
}

func (s *MediaSaveUserAvatarSuite) TestSaveUserAvatar_Success(t provider.T) {
	t.Parallel()
	t.Title("Save user avatar success")
	url := domain.Url("url")
	userID := domain.NewID()
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveUserAvatarSuccessRepositoryMock(objectStorage, url)
	actual, err := mediaService.SaveUserAvatar(context.Background(), userID, NewFileBuilder().Build())
	t.Assert().Nil(err)
	t.Assert().Equal(url, actual)
}

func MediaSaveUserAvatarFailureRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, errs.ErrSaveFileError)
}

func (s *MediaSaveUserAvatarSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Save user avatar failure")
	url := domain.Url("url")
	userID := domain.NewID()
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveUserAvatarFailureRepositoryMock(objectStorage, url)
	_, err := mediaService.SaveUserAvatar(context.Background(), userID, NewFileBuilder().Build())
	t.Assert().ErrorIs(err, errs.ErrSaveFileError)
}

func TestMediaSaveUserAvatarSuite(t *testing.T) {
	suite.RunNamedSuite(t, "SaveUserAvatar", new(MediaSaveUserAvatarSuite))
}

// SaveLessonTheory Suite
type MediaSaveLessonTheorySuite struct {
	MediaSuite
}

func MediaSaveLessonTheorySuccessRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, nil)
}

func (s *MediaSaveLessonTheorySuite) TestSaveLessonTheory_Success(t provider.T) {
	t.Parallel()
	t.Title("Save lesson theory success")
	url := domain.Url("url")
	lessonID := domain.NewID()
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveLessonTheorySuccessRepositoryMock(objectStorage, url)
	actual, err := mediaService.SaveLessonTheory(context.Background(),
		lessonID, strings.NewReader("text"))
	t.Assert().Nil(err)
	t.Assert().Equal(url, actual)
}

func MediaSaveLessonTheoryFailureRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, errs.ErrSaveFileError)
}

func (s *MediaSaveLessonTheorySuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Save lesson theory failure")
	url := domain.Url("url")
	lessonID := domain.NewID()
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveLessonTheoryFailureRepositoryMock(objectStorage, url)
	_, err := mediaService.SaveLessonTheory(context.Background(),
		lessonID, strings.NewReader("text"))
	t.Assert().ErrorIs(err, errs.ErrSaveFileError)
}

func TestMediaSaveLessonTheorySuite(t *testing.T) {
	suite.RunNamedSuite(t, "SaveLessonTheory", new(MediaSaveLessonTheorySuite))
}

// SaveTestQuestion Suite
type MediaSaveTestQuestionSuite struct {
	MediaSuite
}

func MediaSaveTestQuestionSuccessRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, nil)
}

func (s *MediaSaveTestQuestionSuite) TestSaveTestQuestion_Success(t provider.T) {
	t.Parallel()
	t.Title("Save test question success")
	url := domain.Url("url")
	lessonID := domain.NewID()
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveTestQuestionSuccessRepositoryMock(objectStorage, url)
	actual, err := mediaService.SaveTestQuestion(context.Background(),
		lessonID, strings.NewReader("text"))
	t.Assert().Nil(err)
	t.Assert().Equal(url, actual)
}

func MediaSaveTestQuestionFailureRepositoryMock(storage *mocks.ObjectStorage, url domain.Url) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(url, errs.ErrSaveFileError)
}

func (s *MediaSaveTestQuestionSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Save test question failure")
	url := domain.Url("url")
	lessonID := domain.NewID()
	objectStorage := mocks.NewObjectStorage(t)
	mediaService := service.NewMediaService(objectStorage, s.logger)
	MediaSaveTestQuestionFailureRepositoryMock(objectStorage, url)
	_, err := mediaService.SaveTestQuestion(context.Background(),
		lessonID, strings.NewReader("text"))
	t.Assert().ErrorIs(err, errs.ErrSaveFileError)
}

func TestMediaSaveTestQuestionSuite(t *testing.T) {
	suite.RunNamedSuite(t, "SaveTestQuestion", new(MediaSaveTestQuestionSuite))
}
