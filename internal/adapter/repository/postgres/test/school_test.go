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

type SchoolSuite struct {
	suite.Suite
}

func NewSchoolRepository() (port.ISchoolRepository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	conn := sqlx.NewDb(db, "pgx")
	repo := repository.NewSchoolRepo(conn)
	return repo, mock
}

// FindAll Suite
type SchoolFindAllSuite struct {
	SchoolSuite
}

func (s *SchoolFindAllSuite) SchoolFindAllSuccessRepositoryMock(mock sqlmock.Sqlmock, school domain.School) {
	pgSchool := entity.NewPgSchool(school)
	expectedRows := sqlmock.NewRows(EntityColumns(pgSchool))
	expectedRows.AddRow(EntityValues(pgSchool)...)
	mock.ExpectQuery(repository.SchoolFindAllQuery).WillReturnRows(expectedRows)
}

func (s *SchoolFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolFindAllSuccessRepositoryMock(mock, school)
	schools, err := repo.FindAll(context.Background())
	t.Assert().Nil(err)
	t.Assert().Equal(schools[0].Name, school.Name)
}

func (s *SchoolFindAllSuite) SchoolFindAllFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.SchoolFindAllQuery).WillReturnError(sql.ErrNoRows)
}

func (s *SchoolFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewSchoolRepository()
	s.SchoolFindAllFailureRepositoryMock(mock)
	_, err := repo.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindAll", new(SchoolFindAllSuite))
}

type SchoolFindByIDSuite struct {
	SchoolSuite
}

func (s *SchoolFindByIDSuite) SchoolFindByIDSuccessRepositoryMock(mock sqlmock.Sqlmock, school domain.School) {
	pgSchool := entity.NewPgSchool(school)
	expectedRows := sqlmock.NewRows(EntityColumns(pgSchool)).
		AddRow(EntityValues(pgSchool)...)
	mock.ExpectQuery(repository.SchoolFindByIDQuery).WithArgs(school.ID).WillReturnRows(expectedRows)
}

func (s *SchoolFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find by ID success")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolFindByIDSuccessRepositoryMock(mock, school)
	foundSchool, err := repo.FindByID(context.Background(), school.ID)
	t.Assert().Nil(err)
	t.Assert().Equal(foundSchool.ID, school.ID)
}

func (s *SchoolFindByIDSuite) SchoolFindByIDFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.SchoolFindByIDQuery).WillReturnError(sql.ErrNoRows)
}

func (s *SchoolFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find by ID failure")
	repo, mock := NewSchoolRepository()
	s.SchoolFindByIDFailureRepositoryMock(mock)
	_, err := repo.FindByID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindByID", new(SchoolFindByIDSuite))
}

type SchoolFindUserSchoolsSuite struct {
	SchoolSuite
}

func (s *SchoolFindUserSchoolsSuite) SchoolFindUserSchoolsSuccessRepositoryMock(mock sqlmock.Sqlmock,
	school domain.School, userID domain.ID) {
	pgSchool := entity.NewPgSchool(school)
	expectedRows := sqlmock.NewRows(EntityColumns(pgSchool))
	expectedRows.AddRow(EntityValues(pgSchool)...)
	mock.ExpectQuery(repository.SchoolFindUserSchoolsQuery).WithArgs(userID).WillReturnRows(expectedRows)
}

func (s *SchoolFindUserSchoolsSuite) TestFindUserSchools_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	userID := domain.NewID()
	s.SchoolFindUserSchoolsSuccessRepositoryMock(mock, school, userID)
	schools, err := repo.FindUserSchools(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(schools[0].Name, school.Name)
}

func (s *SchoolFindUserSchoolsSuite) SchoolFindUserSchoolsFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.SchoolFindUserSchoolsQuery).WillReturnError(sql.ErrNoRows)
}

func (s *SchoolFindUserSchoolsSuite) TestFindUserSchools_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewSchoolRepository()
	userID := domain.NewID()
	s.SchoolFindUserSchoolsFailureRepositoryMock(mock)
	_, err := repo.FindUserSchools(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindUserSchoolsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindUserSchools", new(SchoolFindUserSchoolsSuite))
}

type SchoolFindSchoolCoursesSuite struct {
	SchoolSuite
}

func (s *SchoolFindSchoolCoursesSuite) SchoolFindSchoolCoursesSuccessRepositoryMock(mock sqlmock.Sqlmock,
	course domain.Course, schoolID domain.ID) {
	pgCourse := entity.NewPgCourse(course)
	expectedRows := sqlmock.NewRows(EntityColumns(pgCourse))
	expectedRows.AddRow(EntityValues(pgCourse)...)
	mock.ExpectQuery(repository.SchoolFindSchoolCoursesQuery).WithArgs(schoolID).WillReturnRows(expectedRows)
}

