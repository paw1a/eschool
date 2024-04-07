package service

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
	domainErr "github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/pkg/errors"
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

func (l *LessonService) FindByID(ctx context.Context, lessonID domain.ID) (domain.Lesson, error) {
	return l.repo.FindByID(ctx, lessonID)
}

func (l *LessonService) FindCourseLessons(ctx context.Context, courseID domain.ID) ([]domain.Lesson, error) {
	return l.repo.FindCourseLessons(ctx, courseID)
}

func (l *LessonService) CreateCourseLesson(ctx context.Context, courseID domain.ID,
	param port.CreateLessonParam) (domain.Lesson, error) {
	switch param.Type {
	case domain.TheoryLesson:
		//TODO: create empty markdown, if it is a theory and set url
	case domain.VideoLesson:
		if !param.ContentUrl.Valid {
			return domain.Lesson{}, domainErr.ErrLessonContentUrlEmpty
		}
	}

	return l.repo.Create(ctx, domain.Lesson{
		ID:         domain.NewID(),
		CourseID:   courseID,
		Title:      param.Title,
		Type:       param.Type,
		ContentUrl: param.ContentUrl,
	})
}

func (l *LessonService) AddLessonTests(ctx context.Context, lessonID domain.ID, tests []port.CreateTestParam) error {
	for i, test := range tests {
		if test.QuestionString == "" {
			return errors.Wrap(domainErr.ErrLessonTestQuestionEmpty, fmt.Sprintf("test %d error", i))
		}
		if len(test.Options) == 0 {
			return errors.Wrap(domainErr.ErrLessonTestOptionsEmpty, fmt.Sprintf("test %d error", i))
		}
		if test.Level < 0 {
			return errors.Wrap(domainErr.ErrLessonTestInvalidLevel, fmt.Sprintf("test %d error", i))
		}
		if test.Mark < 0 {
			return errors.Wrap(domainErr.ErrLessonTestInvalidMark, fmt.Sprintf("test %d error", i))
		}
	}

	testArray := make([]domain.Test, len(tests))
	for i, test := range tests {
		testArray[i] = domain.Test{
			ID:       domain.NewID(),
			LessonID: lessonID,
			//TODO: add question markdown
			QuestionUrl: "url",
			Options:     test.Options,
			Answer:      test.Answer,
			Level:       test.Level,
			Mark:        test.Mark,
		}
	}

	return l.repo.AddLessonTests(ctx, testArray)
}

func (l *LessonService) DeleteLessonTest(ctx context.Context, testID domain.ID) error {
	return l.repo.DeleteLessonTest(ctx, testID)
}

func (l *LessonService) UpdateLessonTest(ctx context.Context, testID domain.ID,
	param port.UpdateTestParam) (domain.Test, error) {
	return l.repo.UpdateLessonTest(ctx, testID, param)
}

func (l *LessonService) UpdateLessonTheory(ctx context.Context, lessonID domain.ID,
	param port.UpdateTheoryParam) error {
	return l.repo.UpdateLessonTheory(ctx, lessonID, param)
}

func (l *LessonService) UpdateLessonVideo(ctx context.Context, lessonID domain.ID,
	param port.UpdateVideoParam) error {
	return l.repo.UpdateLessonVideo(ctx, lessonID, param)
}

func (l *LessonService) Delete(ctx context.Context, lessonID domain.ID) error {
	return l.repo.Delete(ctx, lessonID)
}
