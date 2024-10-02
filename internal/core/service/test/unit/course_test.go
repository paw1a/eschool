package unit

import (
	"context"
	"github.com/guregu/null"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

type CourseSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *CourseSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// FindAll Suite
type CourseFindAllSuite struct {
	CourseSuite
}

func CourseFindAllSuccessRepositoryMock(repository *mocks.CourseRepository) {
	repository.
		On("FindAll", context.Background()).
		Return([]domain.Course{NewCourseBuilder().Build()}, nil)
}

func (s *CourseFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service find all success")
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindAllSuccessRepositoryMock(courseRepository)
	_, err := courseService.FindAll(context.Background())
	t.Assert().Nil(err)
}

func CourseFindAllFailureRepositoryMock(repository *mocks.CourseRepository) {
	repository.
		On("FindAll", context.Background()).
		Return(nil, errs.ErrNotExist)
}

func (s *CourseFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service find all failure")
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindAllFailureRepositoryMock(courseRepository)
	_, err := courseService.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service find all", new(CourseFindAllSuite))
}

// FindByID Suite
type CourseFindByIDSuite struct {
	CourseSuite
}

func CourseFindByIDSuccessRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).Build(), nil)
}

func (s *CourseFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service find by id success")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindByIDSuccessRepositoryMock(courseRepository, courseID)
	course, err := courseService.FindByID(context.Background(), courseID)
	t.Assert().Nil(err)
	t.Assert().Equal(courseID, course.ID)
}

func CourseFindByIDFailureRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(domain.Course{}, errs.ErrNotExist)
}

func (s *CourseFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service find by id failure")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindByIDFailureRepositoryMock(courseRepository, courseID)
	_, err := courseService.FindByID(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service find by id", new(CourseFindByIDSuite))
}

// FindStudentCourses Suite
type CourseFindStudentCoursesSuite struct {
	CourseSuite
}

func CourseFindStudentCoursesSuccessRepositoryMock(repository *mocks.CourseRepository, userID domain.ID) {
	repository.
		On("FindStudentCourses", context.Background(), userID).
		Return([]domain.Course{NewCourseBuilder().Build()}, nil)
}

func (s *CourseFindStudentCoursesSuite) TestFindStudentCourses_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service find student courses success")
	userID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindStudentCoursesSuccessRepositoryMock(courseRepository, userID)
	courses, err := courseService.FindStudentCourses(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(courses[0].Name, NewCourseBuilder().Build().Name)
}

func CourseFindStudentCoursesFailureRepositoryMock(repository *mocks.CourseRepository, userID domain.ID) {
	repository.
		On("FindStudentCourses", context.Background(), userID).
		Return(nil, errs.ErrNotExist)
}

func (s *CourseFindStudentCoursesSuite) TestFindStudentCourses_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service find student courses failure")
	userID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindStudentCoursesFailureRepositoryMock(courseRepository, userID)
	_, err := courseService.FindStudentCourses(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindStudentCoursesSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindStudentCourses", new(CourseFindStudentCoursesSuite))
}

// FindTeacherCourses Suite
type CourseFindTeacherCoursesSuite struct {
	CourseSuite
}

func CourseFindTeacherCoursesSuccessRepositoryMock(repository *mocks.CourseRepository, userID domain.ID) {
	repository.
		On("FindTeacherCourses", context.Background(), userID).
		Return([]domain.Course{NewCourseBuilder().Build()}, nil)
}

func (s *CourseFindTeacherCoursesSuite) TestFindTeacherCourses_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service find teacher courses success")
	userID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindTeacherCoursesSuccessRepositoryMock(courseRepository, userID)
	courses, err := courseService.FindTeacherCourses(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(courses[0].Name, NewCourseBuilder().Build().Name)
}

func CourseFindTeacherCoursesFailureRepositoryMock(repository *mocks.CourseRepository, userID domain.ID) {
	repository.
		On("FindTeacherCourses", context.Background(), userID).
		Return(nil, errs.ErrNotExist)
}

func (s *CourseFindTeacherCoursesSuite) TestFindTeacherCourses_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service find teacher courses failure")
	userID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindTeacherCoursesFailureRepositoryMock(courseRepository, userID)
	_, err := courseService.FindTeacherCourses(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindTeacherCoursesSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service find teacher courses", new(CourseFindTeacherCoursesSuite))
}

// FindCourseTeachers Suite
type CourseFindCourseTeachersSuite struct {
	CourseSuite
}

func CourseFindCourseTeachersSuccessRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("FindCourseTeachers", context.Background(), courseID).
		Return([]domain.User{NewUserBuilder().Build()}, nil)
}

