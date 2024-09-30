package test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/adapter/repository/postgres/entity"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"testing"
)

type CourseSuite struct {
	suite.Suite
}

func NewCourseRepository() (port.ICourseRepository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	conn := sqlx.NewDb(db, "pgx")
	repo := repository.NewCourseRepo(conn)
	return repo, mock
}

// FindAll Suite
type CourseFindAllSuite struct {
	CourseSuite
}

func (s *CourseFindAllSuite) CourseFindAllSuccessRepositoryMock(mock sqlmock.Sqlmock, course domain.Course) {
	pgCourse := entity.NewPgCourse(course)
	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse))
	expectedRows.AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.CourseFindAllQuery).WillReturnRows(expectedRows)
}

func (s *CourseFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository find all success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseFindAllSuccessRepositoryMock(mock, course)
	courses, err := repo.FindAll(context.Background())
	t.Assert().Nil(err)
	t.Assert().Equal(courses[0].Name, course.Name)
}

func (s *CourseFindAllSuite) CourseFindAllFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseFindAllQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository find all failure")
	repo, mock := NewCourseRepository()
	s.CourseFindAllFailureRepositoryMock(mock)
	_, err := repo.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository find all", new(CourseFindAllSuite))
}

type CourseFindByIDSuite struct {
	CourseSuite
}

func (s *CourseFindByIDSuite) CourseFindByIDSuccessRepositoryMock(mock sqlmock.Sqlmock, course domain.Course) {
	pgCourse := entity.NewPgCourse(course)
	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse)).
		AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.CourseFindByIDQuery).WithArgs(course.ID).WillReturnRows(expectedRows)
}

func (s *CourseFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository find by id success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseFindByIDSuccessRepositoryMock(mock, course)
	foundCourse, err := repo.FindByID(context.Background(), course.ID)
	t.Assert().Nil(err)
	t.Assert().Equal(foundCourse.ID, course.ID)
}

func (s *CourseFindByIDSuite) CourseFindByIDFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseFindByIDQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository find by id failure")
	repo, mock := NewCourseRepository()
	s.CourseFindByIDFailureRepositoryMock(mock)
	_, err := repo.FindByID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository find by id", new(CourseFindByIDSuite))
}

type CourseFindTeacherCoursesSuite struct {
	CourseSuite
}

func (s *CourseFindTeacherCoursesSuite) CourseFindTeacherCoursesSuccessRepositoryMock(mock sqlmock.Sqlmock,
	course domain.Course, teacherID domain.ID) {
	pgCourse := entity.NewPgCourse(course)
	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse))
	expectedRows.AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.CourseFindTeacherCoursesQuery).WithArgs(teacherID).WillReturnRows(expectedRows)
}

func (s *CourseFindTeacherCoursesSuite) TestFindTeacherCourses_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository find teacher courses success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	teacherID := domain.NewID()
	s.CourseFindTeacherCoursesSuccessRepositoryMock(mock, course, teacherID)
	courses, err := repo.FindTeacherCourses(context.Background(), teacherID)
	t.Assert().Nil(err)
	t.Assert().Equal(courses[0].Name, course.Name)
}

func (s *CourseFindTeacherCoursesSuite) CourseFindTeacherCoursesFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseFindTeacherCoursesQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseFindTeacherCoursesSuite) TestFindTeacherCourses_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository find teacher courses failure")
	repo, mock := NewCourseRepository()
	teacherID := domain.NewID()
	s.CourseFindTeacherCoursesFailureRepositoryMock(mock)
	_, err := repo.FindTeacherCourses(context.Background(), teacherID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindTeacherCoursesSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository find teacher courses", new(CourseFindTeacherCoursesSuite))
}

type CourseFindStudentCoursesSuite struct {
	CourseSuite
}

func (s *CourseFindStudentCoursesSuite) CourseFindStudentCoursesSuccessRepositoryMock(mock sqlmock.Sqlmock,
	course domain.Course, studentID domain.ID) {
	pgCourse := entity.NewPgCourse(course)
	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse))
	expectedRows.AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.CourseFindStudentCoursesQuery).WithArgs(studentID).WillReturnRows(expectedRows)
}

