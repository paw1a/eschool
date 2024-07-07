package port

import "github.com/guregu/null"

type CourseInfo struct {
	Name     string
	Level    int
	Price    int64
	Language string
}

type CreateCourseParam struct {
	Name     string
	Level    int
	Price    int64
	Language string
}

type UpdateCourseParam struct {
	Name     null.String
	Level    null.Int
	Price    null.Int
	Language null.String
}
