package port

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type CreateLessonParam struct {
	Title      string
	Type       domain.LessonType
	ContentUrl null.String
}

type CreateTestParam struct {
	QuestionString string
	Options        []string
	Answer         string
	Level          int
	Mark           int
}

type UpdateTestParam struct {
	QuestionString null.String
	Options        []string
	Answer         null.String
	Level          null.Int
	Mark           null.Int
}

type UpdateTheoryParam struct {
	ContentString string
}

type UpdateVideoParam struct {
	VideoUrl string
}
