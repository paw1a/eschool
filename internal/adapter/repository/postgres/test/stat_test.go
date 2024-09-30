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

type StatSuite struct {
	suite.Suite
}

func NewStatRepository() (port.IStatRepository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	conn := sqlx.NewDb(db, "pgx")
	repo := repository.NewStatRepo(conn)
	return repo, mock
}

type StatFindLessonStatSuite struct {
	StatSuite
}

func (s *StatFindLessonStatSuite) StatFindLessonStatSuccessRepositoryMock(mock sqlmock.Sqlmock,
	stat domain.LessonStat, userID, lessonID domain.ID) {
	pgStat := entity.NewPgLessonStat(stat)
	expectedRows := sqlmock.NewRows(EntityColumns(pgStat))
	expectedRows.AddRow(EntityValues(pgStat)...)
	mock.ExpectQuery(repository.StatFindByUserLessonQuery).WithArgs(userID, lessonID).WillReturnRows(expectedRows)
	mock.ExpectQuery(repository.StatFindLessonTestsQuery).WithArgs(lessonID).WillReturnError(sql.ErrNoRows)
}

func (s *StatFindLessonStatSuite) TestFindLessonStat_Success(t provider.T) {
	t.Parallel()
	t.Title("Stat repository find lesson stat success")
	repo, mock := NewStatRepository()
	stat := NewLessonStatBuilder().Build()
	userID := domain.NewID()
	lessonID := domain.NewID()
	s.StatFindLessonStatSuccessRepositoryMock(mock, stat, userID, lessonID)
	actual, err := repo.FindLessonStat(context.Background(), userID, lessonID)
	t.Assert().Nil(err)
	t.Assert().Equal(stat.Score, actual.Score)
}

func (s *StatFindLessonStatSuite) StatFindLessonStatFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.StatFindByUserLessonQuery).WillReturnError(sql.ErrNoRows)
}

func (s *StatFindLessonStatSuite) TestFindLessonStat_Failure(t provider.T) {
	t.Parallel()
	t.Title("Stat repository find lesson stat failure")
	repo, mock := NewStatRepository()
	userID := domain.NewID()
	lessonID := domain.NewID()
	s.StatFindLessonStatFailureRepositoryMock(mock)
	_, err := repo.FindLessonStat(context.Background(), userID, lessonID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestStatFindLessonStatSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Stat repository find lesson stat", new(StatFindLessonStatSuite))
}

type StatCreateSuite struct {
	StatSuite
}

func (s *StatCreateSuite) StatCreateSuccessRepositoryMock(mock sqlmock.Sqlmock, stat domain.LessonStat) {
	pgStat := entity.NewPgLessonStat(stat)
	mock.ExpectBegin()
	queryString := InsertQueryString(pgStat, "lesson_stat")
	mock.ExpectExec(queryString).
		WithArgs(EntityValues(pgStat)...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
}

func (s *StatCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Stat repository create lesson stat success")
	repo, mock := NewStatRepository()
	stat := NewLessonStatBuilder().Build()
	s.StatCreateSuccessRepositoryMock(mock, stat)
	err := repo.CreateLessonStat(context.Background(), stat)
	t.Assert().Nil(err)
}

func (s *StatCreateSuite) StatCreateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	queryString := InsertQueryString(entity.PgLessonStat{}, "lesson_stat")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *StatCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Stat repository create lesson stat failure")
	repo, mock := NewStatRepository()
	stat := NewLessonStatBuilder().Build()
	s.StatCreateFailureRepositoryMock(mock)
	err := repo.CreateLessonStat(context.Background(), stat)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestStatCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Stat repository create lesson stat", new(StatCreateSuite))
}

type StatUpdateSuite struct {
	StatSuite
}

func (s *StatUpdateSuite) StatUpdateSuccessRepositoryMock(mock sqlmock.Sqlmock, stat domain.LessonStat) {
	pgStat := entity.NewPgLessonStat(stat)
	mock.ExpectBegin()
	queryString := UpdateQueryString(pgStat, "lesson_stat")
	values := append(EntityValues(pgStat), pgStat.ID)
	mock.ExpectExec(queryString).
		WithArgs(values...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
}

func (s *StatUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Stat repository update lesson stat success")
	repo, mock := NewStatRepository()
	stat := NewLessonStatBuilder().Build()
	s.StatUpdateSuccessRepositoryMock(mock, stat)
	err := repo.UpdateLessonStat(context.Background(), stat)
	t.Assert().Nil(err)
}

func (s *StatUpdateSuite) StatUpdateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	queryString := UpdateQueryString(entity.PgLessonStat{}, "lesson_stat")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *StatUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Stat repository update lesson stat failure")
	repo, mock := NewStatRepository()
	stat := NewLessonStatBuilder().Build()
	s.StatUpdateFailureRepositoryMock(mock)
	err := repo.UpdateLessonStat(context.Background(), stat)
	t.Assert().ErrorIs(err, errs.ErrUpdateFailed)
}

func TestStatUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Stat repository update lesson stat", new(StatUpdateSuite))
}
