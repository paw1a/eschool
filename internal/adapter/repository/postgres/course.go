package repository

import (
	"context"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/domain"
)

type CourseRepository struct {
}

func (c *CourseRepository) FindAll(ctx context.Context) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) FindByID(ctx context.Context, courseID int64) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) FindCourseInfo(ctx context.Context, courseID int64) (dto.CourseInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) FindStudentCourses(ctx context.Context, studentID int64) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) FindTeacherCourses(ctx context.Context, teacherID int64) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) AddCourseStudent(ctx context.Context, studentID, courseID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) AddCourseTeacher(ctx context.Context, teacherID, courseID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) AddCourseLesson(ctx context.Context, courseID, lessonID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) DeleteCourseLesson(ctx context.Context, courseID, lessonID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) Create(ctx context.Context, courseDTO dto.CreateCourseDTO) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) Update(ctx context.Context, courseID int64, courseDTO dto.UpdateCourseDTO) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseRepository) Delete(ctx context.Context, courseID int64) error {
	//TODO implement me
	panic("implement me")
}
