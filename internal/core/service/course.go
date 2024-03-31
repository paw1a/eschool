package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

type CourseService struct {
	repo port.ICourseRepository
}

func NewCourseService(repo port.ICourseRepository) *CourseService {
	return &CourseService{
		repo: repo,
	}
}

func (c *CourseService) FindAll(ctx context.Context) ([]domain.Course, error) {
	return c.repo.FindAll(ctx)
}

func (c *CourseService) FindByID(ctx context.Context, courseID int64) (domain.Course, error) {
	return c.repo.FindByID(ctx, courseID)
}

func (c *CourseService) FindCourseInfo(ctx context.Context, courseID int64) (dto.CourseInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) FindStudentCourses(ctx context.Context, studentID int64) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) FindTeacherCourses(ctx context.Context, teacherID int64) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) AddCourseStudent(ctx context.Context, studentID, courseID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) AddCourseTeacher(ctx context.Context, teacherID, courseID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) AddCourseLesson(ctx context.Context, courseID, lessonID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) DeleteCourseLesson(ctx context.Context, courseID, lessonID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) ConfirmDraftCourse(ctx context.Context, courseID int64) []error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) PublishReadyCourse(ctx context.Context, courseID int64) error {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) Create(ctx context.Context, courseDTO dto.CreateCourseDTO) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) Update(ctx context.Context, courseID int64,
	courseDTO dto.UpdateCourseDTO) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CourseService) Delete(ctx context.Context, courseID int64) error {
	//TODO implement me
	panic("implement me")
}