func (s *CourseFindCourseTeachersSuite) TestFindCourseTeachers_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service find course teachers success")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindCourseTeachersSuccessRepositoryMock(courseRepository, courseID)
	teachers, err := courseService.FindCourseTeachers(context.Background(), courseID)
	t.Assert().Nil(err)
	t.Assert().Equal(teachers[0].Name, NewUserBuilder().Build().Name)
}

func CourseFindCourseTeachersFailureRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("FindCourseTeachers", context.Background(), courseID).
		Return(nil, errs.ErrNotExist)
}

func (s *CourseFindCourseTeachersSuite) TestFindCourseTeachers_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service find course teachers failure")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseFindCourseTeachersFailureRepositoryMock(courseRepository, courseID)
	_, err := courseService.FindCourseTeachers(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindCourseTeachersSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service find course teachers", new(CourseFindCourseTeachersSuite))
}

// IsCourseStudent Suite
type CourseIsCourseStudentSuite struct {
	CourseSuite
}

func CourseIsCourseStudentSuccessRepositoryMock(repository *mocks.CourseRepository, courseID, studentID domain.ID) {
	repository.
		On("IsCourseStudent", context.Background(), courseID, studentID).
		Return(true, nil)
}

func (s *CourseIsCourseStudentSuite) TestIsCourseStudent_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service is course student success")
	courseID := domain.NewID()
	studentID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseIsCourseStudentSuccessRepositoryMock(courseRepository, courseID, studentID)
	isStudent, err := courseService.IsCourseStudent(context.Background(), courseID, studentID)
	t.Assert().Nil(err)
	t.Assert().True(isStudent)
}

func CourseIsCourseStudentFailureRepositoryMock(repository *mocks.CourseRepository, courseID, studentID domain.ID) {
	repository.
		On("IsCourseStudent", context.Background(), courseID, studentID).
		Return(false, errs.ErrNotExist)
}

func (s *CourseIsCourseStudentSuite) TestIsCourseStudent_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service is course student failure")
	courseID := domain.NewID()
	studentID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseIsCourseStudentFailureRepositoryMock(courseRepository, courseID, studentID)
	isStudent, err := courseService.IsCourseStudent(context.Background(), courseID, studentID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
	t.Assert().False(isStudent)
}

func TestCourseIsCourseStudentSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service is course student", new(CourseIsCourseStudentSuite))
}

// IsCourseTeacher Suite
type CourseIsCourseTeacherSuite struct {
	CourseSuite
}

func CourseIsCourseTeacherSuccessRepositoryMock(repository *mocks.CourseRepository, courseID, teacherID domain.ID) {
	repository.
		On("IsCourseTeacher", context.Background(), courseID, teacherID).
		Return(true, nil)
}

func (s *CourseIsCourseTeacherSuite) TestIsCourseTeacher_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service is course teacher success")
	courseID := domain.NewID()
	teacherID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseIsCourseTeacherSuccessRepositoryMock(courseRepository, courseID, teacherID)
	isTeacher, err := courseService.IsCourseTeacher(context.Background(), courseID, teacherID)
	t.Assert().Nil(err)
	t.Assert().True(isTeacher)
}

func CourseIsCourseTeacherFailureRepositoryMock(repository *mocks.CourseRepository, courseID, teacherID domain.ID) {
	repository.
		On("IsCourseTeacher", context.Background(), courseID, teacherID).
		Return(false, errs.ErrNotExist)
}

