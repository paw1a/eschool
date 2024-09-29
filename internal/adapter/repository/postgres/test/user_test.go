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

type UserSuite struct {
	suite.Suite
}

func NewUserRepository() (port.IUserRepository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	conn := sqlx.NewDb(db, "pgx")
	repo := repository.NewUserRepo(conn)
	return repo, mock
}

// FindAll Suite
type UserFindAllSuite struct {
	UserSuite
}

func (s *UserFindAllSuite) UserFindAllSuccessRepositoryMock(mock sqlmock.Sqlmock, user domain.User) {
	pgUser := entity.NewPgUser(user)
	expectedRows := sqlmock.NewRows(EntityColumns(pgUser))
	expectedRows.AddRow(EntityValues(pgUser)...)
	mock.ExpectQuery(repository.UserFindAllQuery).WillReturnRows(expectedRows)
}

func (s *UserFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find all success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserFindAllSuccessRepositoryMock(mock, user)
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

type UserFindByIDSuite struct {
	UserSuite
}

func (s *UserFindByIDSuite) UserFindByIDSuccessRepositoryMock(mock sqlmock.Sqlmock, user domain.User) {
	pgUser := entity.NewPgUser(user)
	expectedRows := sqlmock.NewRows(EntityColumns(pgUser)).
		AddRow(EntityValues(pgUser)...)
	mock.ExpectQuery(repository.UserFindByIDQuery).WithArgs(user.ID).WillReturnRows(expectedRows)
}

func (s *UserFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find by ID success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserFindByIDSuccessRepositoryMock(mock, user)
	foundUser, err := repo.FindByID(context.Background(), user.ID)
	t.Assert().Nil(err)
	t.Assert().Equal(foundUser.ID, user.ID)
}

func (s *UserFindByIDSuite) UserFindByIDFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.UserFindByIDQuery).WillReturnError(sql.ErrNoRows)
}

func (s *UserFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find by ID failure")
	repo, mock := NewUserRepository()
	s.UserFindByIDFailureRepositoryMock(mock)
	_, err := repo.FindByID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindByID", new(UserFindByIDSuite))
}

type UserFindByEmailSuite struct {
	UserSuite
}

func (s *UserFindByEmailSuite) UserFindByEmailSuccessRepositoryMock(mock sqlmock.Sqlmock, user domain.User) {
	pgUser := entity.NewPgUser(user)
	expectedRows := sqlmock.NewRows(EntityColumns(pgUser)).
		AddRow(EntityValues(pgUser)...)
	mock.ExpectQuery(repository.UserFindByEmailQuery).WithArgs(user.Email).WillReturnRows(expectedRows)
}

func (s *UserFindByEmailSuite) TestFindByEmail_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find by email success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserFindByEmailSuccessRepositoryMock(mock, user)
	foundUser, err := repo.FindByEmail(context.Background(), user.Email)
	t.Assert().Nil(err)
	t.Assert().Equal(foundUser.Email, user.Email)
}

func (s *UserFindByEmailSuite) UserFindByEmailFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(repository.UserFindByEmailQuery).WillReturnError(sql.ErrNoRows)
}

func (s *UserFindByEmailSuite) TestFindByEmail_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find by email failure")
	repo, mock := NewUserRepository()
	s.UserFindByEmailFailureRepositoryMock(mock)
	_, err := repo.FindByEmail(context.Background(), "nonexistent@example.com")
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindByEmailSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindByEmail", new(UserFindByEmailSuite))
}

type UserFindByCredentialsSuite struct {
	UserSuite
}

func (s *UserFindByCredentialsSuite) UserFindByCredentialsSuccessRepositoryMock(mock sqlmock.Sqlmock, email, password string, user domain.User) {
	pgUser := entity.NewPgUser(user)
	expectedRows := sqlmock.NewRows(EntityColumns(pgUser)).
		AddRow(EntityValues(pgUser)...)
	mock.ExpectQuery(repository.UserFindByCredentialsQuery).WithArgs(email, password).WillReturnRows(expectedRows)
}

func (s *UserFindByCredentialsSuite) TestFindByCredentials_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find by credentials success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserFindByCredentialsSuccessRepositoryMock(mock, user.Email, user.Password, user)
	foundUser, err := repo.FindByCredentials(context.Background(), user.Email, user.Password)
	t.Assert().Nil(err)
	t.Assert().Equal(foundUser.Email, user.Email)
}

func (s *UserFindByCredentialsSuite) UserFindByCredentialsFailureRepositoryMock(mock sqlmock.Sqlmock, email, password string) {
	mock.ExpectQuery(repository.UserFindByCredentialsQuery).WithArgs(email, password).WillReturnError(sql.ErrNoRows)
}

func (s *UserFindByCredentialsSuite) TestFindByCredentials_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find by credentials failure")
	repo, mock := NewUserRepository()
	s.UserFindByCredentialsFailureRepositoryMock(mock, "nonexistent@example.com", "wrongpassword")
	_, err := repo.FindByCredentials(context.Background(), "nonexistent@example.com", "wrongpassword")
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindByCredentialsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindByCredentials", new(UserFindByCredentialsSuite))
}

type UserFindUserInfoSuite struct {
	UserSuite
}

func (s *UserFindUserInfoSuite) UserFindUserInfoSuccessRepositoryMock(mock sqlmock.Sqlmock,
	userInfo port.UserInfo, userID domain.ID) {
	expectedRows := sqlmock.NewRows([]string{"name", "surname"}).
		AddRow(userInfo.Name, userInfo.Surname)
	mock.ExpectQuery(repository.UserFindUserInfoQuery).WithArgs(userID).WillReturnRows(expectedRows)
}

