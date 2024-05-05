package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
)

type CourseService struct {
	repo       port.ICourseRepository
	lessonRepo port.ILessonRepository
	schoolRepo port.ISchoolRepository
	statRepo   port.IStatRepository
}

func NewCourseService(repo port.ICourseRepository, lessonRepo port.ILessonRepository,
	schoolRepo port.ISchoolRepository, statRepo port.IStatRepository) *CourseService {
	return &CourseService{
		repo:       repo,
		lessonRepo: lessonRepo,
		schoolRepo: schoolRepo,
		statRepo:   statRepo,
	}
}

func (c *CourseService) FindAll(ctx context.Context) ([]domain.Course, error) {
	return c.repo.FindAll(ctx)
}

func (c *CourseService) FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error) {
	return c.repo.FindByID(ctx, courseID)
}

func (c *CourseService) FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error) {
	return c.repo.FindStudentCourses(ctx, studentID)
}

func (c *CourseService) FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error) {
	return c.repo.FindTeacherCourses(ctx, teacherID)
}

func (c *CourseService) FindCourseTeachers(ctx context.Context, courseID domain.ID) ([]domain.User, error) {
	return c.repo.FindCourseTeachers(ctx, courseID)
}

func (c *CourseService) IsCourseStudent(ctx context.Context, studentID, courseID domain.ID) (bool, error) {
	return c.repo.IsCourseStudent(ctx, studentID, courseID)
}

func (c *CourseService) IsCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) (bool, error) {
	return c.repo.IsCourseTeacher(ctx, teacherID, courseID)
}

func (c *CourseService) AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error {
	lessons, err := c.lessonRepo.FindCourseLessons(ctx, courseID)
	if err != nil {
		return err
	}

	for _, lesson := range lessons {
		testStats := make([]domain.TestStat, len(lesson.Tests))
		for i, test := range lesson.Tests {
			testStats[i] = domain.TestStat{
				ID:     domain.NewID(),
				TestID: test.ID,
				UserID: studentID,
				Score:  0,
			}
		}

		err = c.statRepo.CreateLessonStat(ctx, domain.LessonStat{
			ID:        domain.NewID(),
			LessonID:  lesson.ID,
			UserID:    studentID,
			Score:     0,
			TestStats: testStats,
		})
		if err != nil {
			return err
		}
	}

	return c.repo.AddCourseStudent(ctx, studentID, courseID)
}

func (c *CourseService) AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error {
	course, err := c.FindByID(ctx, courseID)
	if err != nil {
		return err
	}

	isSchoolTeacher, err := c.schoolRepo.IsSchoolTeacher(ctx, course.SchoolID, teacherID)
	if err != nil {
		return err
	}

	if !isSchoolTeacher {
		return errs.ErrUserIsNotSchoolTeacher
	}

	return c.repo.AddCourseTeacher(ctx, teacherID, courseID)
}

func (c *CourseService) checkCourseLessons(lessons []domain.Lesson) []error {
	var errList []error
	var theoryCount, practiceCount int
	for _, lesson := range lessons {
		err := lesson.Validate()
		if err != nil {
			errList = append(errList, err)
		}

		switch lesson.Type {
		case domain.PracticeLesson:
			practiceCount++
		case domain.TheoryLesson:
			theoryCount++
		case domain.VideoLesson:
			theoryCount++
		}
	}

	if theoryCount == 0 || practiceCount == 0 {
		errList = append(errList, errs.ErrCourseNotEnoughLessons)
	}

	return errList
}

func (c *CourseService) ConfirmDraftCourse(ctx context.Context, courseID domain.ID) []error {
	course, err := c.FindByID(ctx, courseID)
	if err != nil {
		return []error{err}
	}

	if course.Status != domain.CourseDraft {
		return []error{errs.ErrCourseReadyState}
	}

	var errList []error
	lessons, err := c.lessonRepo.FindCourseLessons(ctx, courseID)
	if err != nil {
		return []error{err}
	}

	if checkErrors := c.checkCourseLessons(lessons); checkErrors != nil {
		errList = append(errList, checkErrors...)
	}

	if errList == nil {
		err := c.repo.UpdateStatus(ctx, courseID, domain.CourseReady)
		if err != nil {
			return []error{err}
		}
	}

	return errList
}

func (c *CourseService) PublishReadyCourse(ctx context.Context, courseID domain.ID) error {
	course, err := c.FindByID(ctx, courseID)
	if err != nil {
		return err
	}

	if course.Status != domain.CourseReady {
		return errs.ErrCoursePublishedState
	}

	return c.repo.UpdateStatus(ctx, courseID, domain.CoursePublished)
}

func (c *CourseService) CreateSchoolCourse(ctx context.Context, schoolID domain.ID,
	param port.CreateCourseParam) (domain.Course, error) {
	if param.Price < 0 {
		return domain.Course{}, errs.ErrCourseInvalidPrice
	}
	if param.Level < 1 {
		return domain.Course{}, errs.ErrCourseInvalidLevel
	}

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
	course, err := c.repo.FindByID(ctx, courseID)
	if err != nil {
		return domain.Course{}, err
	}

	if param.Price.Valid {
		if param.Price.Int64 < 0 {
			return domain.Course{}, errs.ErrCourseInvalidPrice
		}
		course.Price = param.Price.Int64
	}
	if param.Name.Valid {
		course.Name = param.Name.String
	}
	if param.Level.Valid {
		if param.Level.Int64 < 1 {
			return domain.Course{}, errs.ErrCourseInvalidLevel
		}
		course.Level = int(param.Level.Int64)
	}
	if param.Language.Valid {
		course.Language = param.Language.String
	}

	return c.repo.Update(ctx, course)
}

func (c *CourseService) Delete(ctx context.Context, courseID domain.ID) error {
	return c.repo.Delete(ctx, courseID)
}
