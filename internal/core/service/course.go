package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	domainErr "github.com/paw1a/eschool/internal/core/errors"
	"github.com/paw1a/eschool/internal/core/port"
)

type CourseService struct {
	repo       port.ICourseRepository
	lessonRepo port.ILessonRepository
}

func NewCourseService(repo port.ICourseRepository, lessonRepo port.ILessonRepository) *CourseService {
	return &CourseService{
		repo:       repo,
		lessonRepo: lessonRepo,
	}
}

func (c *CourseService) FindAll(ctx context.Context) ([]domain.Course, error) {
	return c.repo.FindAll(ctx)
}

func (c *CourseService) FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error) {
	return c.repo.FindByID(ctx, courseID)
}

func (c *CourseService) FindCourseInfo(ctx context.Context, courseID domain.ID) (port.CourseInfo, error) {
	return c.repo.FindCourseInfo(ctx, courseID)
}

func (c *CourseService) FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error) {
	return c.repo.FindStudentCourses(ctx, studentID)
}

func (c *CourseService) FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error) {
	return c.repo.FindTeacherCourses(ctx, teacherID)
}

func (c *CourseService) AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error {
	return c.repo.AddCourseStudent(ctx, studentID, courseID)
}

func (c *CourseService) AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error {
	return c.repo.AddCourseTeacher(ctx, teacherID, courseID)
}

func (c *CourseService) AddCourseLesson(ctx context.Context, courseID, lessonID domain.ID) error {
	return c.repo.AddCourseLesson(ctx, courseID, lessonID)
}

func (c *CourseService) DeleteCourseLesson(ctx context.Context, courseID, lessonID domain.ID) error {
	return c.repo.DeleteCourseLesson(ctx, courseID, lessonID)
}

func (c *CourseService) ConfirmDraftCourse(ctx context.Context, courseID domain.ID) []error {
	course, err := c.FindByID(ctx, courseID)
	if err != nil {
		return []error{err}
	}

	if course.Status != domain.CourseDraft {
		return []error{domainErr.ErrCourseReadyState}
	}

	var errs []error

	var theoryCount, practiceCount int
	lessons, err := c.lessonRepo.FindCourseLessons(ctx, courseID)
	if err != nil {
		return []error{err}
	}

	for _, lesson := range lessons {
		switch lesson.Type {
		case domain.TheoryLesson:
			fallthrough
		case domain.VideoLesson:
			theoryCount++
		case domain.PracticeLesson:
			practiceCount++
		}
	}

	if theoryCount == 0 || practiceCount == 0 {
		errs = append(errs, domainErr.ErrCourseNotEnoughLessons)
	}

	for _, lesson := range lessons {
		if lesson.Mark <= 0 {
			errs = append(errs, domainErr.ErrCourseLessonInvalidMark)
		}
		switch lesson.Type {
		case domain.PracticeLesson:
			tests, err := c.lessonRepo.FindLessonTests(ctx, lesson.ID)
			if err != nil {
				return []error{err}
			}

			if len(tests) == 0 {
				errs = append(errs, domainErr.ErrCoursePracticeLessonTestsEmpty)
			}
		case domain.TheoryLesson:
			//TODO: load question markdown to string
			var questionString string
			if len(questionString) == 0 {
				errs = append(errs, domainErr.ErrCourseTheoryLessonEmpty)
			}
		case domain.VideoLesson:
			if !lesson.ContentUrl.Valid {
				errs = append(errs, domainErr.ErrCourseContentUrlInvalid)
			}
		}
	}

	if errs == nil {
		err := c.repo.UpdateStatus(ctx, courseID, domain.CourseReady)
		if err != nil {
			return []error{err}
		}
	}

	return errs
}

func (c *CourseService) PublishReadyCourse(ctx context.Context, courseID domain.ID) error {
	course, err := c.FindByID(ctx, courseID)
	if err != nil {
		return err
	}

	if course.Status != domain.CourseReady {
		return domainErr.ErrCoursePublishedState
	}

	return c.repo.UpdateStatus(ctx, courseID, domain.CoursePublished)
}

func (c *CourseService) CreateSchoolCourse(ctx context.Context, schoolID domain.ID,
	param port.CreateCourseParam) (domain.Course, error) {
	return c.repo.Create(ctx, domain.Course{
		ID:       domain.NewID(),
		SchoolID: schoolID,
		Name:     param.Name,
		Level:    param.Level,
		Price:    param.Price,
		Language: param.Language,
		Status:   domain.CourseDraft,
	})
}

func (c *CourseService) Update(ctx context.Context, courseID domain.ID,
	param port.UpdateCourseParam) (domain.Course, error) {
	return c.repo.Update(ctx, courseID, param)
}

func (c *CourseService) Delete(ctx context.Context, courseID domain.ID) error {
	return c.repo.Delete(ctx, courseID)
}