func (s *CourseFindStudentCoursesSuite) TestFindStudentCourses_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository find student courses success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	studentID := domain.NewID()
	s.CourseFindStudentCoursesSuccessRepositoryMock(mock, course, studentID)
	courses, err := repo.FindStudentCourses(context.Background(), studentID)
	t.Assert().Nil(err)
	t.Assert().Equal(courses[0].Name, course.Name)
}

func (s *CourseFindStudentCoursesSuite) CourseFindStudentCoursesFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseFindStudentCoursesQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseFindStudentCoursesSuite) TestFindStudentCourses_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository find student courses failure")
	repo, mock := NewCourseRepository()
	studentID := domain.NewID()
	s.CourseFindStudentCoursesFailureRepositoryMock(mock)
	_, err := repo.FindStudentCourses(context.Background(), studentID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindStudentCoursesSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository find student courses", new(CourseFindStudentCoursesSuite))
}

type CourseFindCourseTeachersSuite struct {
	CourseSuite
}

func (s *CourseFindCourseTeachersSuite) CourseFindCourseTeachersSuccessRepositoryMock(mock sqlmock.Sqlmock,
	teacher domain.User, courseID domain.ID) {
	pgUser := entity.NewPgUser(teacher)
	expectedRows := sqlmock.NewRows(EntityColumns(pgUser))
	expectedRows.AddRow(EntityValues(pgUser)...)
	mock.ExpectQuery(repository.CourseFindCourseTeachersQuery).WithArgs(courseID).WillReturnRows(expectedRows)
}

func (s *CourseFindCourseTeachersSuite) TestFindCourseTeachers_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository find course teachers success")
	repo, mock := NewCourseRepository()
	teacher := NewUserBuilder().Build()
	courseID := domain.NewID()
	s.CourseFindCourseTeachersSuccessRepositoryMock(mock, teacher, courseID)
	teachers, err := repo.FindCourseTeachers(context.Background(), courseID)
	t.Assert().Nil(err)
	t.Assert().Equal(teachers[0].Name, teacher.Name)
}

func (s *CourseFindCourseTeachersSuite) CourseFindCourseTeachersFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseFindCourseTeachersQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseFindCourseTeachersSuite) TestFindCourseTeachers_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository find course teachers failure")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	s.CourseFindCourseTeachersFailureRepositoryMock(mock)
	_, err := repo.FindCourseTeachers(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseFindCourseTeachersSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository find course teachers", new(CourseFindCourseTeachersSuite))
}

type CourseIsCourseTeacherSuite struct {
	CourseSuite
}

func (s *CourseIsCourseTeacherSuite) CourseIsCourseTeacherSuccessRepositoryMock(mock sqlmock.Sqlmock,
	courseID, teacherID domain.ID) {
	expectedRows := sqlmock.NewRows([]string{"?column"})
	expectedRows.AddRow(1)
	mock.ExpectQuery(repository.CourseContainsTeacherQuery).
		WithArgs(courseID, teacherID).WillReturnRows(expectedRows)
}

func (s *CourseIsCourseTeacherSuite) TestIsCourseTeacher_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository is course teacher success")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	teacherID := domain.NewID()
	s.CourseIsCourseTeacherSuccessRepositoryMock(mock, courseID, teacherID)
	isTeacher, err := repo.IsCourseTeacher(context.Background(), teacherID, courseID)
	t.Assert().Nil(err)
	t.Assert().True(isTeacher)
}

func (s *CourseIsCourseTeacherSuite) CourseIsCourseTeacherFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseContainsTeacherQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseIsCourseTeacherSuite) TestIsCourseTeacher_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository is course teacher failure")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	teacherID := domain.NewID()
	s.CourseIsCourseTeacherFailureRepositoryMock(mock)
	_, err := repo.IsCourseTeacher(context.Background(), courseID, teacherID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseIsCourseTeacherSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository is course teacher", new(CourseIsCourseTeacherSuite))
}

type CourseIsCourseStudentSuite struct {
	CourseSuite
}

func (s *CourseIsCourseStudentSuite) CourseIsCourseStudentSuccessRepositoryMock(mock sqlmock.Sqlmock,
	courseID, studentID domain.ID) {
	expectedRows := sqlmock.NewRows([]string{"?column"})
	expectedRows.AddRow(1)
	mock.ExpectQuery(repository.CourseContainsStudentQuery).
		WithArgs(courseID, studentID).WillReturnRows(expectedRows)
}