func (s *SchoolFindSchoolCoursesSuite) TestFindSchoolCourses_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewSchoolRepository()
	course := NewCourseBuilder().Build()
	schoolID := domain.NewID()
	s.SchoolFindSchoolCoursesSuccessRepositoryMock(mock, course, schoolID)
	courses, err := repo.FindSchoolCourses(context.Background(), schoolID)
	t.Assert().Nil(err)
	t.Assert().Equal(courses[0].Name, course.Name)
}

func (s *SchoolFindSchoolCoursesSuite) SchoolFindSchoolCoursesFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.SchoolFindSchoolCoursesQuery).WillReturnError(sql.ErrNoRows)
}

func (s *SchoolFindSchoolCoursesSuite) TestFindSchoolCourses_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewSchoolRepository()
	schoolID := domain.NewID()
	s.SchoolFindSchoolCoursesFailureRepositoryMock(mock)
	_, err := repo.FindSchoolCourses(context.Background(), schoolID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindSchoolCoursesSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindSchoolCourses", new(SchoolFindSchoolCoursesSuite))
}

type SchoolFindSchoolTeachersSuite struct {
	SchoolSuite
}

func (s *SchoolFindSchoolTeachersSuite) SchoolFindSchoolTeachersSuccessRepositoryMock(mock sqlmock.Sqlmock,
	teacher domain.User, schoolID domain.ID) {
	pgUser := entity.NewPgUser(teacher)
	expectedRows := sqlmock.NewRows(EntityColumns(pgUser))
	expectedRows.AddRow(EntityValues(pgUser)...)
	mock.ExpectQuery(repository.SchoolFindSchoolTeachersQuery).WithArgs(schoolID).WillReturnRows(expectedRows)
}

func (s *SchoolFindSchoolTeachersSuite) TestFindSchoolTeachers_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewSchoolRepository()
	teacher := NewUserBuilder().Build()
	schoolID := domain.NewID()
	s.SchoolFindSchoolTeachersSuccessRepositoryMock(mock, teacher, schoolID)
	teachers, err := repo.FindSchoolTeachers(context.Background(), schoolID)
	t.Assert().Nil(err)
	t.Assert().Equal(teachers[0].Name, teacher.Name)
}

func (s *SchoolFindSchoolTeachersSuite) SchoolFindSchoolTeachersFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.SchoolFindSchoolTeachersQuery).WillReturnError(sql.ErrNoRows)
}

func (s *SchoolFindSchoolTeachersSuite) TestFindSchoolTeachers_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewSchoolRepository()
	schoolID := domain.NewID()
	s.SchoolFindSchoolTeachersFailureRepositoryMock(mock)
	_, err := repo.FindSchoolTeachers(context.Background(), schoolID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindSchoolTeachersSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindSchoolTeachers", new(SchoolFindSchoolTeachersSuite))
}

type SchoolIsSchoolTeacherSuite struct {
	SchoolSuite
}

func (s *SchoolIsSchoolTeacherSuite) SchoolIsSchoolTeacherSuccessRepositoryMock(mock sqlmock.Sqlmock,
	schoolID, teacherID domain.ID) {
	expectedRows := sqlmock.NewRows([]string{"?column"})
	expectedRows.AddRow(1)
	mock.ExpectQuery(repository.SchoolContainsTeacherQuery).
		WithArgs(schoolID, teacherID).WillReturnRows(expectedRows)
}

func (s *SchoolIsSchoolTeacherSuite) TestIsSchoolTeacher_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewSchoolRepository()
	schoolID := domain.NewID()
	teacherID := domain.NewID()
	s.SchoolIsSchoolTeacherSuccessRepositoryMock(mock, schoolID, teacherID)
	isTeacher, err := repo.IsSchoolTeacher(context.Background(), schoolID, teacherID)
	t.Assert().Nil(err)
	t.Assert().True(isTeacher)
}

func (s *SchoolIsSchoolTeacherSuite) SchoolIsSchoolTeacherFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.SchoolContainsTeacherQuery).WillReturnError(sql.ErrNoRows)
}

func (s *SchoolIsSchoolTeacherSuite) TestIsSchoolTeacher_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewSchoolRepository()
	schoolID := domain.NewID()
	teacherID := domain.NewID()
	s.SchoolIsSchoolTeacherFailureRepositoryMock(mock)
	_, err := repo.IsSchoolTeacher(context.Background(), schoolID, teacherID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolIsSchoolTeacherSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository IsSchoolTeacher", new(SchoolIsSchoolTeacherSuite))
}

type SchoolAddSchoolTeacherSuite struct {
	SchoolSuite
}

