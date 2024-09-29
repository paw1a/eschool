package test

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type LessonBuilder struct {
	lesson domain.Lesson
}

func NewLessonBuilder() *LessonBuilder {
	return &LessonBuilder{
		lesson: domain.Lesson{
			ID:        domain.NewID(),
			Title:     "title",
			Score:     10,
			Type:      domain.TheoryLesson,
			TheoryUrl: null.StringFrom("url"),
			VideoUrl:  null.String{},
			Tests:     []domain.Test{NewTestBuilder().Build(), NewTestBuilder().Build()},
		},
	}
}

func (b *LessonBuilder) WithID(id domain.ID) *LessonBuilder {
	b.lesson.ID = id
	return b
}

func (b *LessonBuilder) WithCourseID(courseID domain.ID) *LessonBuilder {
	b.lesson.CourseID = courseID
	return b
}

func (b *LessonBuilder) WithTitle(title string) *LessonBuilder {
	b.lesson.Title = title
	return b
}

func (b *LessonBuilder) WithScore(score int) *LessonBuilder {
	b.lesson.Score = score
	return b
}

func (b *LessonBuilder) WithType(lessonType domain.LessonType) *LessonBuilder {
	b.lesson.Type = lessonType
	return b
}

func (b *LessonBuilder) WithTheoryUrl(theoryUrl null.String) *LessonBuilder {
	b.lesson.TheoryUrl = theoryUrl
	return b
}

func (b *LessonBuilder) WithVideoUrl(videoUrl null.String) *LessonBuilder {
	b.lesson.VideoUrl = videoUrl
	return b
}

func (b *LessonBuilder) WithTests(tests []domain.Test) *LessonBuilder {
	b.lesson.Tests = tests
	return b
}

func (b *LessonBuilder) Build() domain.Lesson {
	return b.lesson
}

type TestBuilder struct {
	test domain.Test
}

func NewTestBuilder() *TestBuilder {
	return &TestBuilder{
		test: domain.Test{
			ID:       domain.NewID(),
			LessonID: domain.NewID(),
			TaskUrl:  "url",
			Options:  []string{"opt1", "opt2"},
			Answer:   "opt1",
			Level:    3,
			Score:    10,
		},
	}
}

func (b *TestBuilder) WithID(id domain.ID) *TestBuilder {
	b.test.ID = id
	return b
}

func (b *TestBuilder) WithLessonID(lessonID domain.ID) *TestBuilder {
	b.test.LessonID = lessonID
	return b
}

func (b *TestBuilder) WithTaskUrl(taskUrl string) *TestBuilder {
	b.test.TaskUrl = taskUrl
	return b
}

func (b *TestBuilder) WithOptions(options []string) *TestBuilder {
	b.test.Options = options
	return b
}

func (b *TestBuilder) WithAnswer(answer string) *TestBuilder {
	b.test.Answer = answer
	return b
}

func (b *TestBuilder) WithLevel(level int) *TestBuilder {
	b.test.Level = level
	return b
}

func (b *TestBuilder) WithScore(score int) *TestBuilder {
	b.test.Score = score
	return b
}

func (b *TestBuilder) Build() domain.Test {
	return b.test
}