func (s *CourseIsCourseStudentSuite) TestIsCourseStudent_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository is course student success")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	studentID := domain.NewID()
	s.CourseIsCourseStudentSuccessRepositoryMock(mock, courseID, studentID)
	isStudent, err := repo.IsCourseStudent(context.Background(), studentID, courseID)
	t.Assert().Nil(err)
	t.Assert().True(isStudent)
}

func (s *CourseIsCourseStudentSuite) CourseIsCourseStudentFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseContainsStudentQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseIsCourseStudentSuite) TestIsCourseStudent_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository is course student failure")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	studentID := domain.NewID()
	s.CourseIsCourseStudentFailureRepositoryMock(mock)
	_, err := repo.IsCourseStudent(context.Background(), courseID, studentID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseIsCourseStudentSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository is course student", new(CourseIsCourseStudentSuite))
}

type CourseAddCourseTeacherSuite struct {
	CourseSuite
}

func (s *CourseAddCourseTeacherSuite) CourseAddCourseTeacherSuccessRepositoryMock(mock sqlmock.Sqlmock,
	courseID, teacherID domain.ID) {
	mock.ExpectExec(repository.CourseAddCourseTeacherQuery).
		WithArgs(teacherID, courseID).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *CourseAddCourseTeacherSuite) TestAddCourseTeacher_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository add course teacher success")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	teacherID := domain.NewID()
	s.CourseAddCourseTeacherSuccessRepositoryMock(mock, courseID, teacherID)
	err := repo.AddCourseTeacher(context.Background(), teacherID, courseID)
	t.Assert().Nil(err)
}

func (s *CourseAddCourseTeacherSuite) CourseAddCourseTeacherFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.CourseAddCourseTeacherQuery).WillReturnError(sql.ErrConnDone)
}

func (s *CourseAddCourseTeacherSuite) TestAddCourseTeacher_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository add course teacher failure")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	teacherID := domain.NewID()
	s.CourseAddCourseTeacherFailureRepositoryMock(mock)
	err := repo.AddCourseTeacher(context.Background(), courseID, teacherID)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestCourseAddCourseTeacherSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository add course teacher", new(CourseAddCourseTeacherSuite))
}

type CourseAddCourseStudentSuite struct {
	CourseSuite
}

func (s *CourseAddCourseStudentSuite) CourseAddCourseStudentSuccessRepositoryMock(mock sqlmock.Sqlmock,
	courseID, studentID domain.ID) {
	mock.ExpectExec(repository.CourseAddCourseStudentQuery).
		WithArgs(studentID, courseID).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *CourseAddCourseStudentSuite) TestAddCourseStudent_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository add course student success")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	studentID := domain.NewID()
	s.CourseAddCourseStudentSuccessRepositoryMock(mock, courseID, studentID)
	err := repo.AddCourseStudent(context.Background(), studentID, courseID)
	t.Assert().Nil(err)
}

func (s *CourseAddCourseStudentSuite) CourseAddCourseStudentFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.CourseAddCourseStudentQuery).WillReturnError(sql.ErrConnDone)
}

func (s *CourseAddCourseStudentSuite) TestAddCourseStudent_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository add course student failure")
	repo, mock := NewCourseRepository()
	courseID := domain.NewID()
	studentID := domain.NewID()
	s.CourseAddCourseStudentFailureRepositoryMock(mock)
	err := repo.AddCourseStudent(context.Background(), courseID, studentID)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestCourseAddCourseStudentSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository add course student", new(CourseAddCourseStudentSuite))
}

type CourseCreateSuite struct {
	CourseSuite
}

func (s *CourseCreateSuite) CourseCreateSuccessRepositoryMock(mock sqlmock.Sqlmock, course domain.Course) {
	pgCourse := entity.NewPgCourse(course)
	queryString := InsertQueryString(pgCourse, "course")
	mock.ExpectExec(queryString).
		WithArgs(EntityValues(pgCourse)...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse)).
		AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.CourseFindByIDQuery).WithArgs(pgCourse.ID).WillReturnRows(expectedRows)
}

