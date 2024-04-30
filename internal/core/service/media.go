package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/twinj/uuid"
	"io"
)

type MediaService struct {
	storage port.IObjectStorage
}

func NewMediaService(storage port.IObjectStorage) *MediaService {
	return &MediaService{
		storage: storage,
	}
}

func (m *MediaService) SaveMediaFile(ctx context.Context, file domain.File) (domain.Url, error) {
	return m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + "_" + file.Name,
		Path:   "media/",
		Reader: file.Reader,
	})
}

func (m *MediaService) SaveUserAvatar(ctx context.Context, userID domain.ID, file domain.File) (domain.Url, error) {
	return m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + "_" + file.Name,
		Path:   "avatars/" + userID.String(),
		Reader: file.Reader,
	})
}

func (m *MediaService) SaveLessonTheory(ctx context.Context, lessonID domain.ID, reader io.Reader) (domain.Url, error) {
	return m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + ".md",
		Path:   "markdown/lessons/" + lessonID.String(),
		Reader: reader,
	})
}

func (m *MediaService) SaveTestQuestion(ctx context.Context, testID domain.ID, reader io.Reader) (domain.Url, error) {
	return m.storage.SaveFile(ctx, domain.File{
		Name:   uuid.NewV4().String() + ".md",
		Path:   "markdown/tests/" + testID.String(),
		Reader: reader,
	})
}
