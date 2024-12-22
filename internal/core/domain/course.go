package domain

type CourseStatus int

const (
	CourseDraft CourseStatus = iota
	CourseReady
	CoursePublished
)

type Course struct {
	ID          ID
	SchoolID    ID
	Name        string
	Description string
	Level       int
	Price       int64
	Language    string
	ImageUrl    string
	Status      CourseStatus
}
