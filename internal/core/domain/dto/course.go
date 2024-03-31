package dto

import "github.com/guregu/null"

type CourseInfo struct {
	Name     string
	Level    int
	Price    int64
	Language string
}

type CreateCourseDTO struct {
	Name     string
	Level    int
	Price    int64
	Language string
}

type UpdateCourseDTO struct {
	Name     null.String
	Level    null.Int
	Price    null.Int
	Language null.String
}
