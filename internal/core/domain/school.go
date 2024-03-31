package domain

type School struct {
	ID          int64
	Description string
	Courses     []Course
}
