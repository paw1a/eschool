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

type LessonSuite struct {
	suite.Suite
}

func NewLessonRepository() (port.ILessonRepository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	conn := sqlx.NewDb(db, "pgx")
	repo := repository.NewLessonRepo(conn)
	return repo, mock
}

// FindAll Suite
type LessonFindAllSuite struct {
	LessonSuite
}

func (s *LessonFindAllSuite) LessonFindAllSuccessRepositoryMock(mock sqlmock.Sqlmock, lesson domain.Lesson) {
	pgLesson := entity.NewPgLesson(lesson)
	expectedRows := sqlmock.NewRows(EntityColumns(pgLesson))
	expectedRows.AddRow(EntityValues(pgLesson)...)
	mock.ExpectQuery(repository.LessonFindAllQuery).WillReturnRows(expectedRows)
}

func (s *LessonFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find all success")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonFindAllSuccessRepositoryMock(mock, lesson)
	lessons, err := repo.FindAll(context.Background())
	t.Assert().Nil(err)
	t.Assert().Equal(lessons[0].Title, lesson.Title)
}

func (s *LessonFindAllSuite) LessonFindAllFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.LessonFindAllQuery).WillReturnError(sql.ErrNoRows)
}

func (s *LessonFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find all failure")
	repo, mock := NewLessonRepository()
	s.LessonFindAllFailureRepositoryMock(mock)
	_, err := repo.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson repository find all", new(LessonFindAllSuite))
}

type LessonFindByIDSuite struct {
	LessonSuite
}

func (s *LessonFindByIDSuite) LessonFindByIDSuccessRepositoryMock(mock sqlmock.Sqlmock, lesson domain.Lesson) {
	pgLesson := entity.NewPgLesson(lesson)
	expectedRows := sqlmock.NewRows(EntityColumns(pgLesson)).
		AddRow(EntityValues(pgLesson)...)
	mock.ExpectQuery(repository.LessonFindByIDQuery).WithArgs(pgLesson.ID).WillReturnRows(expectedRows)
}

func (s *LessonFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find by id success")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonFindByIDSuccessRepositoryMock(mock, lesson)
	foundLesson, err := repo.FindByID(context.Background(), lesson.ID)
	t.Assert().Nil(err)
	t.Assert().Equal(foundLesson.ID, lesson.ID)
}

func (s *LessonFindByIDSuite) LessonFindByIDFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.LessonFindByIDQuery).WillReturnError(sql.ErrNoRows)
}

func (s *LessonFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find by id failure")
	repo, mock := NewLessonRepository()
	s.LessonFindByIDFailureRepositoryMock(mock)
	_, err := repo.FindByID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson repository find by id", new(LessonFindByIDSuite))
}

type LessonFindLessonTestsSuite struct {
	LessonSuite
}

func (s *LessonFindLessonTestsSuite) LessonFindLessonTestsSuccessRepositoryMock(mock sqlmock.Sqlmock,
	tests []domain.Test, lessonID domain.ID) {
	pgTest1 := entity.NewPgTest(tests[0])
	pgTest2 := entity.NewPgTest(tests[1])
	expectedRows := sqlmock.NewRows(EntityColumns(pgTest1))
	expectedRows.AddRow(EntityValues(pgTest1)...)
	expectedRows.AddRow(EntityValues(pgTest2)...)
	mock.ExpectQuery(repository.LessonFindLessonTestsQuery).WithArgs(lessonID).WillReturnRows(expectedRows)
}

func (s *LessonFindLessonTestsSuite) TestFindLessonTests_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find lesson tests success")
	repo, mock := NewLessonRepository()
	lessonID := domain.NewID()
	tests := []domain.Test{NewTestBuilder().Build(), NewTestBuilder().Build()}
	s.LessonFindLessonTestsSuccessRepositoryMock(mock, tests, lessonID)
	actual, err := repo.FindLessonTests(context.Background(), lessonID)
	t.Assert().Nil(err)
	t.Assert().Equal(tests[0], actual[0])
	t.Assert().Equal(tests[1], actual[1])
}

func (s *LessonFindLessonTestsSuite) LessonFindLessonTestsFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.LessonFindLessonTestsQuery).WillReturnError(sql.ErrNoRows)
}

func (s *LessonFindLessonTestsSuite) TestFindLessonTests_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find lesson tests failure")
	repo, mock := NewLessonRepository()
	userID := domain.NewID()
	s.LessonFindLessonTestsFailureRepositoryMock(mock)
	_, err := repo.FindLessonTests(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindLessonTestsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson repository find lesson tests", new(LessonFindLessonTestsSuite))
}

type LessonFindCourseLessonsSuite struct {
	LessonSuite
}

func (s *LessonFindCourseLessonsSuite) LessonFindCourseLessonsSuccessRepositoryMock(mock sqlmock.Sqlmock,
	lesson domain.Lesson, courseID domain.ID) {
	pgLesson := entity.NewPgLesson(lesson)
	expectedRows := sqlmock.NewRows(EntityColumns(pgLesson))
	expectedRows.AddRow(EntityValues(pgLesson)...)
	mock.ExpectQuery(repository.LessonFindCourseLessonsQuery).WithArgs(courseID).WillReturnRows(expectedRows)
}

