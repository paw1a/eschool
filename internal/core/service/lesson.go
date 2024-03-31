package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

type LessonService struct {
	repo port.ILessonRepository
}

func NewLessonService(repo port.ILessonRepository) *LessonService {
	return &LessonService{
		repo: repo,
	}
}

func (l *LessonService) FindAll(ctx context.Context) ([]domain.Lesson, error) {
	return l.repo.FindAll(ctx)
}

func (l *LessonService) FindByID(ctx context.Context, lessonID int64) (domain.Lesson, error) {
	return l.repo.FindByID(ctx, lessonID)
}

func (l *LessonService) FindCourseLessons(ctx context.Context, courseID int64) ([]domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonService) Create(ctx context.Context, lessonDTO dto.CreateLessonDTO) (domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonService) AddLessonTests(ctx context.Context, lessonID int64, tests []dto.CreateTestDTO) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonService) DeleteLessonTest(ctx context.Context, lessonID, testID int64) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonService) UpdateLessonTest(ctx context.Context, lessonID, testID int64,
	testDTO dto.UpdateTestDTO) (domain.Test, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonService) UpdateLessonTheory(ctx context.Context, lessonID int64,
	theoryDTO dto.UpdateTheoryDTO) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonService) UpdateLessonVideo(ctx context.Context, lessonID int64,
	videoDTO dto.UpdateVideoDTO) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonService) Delete(ctx context.Context, lessonID int64) error {
	//TODO implement me
	panic("implement me")
}
