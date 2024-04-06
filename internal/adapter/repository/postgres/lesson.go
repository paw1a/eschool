package repository

import (
	"context"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/domain"
)

type LessonRepository struct {
}

func (l *LessonRepository) FindAll(ctx context.Context) ([]domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) FindByID(ctx context.Context, lessonID int64) (domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) FindCourseLessons(ctx context.Context, courseID int64) ([]domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) Create(ctx context.Context, lessonDTO dto.CreateLessonDTO) (domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) AddLessonTests(ctx context.Context, lessonID int64, tests []dto.CreateTestDTO) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) DeleteLessonTest(ctx context.Context, lessonID, testID int64) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) UpdateLessonTest(ctx context.Context, lessonID, testID int64, testDTO dto.UpdateTestDTO) (domain.Test, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) UpdateLessonTheory(ctx context.Context, lessonID int64, theoryDTO dto.UpdateTheoryDTO) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) UpdateLessonVideo(ctx context.Context, lessonID int64, videoDTO dto.UpdateVideoDTO) error {
	//TODO implement me
	panic("implement me")
}

func (l *LessonRepository) Delete(ctx context.Context, lessonID int64) error {
	//TODO implement me
	panic("implement me")
}
