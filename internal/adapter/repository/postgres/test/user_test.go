package test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"testing"
)

type UserSuite struct {
	suite.Suite
}

func NewUserRepository() (port.IUserRepository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	conn := sqlx.NewDb(db, "pgx")
	repo := repository.NewUserRepo(conn)
	return repo, mock
}

func (s *UserSuite) BeforeEach(t provider.T) {
}

// FindAll Suite
type UserFindAllSuite struct {
	UserSuite
}

func (s *UserFindAllSuite) UserFindAllSuccessRepositoryMock(mock sqlmock.Sqlmock, user *domain.User) {
	expectedRows := sqlmock.NewRows([]string{"name", "surname", "email", "password"})
	expectedRows.AddRow(user.Name, user.Surname, user.Email, user.Password)
	mock.ExpectQuery(repository.UserFindAllQuery).WillReturnRows(expectedRows)
}

func (s *UserFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserFindAllSuccessRepositoryMock(mock, &user)
	users, err := repo.FindAll(context.Background())
	t.Assert().Nil(err)
	t.Assert().Equal(users[0].Email, user.Email)
}

func (s *UserFindAllSuite) UserFindAllFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.UserFindAllQuery).WillReturnError(sql.ErrNoRows)
}

func (s *UserFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find all failure")
	repo, mock := NewUserRepository()
	s.UserFindAllFailureRepositoryMock(mock)
	_, err := repo.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindAll", new(UserFindAllSuite))
}