func (s *SchoolAddSchoolTeacherSuite) SchoolAddSchoolTeacherSuccessRepositoryMock(mock sqlmock.Sqlmock,
	schoolID, teacherID domain.ID) {
	mock.ExpectExec(repository.SchoolAddTeacherQuery).
		WithArgs(teacherID, schoolID).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *SchoolAddSchoolTeacherSuite) TestAddSchoolTeacher_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository create school success")
	repo, mock := NewSchoolRepository()
	schoolID := domain.NewID()
	teacherID := domain.NewID()
	s.SchoolAddSchoolTeacherSuccessRepositoryMock(mock, schoolID, teacherID)
	err := repo.AddSchoolTeacher(context.Background(), schoolID, teacherID)
	t.Assert().Nil(err)
}

func (s *SchoolAddSchoolTeacherSuite) SchoolAddSchoolTeacherFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.SchoolAddTeacherQuery).WillReturnError(sql.ErrConnDone)
}

func (s *SchoolAddSchoolTeacherSuite) TestAddSchoolTeacher_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository create school failure")
	repo, mock := NewSchoolRepository()
	schoolID := domain.NewID()
	teacherID := domain.NewID()
	s.SchoolAddSchoolTeacherFailureRepositoryMock(mock)
	err := repo.AddSchoolTeacher(context.Background(), schoolID, teacherID)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestSchoolAddSchoolTeacherSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository AddSchoolTeacher", new(SchoolAddSchoolTeacherSuite))
}

type SchoolCreateSuite struct {
	SchoolSuite
}

func (s *SchoolCreateSuite) SchoolCreateSuccessRepositoryMock(mock sqlmock.Sqlmock, school domain.School) {
	pgSchool := entity.NewPgSchool(school)
	queryString := InsertQueryString(pgSchool, "school")
	mock.ExpectExec(queryString).
		WithArgs(EntityValues(pgSchool)...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedRows := sqlmock.NewRows(EntityColumns(pgSchool)).
		AddRow(EntityValues(pgSchool)...)
	mock.ExpectQuery(repository.SchoolFindByIDQuery).WithArgs(pgSchool.ID).WillReturnRows(expectedRows)
}

func (s *SchoolCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository create school success")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolCreateSuccessRepositoryMock(mock, school)
	createdSchool, err := repo.Create(context.Background(), school)
	t.Assert().Nil(err)
	t.Assert().Equal(createdSchool.Name, school.Name)
}

func (s *SchoolCreateSuite) SchoolCreateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	queryString := InsertQueryString(entity.PgSchool{}, "school")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *SchoolCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository create school failure")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolCreateFailureRepositoryMock(mock)
	_, err := repo.Create(context.Background(), school)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestSchoolCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Create", new(SchoolCreateSuite))
}

type SchoolUpdateSuite struct {
	SchoolSuite
}

func (s *SchoolUpdateSuite) SchoolUpdateSuccessRepositoryMock(mock sqlmock.Sqlmock, school domain.School) {
	pgSchool := entity.NewPgSchool(school)
	queryString := UpdateQueryString(pgSchool, "school")
	values := append(EntityValues(pgSchool), pgSchool.ID)
	mock.ExpectExec(queryString).
		WithArgs(values...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedRows := sqlmock.NewRows(EntityColumns(pgSchool)).
		AddRow(EntityValues(pgSchool)...)
	mock.ExpectQuery(repository.SchoolFindByIDQuery).WithArgs(school.ID).WillReturnRows(expectedRows)
}

func (s *SchoolUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository update school success")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolUpdateSuccessRepositoryMock(mock, school)
	updatedSchool, err := repo.Update(context.Background(), school)
	t.Assert().Nil(err)
	t.Assert().Equal(updatedSchool.Name, school.Name)
}

func (s *SchoolUpdateSuite) SchoolUpdateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	queryString := UpdateQueryString(entity.PgSchool{}, "school")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *SchoolUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository update school failure")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolUpdateFailureRepositoryMock(mock)
	_, err := repo.Update(context.Background(), school)
	t.Assert().ErrorIs(err, errs.ErrUpdateFailed)
}

func TestSchoolUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Update", new(SchoolUpdateSuite))
}

type SchoolDeleteSuite struct {
	SchoolSuite
}

func (s *SchoolDeleteSuite) SchoolDeleteSuccessRepositoryMock(mock sqlmock.Sqlmock, schoolID domain.ID) {
	mock.ExpectExec(repository.SchoolDeleteQuery).WithArgs(schoolID).WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *SchoolDeleteSuite) TestDelete_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository delete school success")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolDeleteSuccessRepositoryMock(mock, school.ID)
	err := repo.Delete(context.Background(), school.ID)
	t.Assert().Nil(err)
}

func (s *SchoolDeleteSuite) SchoolDeleteFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.SchoolDeleteQuery).WillReturnError(sql.ErrConnDone)
}

func (s *SchoolDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository delete school failure")
	repo, mock := NewSchoolRepository()
	school := NewSchoolBuilder().Build()
	s.SchoolDeleteFailureRepositoryMock(mock)
	err := repo.Delete(context.Background(), school.ID)
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestSchoolDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Delete", new(SchoolDeleteSuite))
}
