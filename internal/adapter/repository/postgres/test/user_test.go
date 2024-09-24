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
	db         *sql.DB
	mock       sqlmock.Sqlmock
	repository port.IUserRepository
}

func (s *UserSuite) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	db := sqlx.NewDb(s.db, "pgx")
	s.repository = repository.NewUserRepo(db)
}

// FindAll Suite
type UserFindAllSuite struct {
	UserSuite
}

func (s *UserFindAllSuite) UserFindAllSuccessRepositoryMock(user *domain.User) {
	expectedRows := sqlmock.NewRows([]string{"name", "surname", "email", "password"})
	expectedRows.AddRow(user.Name, user.Surname, user.Email, user.Password)
	s.mock.ExpectQuery(repository.UserFindAllQuery).WillReturnRows(expectedRows)
}

func (s *UserFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Find all success")
	user := NewUserBuilder().Build()
	s.UserFindAllSuccessRepositoryMock(&user)
	users, err := s.repository.FindAll(context.Background())
	t.Assert().Nil(err)
	t.Assert().Equal(users[0].Email, user.Email)
	t.Assert().Equal(users[0].Password, user.Password)
	t.Assert().Equal(users[0].Name, user.Name)
	t.Assert().Equal(users[0].Surname, user.Surname)
}

func (s *UserFindAllSuite) UserFindAllFailureRepositoryMock() {
	s.mock.ExpectQuery(repository.UserFindAllQuery).WillReturnError(sql.ErrNoRows)
}

func (s *UserFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find all failure")
	s.UserFindAllFailureRepositoryMock()
	_, err := s.repository.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindAll", new(UserFindAllSuite))
}
