package domain

type CourseStatus int

const (
	CourseDraft CourseStatus = iota
	CourseReady
	CoursePublished
)

type Course struct {
	ID       ID
	SchoolID ID
	Name     string
	Level    int
	Price    int64
	Language string
	Status   CourseStatus
}
