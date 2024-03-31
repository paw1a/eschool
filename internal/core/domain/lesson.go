package domain

import "github.com/guregu/null"

type LessonType int

const (
	Theory LessonType = iota
	Practice
	Video
)

type Lesson struct {
	ID         int64
	Title      string
	Type       LessonType
	Tests      []Test
	ContentUrl null.String
}

type Test struct {
	ID          int64
	QuestionUrl string
	Options     []string
	Answer      string
	Mark        int
}

type LessonStat struct {
}

type TestStat struct {
}
