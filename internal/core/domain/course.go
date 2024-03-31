package domain

type CourseStatus int

const (
	Draft CourseStatus = iota
	Ready
	Published
)

type Course struct {
	ID       int64
	Name     string
	Level    int
	Price    int64
	Language string
	Status   CourseStatus
	Lessons  []Lesson
	Reviews  []Review
}