func (s *CourseIsCourseTeacherSuite) TestIsCourseTeacher_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service is course teacher failure")
	courseID := domain.NewID()
	teacherID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseIsCourseTeacherFailureRepositoryMock(courseRepository, courseID, teacherID)
	isTeacher, err := courseService.IsCourseTeacher(context.Background(), courseID, teacherID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
	t.Assert().False(isTeacher)
}

func TestCourseIsCourseTeacherSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service is course teacher", new(CourseIsCourseTeacherSuite))
}

// AddCourseStudent Suite
type CourseAddCourseStudentSuite struct {
	CourseSuite
}

func CourseAddCourseStudentSuccessRepositoryMock(repository *mocks.CourseRepository,
	lessonRepository *mocks.LessonRepository, statRepository *mocks.StatRepository,
	courseID, studentID domain.ID) {
	repository.
		On("AddCourseStudent", context.Background(), studentID, courseID).
		Return(nil)
	lessonRepository.
		On("FindCourseLessons", context.Background(), courseID).
		Return([]domain.Lesson{NewLessonBuilder().Build()}, nil)
	statRepository.
		On("CreateLessonStat", context.Background(), mock.Anything).
		Return(nil)
}

func (s *CourseAddCourseStudentSuite) TestAddCourseStudent_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service add course student success")
	courseID := domain.NewID()
	studentID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseAddCourseStudentSuccessRepositoryMock(courseRepository, lessonRepository,
		statRepository, courseID, studentID)
	err := courseService.AddCourseStudent(context.Background(), studentID, courseID)
	t.Assert().Nil(err)
}

func CourseAddCourseStudentFailureRepositoryMock(repository *mocks.CourseRepository,
	lessonRepository *mocks.LessonRepository,
	statRepository *mocks.StatRepository, courseID, studentID domain.ID) {
	repository.
		On("AddCourseStudent", context.Background(), studentID, courseID).
		Return(errs.ErrNotExist)
	lessonRepository.
		On("FindCourseLessons", context.Background(), courseID).
		Return([]domain.Lesson{NewLessonBuilder().Build()}, nil)
	statRepository.
		On("CreateLessonStat", context.Background(), mock.Anything).
		Return(nil)
}

func (s *CourseAddCourseStudentSuite) TestAddCourseStudent_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service add course student failure")
	courseID := domain.NewID()
	studentID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseAddCourseStudentFailureRepositoryMock(courseRepository, lessonRepository,
		statRepository, courseID, studentID)
	err := courseService.AddCourseStudent(context.Background(), studentID, courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseAddCourseStudentSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service add course student", new(CourseAddCourseStudentSuite))
}

// AddCourseTeacher Suite
type CourseAddCourseTeacherSuite struct {
	CourseSuite
}

func CourseAddCourseTeacherSuccessRepositoryMock(repository *mocks.CourseRepository,
	schoolRepository *mocks.SchoolRepository, courseID, teacherID domain.ID) {
	repository.
		On("AddCourseTeacher", context.Background(), teacherID, courseID).
		Return(nil)
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).Build(), nil)
	schoolRepository.
		On("IsSchoolTeacher", context.Background(), mock.Anything, teacherID).
		Return(true, nil)
}

func (s *CourseAddCourseTeacherSuite) TestAddCourseTeacher_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service add course teacher success")
	courseID := domain.NewID()
	teacherID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseAddCourseTeacherSuccessRepositoryMock(courseRepository, schoolRepository, courseID, teacherID)
	err := courseService.AddCourseTeacher(context.Background(), teacherID, courseID)
	t.Assert().Nil(err)
}

func CourseAddCourseTeacherFailureRepositoryMock(repository *mocks.CourseRepository,
	schoolRepository *mocks.SchoolRepository, courseID, teacherID domain.ID) {
	repository.
		On("AddCourseTeacher", context.Background(), teacherID, courseID).
		Return(errs.ErrNotExist)
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).Build(), nil)
	schoolRepository.
		On("IsSchoolTeacher", context.Background(), mock.Anything, teacherID).
		Return(true, nil)
}

