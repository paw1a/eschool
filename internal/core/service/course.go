package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/zap"
)

type CourseService struct {
	repo       port.ICourseRepository
	lessonRepo port.ILessonRepository
	schoolRepo port.ISchoolRepository
	statRepo   port.IStatRepository
	logger     *zap.Logger
}

func NewCourseService(repo port.ICourseRepository, lessonRepo port.ILessonRepository,
	schoolRepo port.ISchoolRepository, statRepo port.IStatRepository, logger *zap.Logger) *CourseService {
	return &CourseService{
		repo:       repo,
		lessonRepo: lessonRepo,
		schoolRepo: schoolRepo,
		statRepo:   statRepo,
		logger:     logger,
	}
}

func (c *CourseService) FindAll(ctx context.Context) ([]domain.Course, error) {
	courses, err := c.repo.FindAll(ctx)
	if err != nil {
		c.logger.Error("failed to find all courses", zap.Error(err))
		return nil, err
	}
	return courses, nil
}

func (c *CourseService) FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error) {
	course, err := c.repo.FindByID(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return domain.Course{}, err
	}
	return course, nil
}

func (c *CourseService) FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error) {
	courses, err := c.repo.FindStudentCourses(ctx, studentID)
	if err != nil {
		c.logger.Error("failed to find student courses", zap.Error(err),
			zap.String("userID", studentID.String()))
		return nil, err
	}
	return courses, nil
}

func (c *CourseService) FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error) {
	courses, err := c.repo.FindTeacherCourses(ctx, teacherID)
	if err != nil {
		c.logger.Error("failed to find teacher courses", zap.Error(err),
			zap.String("userID", teacherID.String()))
		return nil, err
	}
	return courses, nil
}

func (c *CourseService) FindCourseTeachers(ctx context.Context, courseID domain.ID) ([]domain.User, error) {
	teachers, err := c.repo.FindCourseTeachers(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course teachers", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return nil, err
	}
	return teachers, nil
}

func (c *CourseService) IsCourseStudent(ctx context.Context, studentID, courseID domain.ID) (bool, error) {
	flag, err := c.repo.IsCourseStudent(ctx, studentID, courseID)
	if err != nil {
		c.logger.Error("failed to check if user is a course student", zap.Error(err),
			zap.String("userID", studentID.String()), zap.String("courseID", courseID.String()))
		return false, err
	}
	return flag, nil
}

func (c *CourseService) IsCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) (bool, error) {
	flag, err := c.repo.IsCourseTeacher(ctx, teacherID, courseID)
	if err != nil {
		c.logger.Error("failed to check if user is a course student", zap.Error(err),
			zap.String("userID", teacherID.String()), zap.String("courseID", courseID.String()))
		return false, err
	}
	return flag, nil
}

func (c *CourseService) AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error {
	lessons, err := c.lessonRepo.FindCourseLessons(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course lessons", zap.Error(err),
			zap.String("userID", studentID.String()), zap.String("courseID", courseID.String()))
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
			c.logger.Error("failed to create lesson statistics entry", zap.Error(err),
				zap.String("lessonID", lesson.ID.String()), zap.String("courseID", courseID.String()))
			return err
		}
	}

	err = c.repo.AddCourseStudent(ctx, studentID, courseID)
	if err != nil {
		c.logger.Error("failed to add course student", zap.Error(err),
			zap.String("userID", studentID.String()), zap.String("courseID", courseID.String()))
		return err
	}

	c.logger.Info("course student is successfully added",
		zap.String("userID", studentID.String()), zap.String("courseID", courseID.String()))
	return nil
}

