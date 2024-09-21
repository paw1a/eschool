package domain

type Review struct {
	ID       ID
	UserID   ID
	CourseID ID
	Text     string
	Rating   int64
}
