package domain

import "github.com/guregu/null"

type LessonType int

const (
	TheoryLesson LessonType = iota
	PracticeLesson
	VideoLesson
)

type Lesson struct {
	ID         ID
	CourseID   ID
	Title      string
	Mark       int
	Type       LessonType
	ContentUrl null.String
}

type Test struct {
	ID          ID
	LessonID    ID
	QuestionUrl string
	Options     []string
	Answer      string
	Level       int
	Mark        int
}

type LessonStat struct {
	ID       ID
	LessonID ID
	UserID   ID
	Mark     int
}

type TestStat struct {
	ID     ID
	TestID ID
	UserID ID
	Mark   int
}
