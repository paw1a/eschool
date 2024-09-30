package test

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type CourseBuilder struct {
	course domain.Course
}

func NewCourseBuilder() *CourseBuilder {
	return &CourseBuilder{
		course: domain.Course{
			Name:   "course name",
			Status: domain.CourseDraft,
			Level:  3,
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

type CourseInfoBuilder struct {
	info port.CourseInfo
}

func NewCourseInfoBuilder() *CourseInfoBuilder {
	return &CourseInfoBuilder{}
}

func (b *CourseInfoBuilder) WithName(name string) *CourseInfoBuilder {
	b.info.Name = name
	return b
}

func (b *CourseInfoBuilder) WithLevel(level int) *CourseInfoBuilder {
	b.info.Level = level
	return b
}

func (b *CourseInfoBuilder) WithPrice(price int64) *CourseInfoBuilder {
	b.info.Price = price
	return b
}

func (b *CourseInfoBuilder) WithLanguage(language string) *CourseInfoBuilder {
	b.info.Language = language
	return b
}

func (b *CourseInfoBuilder) Build() port.CourseInfo {
	return b.info
}

type CreateCourseParamBuilder struct {
	param port.CreateCourseParam
}

func NewCreateCourseParamBuilder() *CreateCourseParamBuilder {
	return &CreateCourseParamBuilder{
		param: port.CreateCourseParam{
			Name:     "course name",
			Level:    3,
			Price:    0,
			Language: "english",
		},
	}
}

func (b *CreateCourseParamBuilder) WithName(name string) *CreateCourseParamBuilder {
	b.param.Name = name
	return b
}

func (b *CreateCourseParamBuilder) WithLevel(level int) *CreateCourseParamBuilder {
	b.param.Level = level
	return b
}

func (b *CreateCourseParamBuilder) WithPrice(price int64) *CreateCourseParamBuilder {
	b.param.Price = price
	return b
}

func (b *CreateCourseParamBuilder) WithLanguage(language string) *CreateCourseParamBuilder {
	b.param.Language = language
	return b
}

func (b *CreateCourseParamBuilder) Build() port.CreateCourseParam {
	return b.param
}

type UpdateCourseParamBuilder struct {
	param port.UpdateCourseParam
}

func NewUpdateCourseParamBuilder() *UpdateCourseParamBuilder {
	return &UpdateCourseParamBuilder{}
}

func (b *UpdateCourseParamBuilder) WithName(name null.String) *UpdateCourseParamBuilder {
	b.param.Name = name
	return b
}

func (b *UpdateCourseParamBuilder) WithLevel(level null.Int) *UpdateCourseParamBuilder {
	b.param.Level = level
	return b
}

func (b *UpdateCourseParamBuilder) WithPrice(price null.Int) *UpdateCourseParamBuilder {
	b.param.Price = price
	return b
}

func (b *UpdateCourseParamBuilder) WithLanguage(language null.String) *UpdateCourseParamBuilder {
	b.param.Language = language
	return b
}

func (b *UpdateCourseParamBuilder) Build() port.UpdateCourseParam {
	return b.param
}