func (s *CourseAddCourseTeacherSuite) TestAddCourseTeacher_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service add course teacher failure")
	courseID := domain.NewID()
	teacherID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseAddCourseTeacherFailureRepositoryMock(courseRepository, schoolRepository, courseID, teacherID)
	err := courseService.AddCourseTeacher(context.Background(), teacherID, courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseAddCourseTeacherSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service add course teacher", new(CourseAddCourseTeacherSuite))
}

// Create Suite
type CourseCreateSuite struct {
	CourseSuite
}

func CourseCreateSuccessRepositoryMock(repository *mocks.CourseRepository, name string) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewCourseBuilder().WithName(name).Build(), nil)
}

func (s *CourseCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service create course success")
	schoolID := domain.NewID()
	name := "course name"
	param := NewCreateCourseParamBuilder().WithName(name).Build()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseCreateSuccessRepositoryMock(courseRepository, name)
	course, err := courseService.CreateSchoolCourse(context.Background(), schoolID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Name, course.Name)
}

func CourseCreateFailureRepositoryMock(repository *mocks.CourseRepository) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.Course{}, errors.New("error"))
}

func (s *CourseCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service create course failure")
	schoolID := domain.NewID()
	param := NewCreateCourseParamBuilder().Build()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseCreateFailureRepositoryMock(courseRepository)
	_, err := courseService.CreateSchoolCourse(context.Background(), schoolID, param)
	t.Assert().NotNil(err)
}

func TestCourseCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service create course", new(CourseCreateSuite))
}

// Update Suite
type CourseUpdateSuite struct {
	CourseSuite
}

func CourseUpdateSuccessRepositoryMock(repository *mocks.CourseRepository,
	courseID domain.ID, name string) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).Build(), nil)
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewCourseBuilder().WithName(name).Build(), nil)
}

func (s *CourseUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service update course success")
	courseID := domain.NewID()
	name := "course name"
	param := NewUpdateCourseParamBuilder().WithName(null.StringFrom(name)).Build()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseUpdateSuccessRepositoryMock(courseRepository, courseID, name)
	course, err := courseService.Update(context.Background(), courseID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Name.String, course.Name)
}

func CourseUpdateFailureRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(domain.Course{}, errs.ErrNotExist)
}

func (s *CourseUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service update course failure")
	courseID := domain.NewID()
	param := NewUpdateCourseParamBuilder().Build()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseUpdateFailureRepositoryMock(courseRepository, courseID)
	_, err := courseService.Update(context.Background(), courseID, param)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service update course", new(CourseUpdateSuite))
}

// Delete Suite
type CourseDeleteSuite struct {
	CourseSuite
}

func CourseDeleteSuccessRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("Delete", context.Background(), courseID).
		Return(nil)
}

func (s *CourseDeleteSuite) TestDelete_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service delete course success")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseDeleteSuccessRepositoryMock(courseRepository, courseID)
	err := courseService.Delete(context.Background(), courseID)
	t.Assert().Nil(err)
}

func CourseDeleteFailureRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("Delete", context.Background(), courseID).
		Return(errs.ErrNotExist)
}

func (s *CourseDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service delete course failure")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseDeleteFailureRepositoryMock(courseRepository, courseID)
	err := courseService.Delete(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service delete course", new(CourseDeleteSuite))
}

// PublishReadyCourse Suite
type CoursePublishReadyCourseSuite struct {
	CourseSuite
}

func CoursePublishReadyCourseSuccessRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).WithStatus(domain.CourseReady).Build(), nil)
	repository.
		On("UpdateStatus", context.Background(), courseID, mock.Anything).
		Return(nil)
}

func (s *CoursePublishReadyCourseSuite) TestPublishReadyCourse_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service publish ready course success")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CoursePublishReadyCourseSuccessRepositoryMock(courseRepository, courseID)
	err := courseService.PublishReadyCourse(context.Background(), courseID)
	t.Assert().Nil(err)
}

