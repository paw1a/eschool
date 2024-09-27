package test

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type LessonBuilder struct {
	lesson domain.Lesson
}

func NewLessonBuilder() *LessonBuilder {
	return &LessonBuilder{
		lesson: domain.Lesson{
			Title:     "title",
			Score:     10,
			Type:      domain.TheoryLesson,
			TheoryUrl: null.String{},
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
			TaskUrl: "url",
			Options: []string{"opt1", "opt2"},
			Answer:  "opt1",
			Level:   3,
			Score:   10,
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

type CreateTheoryParamBuilder struct {
	param port.CreateTheoryParam
}

func NewCreateTheoryParamBuilder() *CreateTheoryParamBuilder {
	return &CreateTheoryParamBuilder{
		param: port.CreateTheoryParam{
			Title:  "title",
			Score:  10,
			Theory: "text",
		},
	}
}

func (b *CreateTheoryParamBuilder) WithTitle(title string) *CreateTheoryParamBuilder {
	b.param.Title = title
	return b
}

func (b *CreateTheoryParamBuilder) WithScore(score int) *CreateTheoryParamBuilder {
	b.param.Score = score
	return b
}

func (b *CreateTheoryParamBuilder) WithTheory(theory string) *CreateTheoryParamBuilder {
	b.param.Theory = theory
	return b
}

func (b *CreateTheoryParamBuilder) Build() port.CreateTheoryParam {
	return b.param
}

type CreateVideoParamBuilder struct {
	param port.CreateVideoParam
}

func NewCreateVideoParamBuilder() *CreateVideoParamBuilder {
	return &CreateVideoParamBuilder{
		param: port.CreateVideoParam{
			Title:    "title",
			Score:    20,
			VideoUrl: "url",
		},
	}
}

func (b *CreateVideoParamBuilder) WithTitle(title string) *CreateVideoParamBuilder {
	b.param.Title = title
	return b
}

func (b *CreateVideoParamBuilder) WithScore(score int) *CreateVideoParamBuilder {
	b.param.Score = score
	return b
}

func (b *CreateVideoParamBuilder) WithVideoUrl(videoUrl string) *CreateVideoParamBuilder {
	b.param.VideoUrl = videoUrl
	return b
}

func (b *CreateVideoParamBuilder) Build() port.CreateVideoParam {
	return b.param
}

type CreatePracticeParamBuilder struct {
	param port.CreatePracticeParam
}

func NewCreatePracticeParamBuilder() *CreatePracticeParamBuilder {
	return &CreatePracticeParamBuilder{
		param: port.CreatePracticeParam{
			Title: "title",
			Score: 10,
			Tests: []port.CreateTestParam{NewCreateTestParamBuilder().Build()},
		},
	}
}

func (b *CreatePracticeParamBuilder) WithTitle(title string) *CreatePracticeParamBuilder {
	b.param.Title = title
	return b
}

func (b *CreatePracticeParamBuilder) WithScore(score int) *CreatePracticeParamBuilder {
	b.param.Score = score
	return b
}

func (b *CreatePracticeParamBuilder) WithTests(tests []port.CreateTestParam) *CreatePracticeParamBuilder {
	b.param.Tests = tests
	return b
}

func (b *CreatePracticeParamBuilder) Build() port.CreatePracticeParam {
	return b.param
}

type UpdateTheoryParamBuilder struct {
	param port.UpdateTheoryParam
}

func NewUpdateTheoryParamBuilder() *UpdateTheoryParamBuilder {
	return &UpdateTheoryParamBuilder{
		param: port.UpdateTheoryParam{},
	}
}

func (b *UpdateTheoryParamBuilder) WithTitle(title null.String) *UpdateTheoryParamBuilder {
	b.param.Title = title
	return b
}

func (b *UpdateTheoryParamBuilder) WithScore(score null.Int) *UpdateTheoryParamBuilder {
	b.param.Score = score
	return b
}

func (b *UpdateTheoryParamBuilder) WithTheory(theory null.String) *UpdateTheoryParamBuilder {
	b.param.Theory = theory
	return b
}

func (b *UpdateTheoryParamBuilder) Build() port.UpdateTheoryParam {
	return b.param
}

type UpdateVideoParamBuilder struct {
	param port.UpdateVideoParam
}

func NewUpdateVideoParamBuilder() *UpdateVideoParamBuilder {
	return &UpdateVideoParamBuilder{
		param: port.UpdateVideoParam{},
	}
}

func (b *UpdateVideoParamBuilder) WithTitle(title null.String) *UpdateVideoParamBuilder {
	b.param.Title = title
	return b
}

func (b *UpdateVideoParamBuilder) WithScore(score null.Int) *UpdateVideoParamBuilder {
	b.param.Score = score
	return b
}

func (b *UpdateVideoParamBuilder) WithVideoUrl(videoUrl null.String) *UpdateVideoParamBuilder {
	b.param.VideoUrl = videoUrl
	return b
}

func (b *UpdateVideoParamBuilder) Build() port.UpdateVideoParam {
	return b.param
}

type UpdatePracticeParamBuilder struct {
	param port.UpdatePracticeParam
}

func NewUpdatePracticeParamBuilder() *UpdatePracticeParamBuilder {
	return &UpdatePracticeParamBuilder{
		param: port.UpdatePracticeParam{
			Title: null.String{},
			Score: null.Int{},
			Tests: []port.UpdateTestParam{NewUpdateTestParamBuilder().Build()},
		},
	}
}

func (b *UpdatePracticeParamBuilder) WithTitle(title null.String) *UpdatePracticeParamBuilder {
	b.param.Title = title
	return b
}

func (b *UpdatePracticeParamBuilder) WithScore(score null.Int) *UpdatePracticeParamBuilder {
	b.param.Score = score
	return b
}

func (b *UpdatePracticeParamBuilder) WithTests(tests []port.UpdateTestParam) *UpdatePracticeParamBuilder {
	b.param.Tests = tests
	return b
}

func (b *UpdatePracticeParamBuilder) Build() port.UpdatePracticeParam {
	return b.param
}

// CreateTestParamBuilder is a builder for CreateTestParam.
type CreateTestParamBuilder struct {
	param port.CreateTestParam
}

// NewCreateTestParamBuilder creates a new instance of CreateTestParamBuilder.
func NewCreateTestParamBuilder() *CreateTestParamBuilder {
	return &CreateTestParamBuilder{
		param: port.CreateTestParam{
			Task:    "task",
			Options: []string{"option1", "option2"},
			Answer:  "option1",
			Level:   3,
			Score:   10,
		},
	}
}

// WithTask sets the task of the test.
func (b *CreateTestParamBuilder) WithTask(task string) *CreateTestParamBuilder {
	b.param.Task = task
	return b
}

// WithOptions sets the options of the test.
func (b *CreateTestParamBuilder) WithOptions(options []string) *CreateTestParamBuilder {
	b.param.Options = options
	return b
}

// WithAnswer sets the answer of the test.
func (b *CreateTestParamBuilder) WithAnswer(answer string) *CreateTestParamBuilder {
	b.param.Answer = answer
	return b
}

// WithLevel sets the level of the test.
func (b *CreateTestParamBuilder) WithLevel(level int) *CreateTestParamBuilder {
	b.param.Level = level
	return b
}

// WithScore sets the score of the test.
func (b *CreateTestParamBuilder) WithScore(score int) *CreateTestParamBuilder {
	b.param.Score = score
	return b
}

// Build constructs the CreateTestParam object.
func (b *CreateTestParamBuilder) Build() port.CreateTestParam {
	return b.param
}

// UpdateTestParamBuilder is a builder for UpdateTestParam.
type UpdateTestParamBuilder struct {
	param port.UpdateTestParam
}

// NewUpdateTestParamBuilder creates a new instance of UpdateTestParamBuilder.
func NewUpdateTestParamBuilder() *UpdateTestParamBuilder {
	return &UpdateTestParamBuilder{
		param: port.UpdateTestParam{
			Task:    "task",
			Options: []string{"option1", "option2"},
			Answer:  "option1",
			Level:   2,
			Score:   10,
		},
	}
}

// WithTask sets the task of the test.
func (b *UpdateTestParamBuilder) WithTask(task string) *UpdateTestParamBuilder {
	b.param.Task = task
	return b
}

// WithOptions sets the options of the test.
func (b *UpdateTestParamBuilder) WithOptions(options []string) *UpdateTestParamBuilder {
	b.param.Options = options
	return b
}

// WithAnswer sets the answer of the test.
func (b *UpdateTestParamBuilder) WithAnswer(answer string) *UpdateTestParamBuilder {
	b.param.Answer = answer
	return b
}

// WithLevel sets the level of the test.
func (b *UpdateTestParamBuilder) WithLevel(level int) *UpdateTestParamBuilder {
	b.param.Level = level
	return b
}

// WithScore sets the score of the test.
func (b *UpdateTestParamBuilder) WithScore(score int) *UpdateTestParamBuilder {
	b.param.Score = score
	return b
}

// Build constructs the UpdateTestParam object.
func (b *UpdateTestParamBuilder) Build() port.UpdateTestParam {
	return b.param
}
