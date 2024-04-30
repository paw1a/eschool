package dto

import "github.com/paw1a/eschool/internal/core/domain"

type CreateLessonDTO struct {
	Title string
	Type  domain.LessonType
}

type CreateTestDTO struct {
	QuestionString string
	Options        []string
	Answer         string
	Level          int
	Mark           int
}

type UpdateTestDTO struct {
	QuestionString string
	Options        []string
	Answer         string
	Level          int
	Mark           int
}

type UpdateTheoryDTO struct {
	ContentString string
}

type UpdateVideoDTO struct {
	VideoUrl string
}
