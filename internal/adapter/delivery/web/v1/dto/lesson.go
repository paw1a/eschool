package dto

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

const (
	LessonDTOTheory   = "theory"
	LessonDTOVideo    = "video"
	LessonDTOPractice = "practice"
)

type CreateLessonDTO struct {
	Title    string          `json:"title" binding:"required"`
	Type     string          `json:"type" binding:"required,oneof=theory video practice"`
	Score    null.Int        `json:"score" binding:"required"`
	Theory   null.String     `json:"theory" binding:"omitempty"`
	VideoUrl null.String     `json:"video_url" binding:"omitempty,url"`
	Tests    []CreateTestDTO `json:"tests" binding:"omitempty"`
}

type CreateTestDTO struct {
	Task    string   `json:"task" binding:"required"`
	Options []string `json:"options" binding:"required"`
	Answer  string   `json:"answer" binding:"required"`
	Level   null.Int `json:"level" binding:"required"`
	Score   null.Int `json:"score" binding:"required"`
}

type UpdateLessonDTO struct {
	Title    null.String     `json:"title" binding:"required"`
	Score    null.Int        `json:"score" binding:"required"`
	Theory   null.String     `json:"theory" binding:"omitempty"`
	VideoUrl null.String     `json:"video_url" binding:"omitempty,url"`
	Tests    []CreateTestDTO `json:"tests" binding:"omitempty"`
}

type PassLessonDTO struct {
	LessonID  string        `json:"lesson_id" binding:"required,uuid"`
	PassTests []PassTestDTO `json:"tests" binding:"omitempty"`
}

type PassTestDTO struct {
	TestID string `json:"test_id" binding:"required,uuid"`
	Answer string `json:"answer" binding:"required"`
}

type LessonDTO struct {
	ID       string `json:"id"`
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Score    int    `json:"score"`
	Type     string `json:"type"`

	TheoryUrl null.String `json:"theory_url" binding:"omitempty"`
	VideoUrl  null.String `json:"video_url" binding:"omitempty"`
	Tests     []TestDTO   `json:"tests" binding:"omitempty"`
}

type TestDTO struct {
	ID       string   `json:"id"`
	LessonID string   `json:"lesson_id"`
	TaskUrl  string   `json:"task_url"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
	Level    int      `json:"level"`
	Score    int      `json:"score"`
}

func NewLessonDTO(lesson domain.Lesson) LessonDTO {
	var lessonType string
	switch lesson.Type {
	case domain.TheoryLesson:
		lessonType = LessonDTOTheory
	case domain.VideoLesson:
		lessonType = LessonDTOVideo
	case domain.PracticeLesson:
		lessonType = LessonDTOPractice
	}

	tests := make([]TestDTO, len(lesson.Tests))
	if lesson.Type == domain.PracticeLesson {
		for i, test := range lesson.Tests {
			tests[i] = NewTestDTO(test)
		}
	}

	return LessonDTO{
		ID:        lesson.ID.String(),
		CourseID:  lesson.CourseID.String(),
		Title:     lesson.Title,
		Score:     lesson.Score,
		Type:      lessonType,
		TheoryUrl: lesson.TheoryUrl,
		VideoUrl:  lesson.VideoUrl,
		Tests:     tests,
	}
}

func NewTestDTO(test domain.Test) TestDTO {
	return TestDTO{
		ID:       test.ID.String(),
		LessonID: test.LessonID.String(),
		TaskUrl:  test.TaskUrl,
		Options:  test.Options,
		Answer:   test.Answer,
		Level:    test.Level,
		Score:    test.Score,
	}
}