func (s *UserFindUserInfoSuite) TestFindUserInfo_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository find user info success")
	repo, mock := NewUserRepository()
	userID := domain.NewID()
	userInfo := port.UserInfo{Name: "name", Surname: "surname"}
	s.UserFindUserInfoSuccessRepositoryMock(mock, userInfo, userID)
	foundUserInfo, err := repo.FindUserInfo(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(foundUserInfo.Name, userInfo.Name)
	t.Assert().Equal(foundUserInfo.Surname, userInfo.Surname)
}

func (s *UserFindUserInfoSuite) UserFindUserInfoFailureRepositoryMock(mock sqlmock.Sqlmock, userID domain.ID) {
	mock.ExpectQuery(repository.UserFindUserInfoQuery).WithArgs(userID).WillReturnError(sql.ErrNoRows)
}

func (s *UserFindUserInfoSuite) TestFindUserInfo_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository find user info failure")
	userID := domain.NewID()
	repo, mock := NewUserRepository()
	s.UserFindUserInfoFailureRepositoryMock(mock, userID)
	_, err := repo.FindUserInfo(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindUserInfoSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository FindUserInfo", new(UserFindUserInfoSuite))
}

type UserCreateSuite struct {
	UserSuite
}

func (s *UserCreateSuite) UserCreateSuccessRepositoryMock(mock sqlmock.Sqlmock, user domain.User) {
	pgUser := entity.NewPgUser(user)
	queryString := InsertQueryString(pgUser, "user")
	mock.ExpectExec(queryString).
		WithArgs(sqlmock.AnyArg(),
			user.Name,
			user.Surname,
			user.Phone,
			user.City,
			user.AvatarUrl,
			user.Email,
			user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedRows := sqlmock.NewRows(EntityColumns(pgUser)).
		AddRow(user.ID.String(),
			user.Name,
			user.Surname,
			user.Phone,
			user.City,
			user.AvatarUrl,
			user.Email,
			user.Password)
	mock.ExpectQuery(repository.UserFindByIDQuery).WithArgs(user.ID).WillReturnRows(expectedRows)
}

func (s *UserCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository create user success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserCreateSuccessRepositoryMock(mock, user)
	createdUser, err := repo.Create(context.Background(), user)
	t.Assert().Nil(err)
	t.Assert().Equal(createdUser.Email, user.Email)
}

func (s *UserCreateSuite) UserCreateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	queryString := InsertQueryString(entity.PgUser{}, "user")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *UserCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository create user failure")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserCreateFailureRepositoryMock(mock)
	_, err := repo.Create(context.Background(), user)
	t.Assert().ErrorIs(err, errs.ErrPersistenceFailed)
}

func TestUserCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Create", new(UserCreateSuite))
}

type UserUpdateSuite struct {
	UserSuite
}

func (s *UserUpdateSuite) UserUpdateSuccessRepositoryMock(mock sqlmock.Sqlmock, user domain.User) {
	pgUser := entity.NewPgUser(user)
	queryString := UpdateQueryString(pgUser, "user")
	mock.ExpectExec(queryString).
		WithArgs(sqlmock.AnyArg(),
			user.Name,
			user.Surname,
			user.Phone,
			user.City,
			user.AvatarUrl,
			user.Email,
			user.Password,
			pgUser.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expectedRows := sqlmock.NewRows(EntityColumns(pgUser)).
		AddRow(user.ID.String(),
			user.Name,
			user.Surname,
			user.Phone,
			user.City,
			user.AvatarUrl,
			user.Email,
			user.Password)
	mock.ExpectQuery(repository.UserFindByIDQuery).WithArgs(user.ID).WillReturnRows(expectedRows)
}

func (s *UserUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository update user success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserUpdateSuccessRepositoryMock(mock, user)
	updatedUser, err := repo.Update(context.Background(), user)
	t.Assert().Nil(err)
	t.Assert().Equal(updatedUser.Email, user.Email)
}

func (s *UserUpdateSuite) UserUpdateFailureRepositoryMock(mock sqlmock.Sqlmock) {
	queryString := UpdateQueryString(entity.PgUser{}, "user")
	mock.ExpectExec(queryString).WillReturnError(sql.ErrConnDone)
}

func (s *UserUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository update user failure")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserUpdateFailureRepositoryMock(mock)
	_, err := repo.Update(context.Background(), user)
	t.Assert().ErrorIs(err, errs.ErrUpdateFailed)
}

func TestUserUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Update", new(UserUpdateSuite))
}

type UserDeleteSuite struct {
	UserSuite
}

func (s *UserDeleteSuite) UserDeleteSuccessRepositoryMock(mock sqlmock.Sqlmock, userID domain.ID) {
	mock.ExpectExec(repository.UserDeleteQuery).WithArgs(userID).WillReturnResult(sqlmock.NewResult(1, 1))
}

func (s *UserDeleteSuite) TestDelete_Success(t provider.T) {
	t.Parallel()
	t.Title("Repository delete user success")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserDeleteSuccessRepositoryMock(mock, user.ID)
	err := repo.Delete(context.Background(), user.ID)
	t.Assert().Nil(err)
}

func (s *UserDeleteSuite) UserDeleteFailureRepositoryMock(mock sqlmock.Sqlmock) {
	mock.ExpectExec(repository.UserDeleteQuery).WillReturnError(sql.ErrConnDone)
}

func (s *UserDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Parallel()
	t.Title("Repository delete user failure")
	repo, mock := NewUserRepository()
	user := NewUserBuilder().Build()
	s.UserDeleteFailureRepositoryMock(mock)
	err := repo.Delete(context.Background(), user.ID)
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestUserDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Repository Delete", new(UserDeleteSuite))
}
