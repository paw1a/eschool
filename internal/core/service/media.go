package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/twinj/uuid"
	"go.uber.org/zap"
	"io"
	"path/filepath"
)

type MediaService struct {
	storage port.IObjectStorage
	logger  *zap.Logger
}

func NewMediaService(storage port.IObjectStorage, logger *zap.Logger) *MediaService {
	return &MediaService{
		storage: storage,
		logger:  logger,
	}
}

func (m *MediaService) SaveMediaFile(ctx context.Context, file domain.File) (domain.Url, error) {
	url, err := m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + "_" + file.Name,
		Path:   "media/",
		Reader: file.Reader,
	})
	if err != nil {
		m.logger.Error("failed to save media file", zap.Error(err),
			zap.String("filename", filepath.Join(file.Path, file.Name)))
		return "", err
	}

	m.logger.Info("media file is successfully saved",
		zap.String("filename", filepath.Join(file.Path, file.Name)))
	return url, nil
}

func (m *MediaService) SaveUserAvatar(ctx context.Context, userID domain.ID, file domain.File) (domain.Url, error) {
	url, err := m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + "_" + file.Name,
		Path:   "avatars/" + userID.String(),
		Reader: file.Reader,
	})
	if err != nil {
		m.logger.Error("failed to save user avatar", zap.Error(err),
			zap.String("filename", filepath.Join(file.Path, file.Name)))
		return "", err
	}

	m.logger.Info("user avatar is successfully saved",
		zap.String("filename", filepath.Join(file.Path, file.Name)))
	return url, nil
}

func (m *MediaService) SaveLessonTheory(ctx context.Context, lessonID domain.ID, reader io.Reader) (domain.Url, error) {
	url, err := m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + ".md",
		Path:   "markdown/lessons/" + lessonID.String(),
		Reader: reader,
	})
	if err != nil {
		m.logger.Error("failed to save lesson theory", zap.Error(err),
			zap.String("lessonID", lessonID.String()))
		return "", err
	}

	m.logger.Info("lesson theory is successfully saved",
		zap.String("lessonID", lessonID.String()))
	return url, nil
}

func (m *MediaService) SaveTestQuestion(ctx context.Context, testID domain.ID, reader io.Reader) (domain.Url, error) {
	url, err := m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + ".md",
		Path:   "markdown/tests/" + testID.String(),
		Reader: reader,
	})
	if err != nil {
		m.logger.Error("failed to save lesson theory", zap.Error(err),
			zap.String("testID", testID.String()))
		return "", err
	}

	m.logger.Info("test question is successfully saved",
		zap.String("testID", testID.String()))
	return url, nil
}
