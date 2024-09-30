package test

import (
	"github.com/paw1a/eschool/internal/core/domain"
)

// CourseBuilder - билдер для структуры domain.Course
type CourseBuilder struct {
	course domain.Course
}

func NewCourseBuilder() *CourseBuilder {
	return &CourseBuilder{
		course: domain.Course{
			ID:       domain.NewID(),
			SchoolID: domain.NewID(),
			Name:     "course name",
			Level:    3,
			Price:    1000,
			Language: "english",
			Status:   domain.CourseDraft,
		},
	}
}

func (b *CourseBuilder) WithID(id domain.ID) *CourseBuilder {
	b.course.ID = id
	return b
}

func (b *CourseBuilder) WithSchoolID(schoolID domain.ID) *CourseBuilder {
	b.course.SchoolID = schoolID
	return b
}

func (b *CourseBuilder) WithName(name string) *CourseBuilder {
	b.course.Name = name
	return b
}

func (b *CourseBuilder) WithLevel(level int) *CourseBuilder {
	b.course.Level = level
	return b
}

func (b *CourseBuilder) WithPrice(price int64) *CourseBuilder {
	b.course.Price = price
	return b
}

func (b *CourseBuilder) WithLanguage(language string) *CourseBuilder {
	b.course.Language = language
	return b
}

func (b *CourseBuilder) WithStatus(status domain.CourseStatus) *CourseBuilder {
	b.course.Status = status
	return b
}

func (b *CourseBuilder) Build() domain.Course {
	return b.course
}