func (s *CourseCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository create course success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseCreateSuccessRepositoryMock(mock, course)
	createdCourse, err := repo.Create(context.Background(), course)
	t.Assert().Nil(err)
	t.Assert().Equal(createdCourse.Name, course.Name)
}

func (s *CourseCreateSuite) CourseCreateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	queryString := InsertQueryString(entity.PgCourse{}, "course")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *CourseCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository create course failure")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseCreateFailureRepositoryMock(mock)
	_, err := repo.Create(context.Background(), course)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestCourseCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository create course", new(CourseCreateSuite))
}

type CourseUpdateSuite struct {
	CourseSuite
}

func (s *CourseUpdateSuite) CourseUpdateSuccessRepositoryMock(mock sqlmock.Sqlmock, course domain.Course) {
	pgCourse := entity.NewPgCourse(course)
	queryString := UpdateQueryString(pgCourse, "course")
	values := append(EntityValues(pgCourse), pgCourse.ID)
	mock.ExpectExec(queryString).
		WithArgs(values...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse)).
		AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.CourseFindByIDQuery).WithArgs(course.ID).WillReturnRows(expectedRows)
}

func (s *CourseUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository update course success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseUpdateSuccessRepositoryMock(mock, course)
	updatedCourse, err := repo.Update(context.Background(), course)
	t.Assert().Nil(err)
	t.Assert().Equal(updatedCourse.Name, course.Name)
}

func (s *CourseUpdateSuite) CourseUpdateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	queryString := UpdateQueryString(entity.PgCourse{}, "course")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *CourseUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository update course failure")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseUpdateFailureRepositoryMock(mock)
	_, err := repo.Update(context.Background(), course)
	t.Assert().ErrorIs(err, errs.ErrUpdateFailed)
}

func TestCourseUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository update course", new(CourseUpdateSuite))
}

type CourseUpdateStatusSuite struct {
	CourseSuite
}

func (s *CourseUpdateStatusSuite) CourseUpdateStatusSuccessRepositoryMock(mock sqlmock.Sqlmock, course domain.Course) {
	pgCourse := entity.NewPgCourse(course)

	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse)).
		AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.CourseFindByIDQuery).WithArgs(pgCourse.ID).WillReturnRows(expectedRows)

	queryString := UpdateQueryString(pgCourse, "course")
	values := append(EntityValues(pgCourse), pgCourse.ID)
	mock.ExpectExec(queryString).
		WithArgs(values...).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *CourseUpdateStatusSuite) TestUpdateStatus_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository update course status success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseUpdateStatusSuccessRepositoryMock(mock, course)
	err := repo.UpdateStatus(context.Background(), course.ID, domain.CourseDraft)
	t.Assert().Nil(err)
}

func (s *CourseUpdateStatusSuite) CourseUpdateStatusFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.CourseFindByIDQuery).WillReturnError(sql.ErrNoRows)
}

func (s *CourseUpdateStatusSuite) TestUpdateStatus_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository update course status failure")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseUpdateStatusFailureRepositoryMock(mock)
	err := repo.UpdateStatus(context.Background(), course.ID, domain.CourseReady)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestCourseUpdateStatusSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository update course status", new(CourseUpdateStatusSuite))
}

type CourseDeleteSuite struct {
	CourseSuite
}

func (s *CourseDeleteSuite) CourseDeleteSuccessRepositoryMock(mock sqlmock.Sqlmock, courseID domain.ID) {
	mock.ExpectExec(repository.CourseDeleteQuery).WithArgs(courseID).WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *CourseDeleteSuite) TestDelete_Success(t provider.T) {
	t.Parallel()
	t.Title("Course repository delete course success")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseDeleteSuccessRepositoryMock(mock, course.ID)
	err := repo.Delete(context.Background(), course.ID)
	t.Assert().Nil(err)
}

func (s *CourseDeleteSuite) CourseDeleteFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.CourseDeleteQuery).WillReturnError(sql.ErrConnDone)
}

func (s *CourseDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Parallel()
	t.Title("Course repository delete course failure")
	repo, mock := NewCourseRepository()
	course := NewCourseBuilder().Build()
	s.CourseDeleteFailureRepositoryMock(mock)
	err := repo.Delete(context.Background(), course.ID)
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestCourseDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course repository delete course", new(CourseDeleteSuite))
}