func (s *LessonFindCourseLessonsSuite) TestFindCourseLessons_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find course lessons success")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	courseID := domain.NewID()
	s.LessonFindCourseLessonsSuccessRepositoryMock(mock, lesson, courseID)
	lessons, err := repo.FindCourseLessons(context.Background(), courseID)
	t.Assert().Nil(err)
	t.Assert().Equal(lessons[0].Title, lesson.Title)
}

func (s *LessonFindCourseLessonsSuite) LessonFindCourseLessonsFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.LessonFindCourseLessonsQuery).WillReturnError(sql.ErrNoRows)
}

func (s *LessonFindCourseLessonsSuite) TestFindCourseLessons_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository find course lessons failure")
	repo, mock := NewLessonRepository()
	courseID := domain.NewID()
	s.LessonFindCourseLessonsFailureRepositoryMock(mock)
	_, err := repo.FindCourseLessons(context.Background(), courseID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestLessonFindCourseLessonsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson repository find course lessons", new(LessonFindCourseLessonsSuite))
}

type LessonCreateSuite struct {
	LessonSuite
}

func (s *LessonCreateSuite) LessonCreateSuccessRepositoryMock(mock sqlmock.Sqlmock, lesson domain.Lesson) {
	pgLesson := entity.NewPgLesson(lesson)
	mock.ExpectBegin()
	queryString := InsertQueryString(pgLesson, "lesson")
	mock.ExpectExec(queryString).
		WithArgs(EntityValues(pgLesson)...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	expectedRows := sqlmock.NewRows(EntityColumns(pgLesson)).
		AddRow(EntityValues(pgLesson)...)
	mock.ExpectQuery(repository.LessonFindByIDQuery).WithArgs(pgLesson.ID).WillReturnRows(expectedRows)
}

func (s *LessonCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository create lesson success")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonCreateSuccessRepositoryMock(mock, lesson)
	createdLesson, err := repo.Create(context.Background(), lesson)
	t.Assert().Nil(err)
	t.Assert().Equal(createdLesson.Title, lesson.Title)
}

func (s *LessonCreateSuite) LessonCreateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	queryString := InsertQueryString(entity.PgLesson{}, "lesson")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *LessonCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository create lesson failure")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonCreateFailureRepositoryMock(mock)
	_, err := repo.Create(context.Background(), lesson)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestLessonCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson repository create lesson", new(LessonCreateSuite))
}

type LessonUpdateSuite struct {
	LessonSuite
}

func (s *LessonUpdateSuite) LessonUpdateSuccessRepositoryMock(mock sqlmock.Sqlmock, lesson domain.Lesson) {
	pgLesson := entity.NewPgLesson(lesson)
	mock.ExpectBegin()
	queryString := UpdateQueryString(pgLesson, "lesson")
	values := append(EntityValues(pgLesson), pgLesson.ID)
	mock.ExpectExec(queryString).
		WithArgs(values...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	expectedRows := sqlmock.NewRows(EntityColumns(pgLesson)).
		AddRow(EntityValues(pgLesson)...)
	mock.ExpectQuery(repository.LessonFindByIDQuery).WithArgs(lesson.ID).WillReturnRows(expectedRows)
}

func (s *LessonUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository update lesson success")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonUpdateSuccessRepositoryMock(mock, lesson)
	updatedLesson, err := repo.Update(context.Background(), lesson)
	t.Assert().Nil(err)
	t.Assert().Equal(updatedLesson.Title, lesson.Title)
}

func (s *LessonUpdateSuite) LessonUpdateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	queryString := UpdateQueryString(entity.PgLesson{}, "lesson")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *LessonUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository update lesson failure")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonUpdateFailureRepositoryMock(mock)
	_, err := repo.Update(context.Background(), lesson)
	t.Assert().ErrorIs(err, errs.ErrUpdateFailed)
}

func TestLessonUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson repository update lesson", new(LessonUpdateSuite))
}

type LessonDeleteSuite struct {
	LessonSuite
}

func (s *LessonDeleteSuite) LessonDeleteSuccessRepositoryMock(mock sqlmock.Sqlmock, lessonID domain.ID) {
	mock.ExpectExec(repository.LessonDeleteQuery).WithArgs(lessonID).WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *LessonDeleteSuite) TestDelete_Success(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository delete lesson success")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonDeleteSuccessRepositoryMock(mock, lesson.ID)
	err := repo.Delete(context.Background(), lesson.ID)
	t.Assert().Nil(err)
}

func (s *LessonDeleteSuite) LessonDeleteFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.LessonDeleteQuery).WillReturnError(sql.ErrConnDone)
}

func (s *LessonDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Parallel()
	t.Title("Lesson repository delete lesson failure")
	repo, mock := NewLessonRepository()
	lesson := NewLessonBuilder().Build()
	s.LessonDeleteFailureRepositoryMock(mock)
	err := repo.Delete(context.Background(), lesson.ID)
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestLessonDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson repository delete lesson", new(LessonDeleteSuite))
}
