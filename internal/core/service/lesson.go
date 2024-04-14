package service

import (
	"context"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"strings"
)

type LessonService struct {
	repo    port.ILessonRepository
	storage port.IObjectStorage
}

func NewLessonService(repo port.ILessonRepository, storage port.IObjectStorage) *LessonService {
	return &LessonService{
		repo:    repo,
		storage: storage,
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

func (l *LessonService) CreateTheoryLesson(ctx context.Context, courseID domain.ID,
	param port.CreateTheoryParam) (domain.Lesson, error) {
	lessonID := domain.NewID()
	url, err := l.storage.SaveFile(ctx, domain.File{
		Name:   lessonID.String() + ".md",
		Path:   "course/" + courseID.String(),
		Reader: strings.NewReader(param.Theory),
	})
	if err != nil {
		return domain.Lesson{}, err
	}

	lesson := domain.Lesson{
		ID:        lessonID,
		CourseID:  courseID,
		Title:     param.Title,
		Score:     param.Score,
		Type:      domain.TheoryLesson,
		TheoryUrl: null.StringFrom(url.String()),
	}
	if err := lesson.Validate(); err != nil {
		return domain.Lesson{}, err
	}

	return l.repo.Create(ctx, lesson)
}

func (l *LessonService) CreateVideoLesson(ctx context.Context, courseID domain.ID,
	param port.CreateVideoParam) (domain.Lesson, error) {
	lessonID := domain.NewID()
	lesson := domain.Lesson{
		ID:       lessonID,
		CourseID: courseID,
		Title:    param.Title,
		Score:    param.Score,
		Type:     domain.VideoLesson,
		VideoUrl: null.StringFrom(param.VideoUrl),
	}
	if err := lesson.Validate(); err != nil {
		return domain.Lesson{}, err
	}

	return l.repo.Create(ctx, lesson)
}

func (l *LessonService) CreatePracticeLesson(ctx context.Context, courseID domain.ID,
	param port.CreatePracticeParam) (domain.Lesson, error) {

	lessonID := domain.NewID()
	tests := make([]domain.Test, len(param.Tests))
	for i, test := range param.Tests {
		tests[i] = domain.Test{
			ID:       domain.NewID(),
			LessonID: lessonID,
			TaskUrl:  "undefined",
			Options:  test.Options,
			Answer:   test.Answer,
			Level:    test.Level,
			Score:    test.Score,
		}
	}

	lesson := domain.Lesson{
		ID:       lessonID,
		CourseID: courseID,
		Title:    param.Title,
		Score:    param.Score,
		Type:     domain.PracticeLesson,
		Tests:    tests,
	}
	if err := lesson.Validate(); err != nil {
		return domain.Lesson{}, err
	}

	for i := range tests {
		url, err := l.storage.SaveFile(ctx, domain.File{
			Name:   tests[i].ID.String() + ".md",
			Path:   "course/" + courseID.String() + "/" + lessonID.String(),
			Reader: strings.NewReader(param.Tests[i].Task),
		})
		if err != nil {
			return domain.Lesson{}, err
		}
		tests[i].TaskUrl = url.String()
	}

	return l.repo.Create(ctx, lesson)
}

func (l *LessonService) UpdateTheoryLesson(ctx context.Context, lessonID domain.ID,
	param port.UpdateTheoryParam) (domain.Lesson, error) {
	lesson, err := l.repo.FindByID(ctx, lessonID)
	if err != nil {
		return domain.Lesson{}, err
	}

	if param.Theory.Valid {
		url, err := l.storage.SaveFile(ctx, domain.File{
			Name:   lesson.ID.String() + ".md",
			Path:   "course/" + lesson.CourseID.String(),
			Reader: strings.NewReader(param.Theory.String),
		})
		if err != nil {
			return domain.Lesson{}, err
		}
		lesson.TheoryUrl = null.StringFrom(url.String())
	}
	if param.Score.Valid {
		lesson.Score = int(param.Score.Int64)
	}
	if param.Title.Valid {
		lesson.Title = param.Title.String
	}

	return l.repo.Update(ctx, lesson)
}

func (l *LessonService) UpdateVideoLesson(ctx context.Context, lessonID domain.ID,
	param port.UpdateVideoParam) (domain.Lesson, error) {
	lesson, err := l.repo.FindByID(ctx, lessonID)
	if err != nil {
		return domain.Lesson{}, err
	}

	if param.VideoUrl.Valid {
		lesson.VideoUrl = null.StringFrom(param.VideoUrl.String)
	}
	if param.Score.Valid {
		lesson.Score = int(param.Score.Int64)
	}
	if param.Title.Valid {
		lesson.Title = param.Title.String
	}

	return l.repo.Update(ctx, lesson)
}

func (l *LessonService) UpdatePracticeLesson(ctx context.Context, lessonID domain.ID,
	param port.UpdatePracticeParam) (domain.Lesson, error) {

	lesson, err := l.repo.FindByID(ctx, lessonID)
	if err != nil {
		return domain.Lesson{}, err
	}

	if param.Score.Valid {
		lesson.Score = int(param.Score.Int64)
	}
	if param.Title.Valid {
		lesson.Title = param.Title.String
	}

	tests := make([]domain.Test, len(param.Tests))
	for i, test := range param.Tests {
		tests[i] = domain.Test{
			ID:       domain.NewID(),
			LessonID: lessonID,
			TaskUrl:  "undefined",
			Options:  test.Options,
			Answer:   test.Answer,
			Level:    test.Level,
			Score:    test.Score,
		}
	}
	lesson.Tests = tests
	if err := lesson.Validate(); err != nil {
		return domain.Lesson{}, err
	}

	for i := range lesson.Tests {
		url, err := l.storage.SaveFile(ctx, domain.File{
			Name:   tests[i].ID.String() + ".md",
			Path:   "course/" + lesson.CourseID.String() + "/" + lessonID.String(),
			Reader: strings.NewReader(param.Tests[i].Task),
		})
		if err != nil {
			return domain.Lesson{}, err
		}
		lesson.Tests[i].TaskUrl = url.String()
	}

	return l.repo.Update(ctx, lesson)
}

func (l *LessonService) Delete(ctx context.Context, lessonID domain.ID) error {
	return l.repo.Delete(ctx, lessonID)
}
