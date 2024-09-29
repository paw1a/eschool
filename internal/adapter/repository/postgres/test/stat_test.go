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
	t.Title("Repository find all success")
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
	t.Title("Repository find all failure")
	repo, mock := NewStatRepository()
	userID := domain.NewID()
	lessonID := domain.NewID()
	s.StatFindLessonStatFailureRepositoryMock(mock)
	_, err := repo.FindLessonStat(context.Background(), userID, lessonID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestStatFindLessonStatSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindLessonStat", new(StatFindLessonStatSuite))
}