func CoursePublishReadyCourseFailureRepositoryMock(repository *mocks.CourseRepository, courseID domain.ID) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).WithStatus(domain.CourseReady).Build(), nil)
	repository.
		On("UpdateStatus", context.Background(), courseID, mock.Anything).
		Return(errs.ErrUpdateFailed)
}

func (s *CoursePublishReadyCourseSuite) TestPublishReadyCourse_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service publish ready course failure")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CoursePublishReadyCourseFailureRepositoryMock(courseRepository, courseID)
	err := courseService.PublishReadyCourse(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrUpdateFailed)
}

func TestCoursePublishReadyCourseSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service publish ready course", new(CoursePublishReadyCourseSuite))
}

// ConfirmDraftCourse Suite
type CourseConfirmDraftCourseSuite struct {
	CourseSuite
}

func CourseConfirmDraftCourseSuccessRepositoryMock(repository *mocks.CourseRepository,
	lessonRepository *mocks.LessonRepository, courseID domain.ID) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).WithStatus(domain.CourseDraft).Build(), nil)
	repository.
		On("UpdateStatus", context.Background(), courseID, mock.Anything).
		Return(nil)
	lessonRepository.
		On("FindCourseLessons", context.Background(), courseID).
		Return([]domain.Lesson{
			NewLessonBuilder().
				WithType(domain.TheoryLesson).
				WithTheoryUrl(null.StringFrom("url")).
				Build(),
			NewLessonBuilder().
				WithType(domain.PracticeLesson).
				WithTests([]domain.Test{NewTestBuilder().Build()}).
				Build()}, nil)
}

func (s *CourseConfirmDraftCourseSuite) TestConfirmDraftCourse_Success(t provider.T) {
	t.Parallel()
	t.Title("Course service confirm draft course success")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseConfirmDraftCourseSuccessRepositoryMock(courseRepository, lessonRepository, courseID)
	err := courseService.ConfirmDraftCourse(context.Background(), courseID)
	t.Assert().Nil(err)
}

func CourseConfirmDraftCourseFailureRepositoryMock(repository *mocks.CourseRepository,
	lessonRepository *mocks.LessonRepository, courseID domain.ID) {
	repository.
		On("FindByID", context.Background(), courseID).
		Return(NewCourseBuilder().WithID(courseID).WithStatus(domain.CourseDraft).Build(), nil)
	repository.
		On("UpdateStatus", context.Background(), courseID, mock.Anything).
		Return(errs.ErrUpdateFailed)
	lessonRepository.
		On("FindCourseLessons", context.Background(), courseID).
		Return([]domain.Lesson{
			NewLessonBuilder().
				WithType(domain.TheoryLesson).
				WithTheoryUrl(null.StringFrom("url")).
				Build(),
			NewLessonBuilder().
				WithType(domain.PracticeLesson).
				WithTests([]domain.Test{NewTestBuilder().Build()}).
				Build()}, nil)
}

func (s *CourseConfirmDraftCourseSuite) TestConfirmDraftCourse_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course service confirm draft course failure")
	courseID := domain.NewID()
	courseRepository := mocks.NewCourseRepository(t)
	lessonRepository := mocks.NewLessonRepository(t)
	schoolRepository := mocks.NewSchoolRepository(t)
	statRepository := mocks.NewStatRepository(t)
	courseService := service.NewCourseService(courseRepository, lessonRepository,
		schoolRepository, statRepository, s.logger)
	CourseConfirmDraftCourseFailureRepositoryMock(courseRepository, lessonRepository, courseID)
	errArray := courseService.ConfirmDraftCourse(context.Background(), courseID)
	t.Assert().ErrorIs(errArray[0], errs.ErrUpdateFailed)
}

func TestCourseConfirmDraftCourseSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service confirm draft course", new(CourseConfirmDraftCourseSuite))
}