func (c *CourseService) AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error {
	course, err := c.FindByID(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return err
	}

	isSchoolTeacher, err := c.schoolRepo.IsSchoolTeacher(ctx, course.SchoolID, teacherID)
	if err != nil {
		c.logger.Error("failed to check if user is a school teacher", zap.Error(err),
			zap.String("userID", teacherID.String()), zap.String("courseID", courseID.String()))
		return err
	}

	if !isSchoolTeacher {
		c.logger.Error("user is not a school teacher")
		return errs.ErrUserIsNotSchoolTeacher
	}

	err = c.repo.AddCourseTeacher(ctx, teacherID, courseID)
	if err != nil {
		c.logger.Error("failed to add course teacher", zap.Error(err),
			zap.String("userID", teacherID.String()), zap.String("courseID", courseID.String()))
		return err
	}

	c.logger.Info("course teacher is successfully added",
		zap.String("userID", teacherID.String()), zap.String("courseID", courseID.String()))
	return nil
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
		c.logger.Error("failed to find course by id", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return []error{err}
	}

	if course.Status != domain.CourseDraft {
		c.logger.Error("course is not draft to make it ready",
			zap.String("courseID", courseID.String()))
		return []error{errs.ErrCourseReadyState}
	}

	var errList []error
	lessons, err := c.lessonRepo.FindCourseLessons(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course lessons", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return []error{err}
	}

	if checkErrors := c.checkCourseLessons(lessons); checkErrors != nil {
		errList = append(errList, checkErrors...)
	}

	if errList == nil {
		err := c.repo.UpdateStatus(ctx, courseID, domain.CourseReady)
		if err != nil {
			c.logger.Error("failed to update course status", zap.Error(err),
				zap.String("courseID", courseID.String()))
			return []error{err}
		}
	} else {
		c.logger.Error("failed to confirm draft course", zap.Error(errList[0]),
			zap.String("courseID", courseID.String()))
		return errList
	}

	c.logger.Info("course status is successfully changed to ready",
		zap.String("courseID", courseID.String()))
	return nil
}

func (c *CourseService) PublishReadyCourse(ctx context.Context, courseID domain.ID) error {
	course, err := c.FindByID(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course by id", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return err
	}

	if course.Status != domain.CourseReady {
		c.logger.Error("course is not ready to be published", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return errs.ErrCoursePublishedState
	}

	err = c.repo.UpdateStatus(ctx, courseID, domain.CoursePublished)
	if err != nil {
		c.logger.Error("failed to update course status",
			zap.String("courseID", courseID.String()))
		return err
	}

	c.logger.Info("course status is successfully changed to published",
		zap.String("courseID", courseID.String()))
	return nil
}

func (c *CourseService) CreateSchoolCourse(ctx context.Context, schoolID domain.ID,
	param port.CreateCourseParam) (domain.Course, error) {
	if param.Price < 0 {
		c.logger.Error("failed to create course, price is < 0")
		return domain.Course{}, errs.ErrCourseInvalidPrice
	}
	if param.Level < 1 {
		c.logger.Error("failed to create course, level is < 1")
		return domain.Course{}, errs.ErrCourseInvalidLevel
	}

	course, err := c.repo.Create(ctx, domain.Course{
		ID:       domain.NewID(),
		SchoolID: schoolID,
		Name:     param.Name,
		Level:    param.Level,
		Price:    param.Price,
		Language: param.Language,
		Status:   domain.CourseDraft,
	})
	if err != nil {
		c.logger.Error("failed to create course", zap.Error(err))
		return domain.Course{}, err
	}

	c.logger.Info("course is successfully created",
		zap.String("courseID", course.ID.String()), zap.String("schoolID", schoolID.String()))
	return course, nil
}

func (c *CourseService) Update(ctx context.Context, courseID domain.ID,
	param port.UpdateCourseParam) (domain.Course, error) {
	course, err := c.repo.FindByID(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course by id", zap.Error(err),
			zap.String("courseID", courseID.String()))
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

	course, err = c.repo.Update(ctx, course)
	if err != nil {
		c.logger.Error("failed to update course", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return domain.Course{}, err
	}

	c.logger.Info("course is successfully updated",
		zap.String("courseID", course.ID.String()), zap.String("schoolID", course.SchoolID.String()))
	return course, nil
}

func (c *CourseService) Delete(ctx context.Context, courseID domain.ID) error {
	err := c.repo.Delete(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to delete course", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return err
	}

	c.logger.Info("course is successfully deleted",
		zap.String("courseID", courseID.String()))
	return nil
}
