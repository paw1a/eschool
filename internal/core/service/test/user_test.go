package test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"go.uber.org/zap"
)

type UserSuite struct {
	suite.Suite
	logger         *zap.Logger
	userRepository *mocks.UserRepository
	userService    *service.UserService
}

func (s *UserSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()

	s.userRepository = mocks.NewUserRepository(t)
	s.userService = service.NewUserService(s.userRepository, s.logger)
}

// FindAll Suite
type UserFindAllSuite struct {
	UserSuite
}

func UserFindAllSuccessRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("FindAll", context.Background()).
		Return([]domain.User{NewUserBuilder().Build()}, nil)
}

func (s *UserFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Find all success")
	UserFindAllSuccessRepositoryMock(s.userRepository)
	_, err := s.userService.FindAll(context.Background())
	t.Assert().Nil(err)
}

func UserFindAllFailureRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("FindAll", context.Background()).
		Return(nil, errs.ErrNotExist)
}

func (s *UserFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find all failure")
	UserFindAllFailureRepositoryMock(s.userRepository)
	_, err := s.userService.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindAll", new(UserFindAllSuite))
}

// FindByID Suite
type UserFindByIDSuite struct {
	UserSuite
}

func UserFindByIDSuccessRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("FindByID", context.Background(), userID).
		Return(NewUserBuilder().WithID(userID).Build(), nil)
}

func (s *UserFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Find by id success")
	userID := domain.NewID()
	UserFindByIDSuccessRepositoryMock(s.userRepository, userID)
	user, err := s.userService.FindByID(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(userID, user.ID)
}

func UserFindByIDFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("FindByID", context.Background(), userID).
		Return(domain.User{}, errs.ErrNotExist)
}

func (s *UserFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find by id failure")
	userID := domain.NewID()
	UserFindByIDFailureRepositoryMock(s.userRepository, userID)
	_, err := s.userService.FindByID(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindByID", new(UserFindByIDSuite))
}

// FindByCredentials Suite
type UserFindByCredentialsSuite struct {
	UserSuite
}

func UserFindByCredentialsSuccessRepositoryMock(repository *mocks.UserRepository, email, password string) {
	repository.
		On("FindByCredentials", context.Background(), email, password).
		Return(NewUserBuilder().
			WithEmail(email).
			WithPassword(password).
			Build(), nil)
}

func (s *UserFindByCredentialsSuite) TestFindByCredentials_Success(t provider.T) {
	t.Parallel()
	t.Title("Find by credentials success")
	email := "test@example.com"
	password := "password"
	UserFindByCredentialsSuccessRepositoryMock(s.userRepository, email, password)
	user, err := s.userService.FindByCredentials(context.Background(),
		port.UserCredentials{Email: email, Password: password})
	t.Assert().Nil(err)
	t.Assert().Equal(email, user.Email)
}

func UserFindByCredentialsFailureRepositoryMock(repository *mocks.UserRepository, email, password string) {
	repository.
		On("FindByCredentials", context.Background(), email, password).
		Return(domain.User{}, errs.ErrInvalidCredentials)
}

func (s *UserFindByCredentialsSuite) TestFindByCredentials_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find by credentials failure")
	email := "test@example.com"
	password := "password"
	UserFindByCredentialsFailureRepositoryMock(s.userRepository, email, password)
	_, err := s.userService.FindByCredentials(context.Background(),
		port.UserCredentials{Email: email, Password: password})
	t.Assert().ErrorIs(err, errs.ErrInvalidCredentials)
}

func TestUserFindByCredentialsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindByCredentials", new(UserFindByCredentialsSuite))
}

// Create Suite
type UserCreateSuite struct {
	UserSuite
}

func UserCreateSuccessRepositoryMock(repository *mocks.UserRepository, email string) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewUserBuilder().WithEmail(email).Build(), nil)
}

func (s *UserCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Create success")
	param := NewCreateUserParamBuilder().Build()
	UserCreateSuccessRepositoryMock(s.userRepository, param.Email)
	user, err := s.userService.Create(context.Background(), param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Email, user.Email)
}

func UserCreateFailureRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.User{}, errs.ErrNotUniqueEmail)
}

func (s *UserCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Create failure")
	param := NewCreateUserParamBuilder().Build()
	UserCreateFailureRepositoryMock(s.userRepository)
	_, err := s.userService.Create(context.Background(), param)
	t.Assert().ErrorIs(err, errs.ErrNotUniqueEmail)
}

func TestUserCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Create", new(UserCreateSuite))
}

// Update Suite
type UserUpdateSuite struct {
	UserSuite
}

func UserUpdateSuccessRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("FindByID", context.Background(), userID).
		Return(NewUserBuilder().WithID(userID).Build(), nil)
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewUserBuilder().WithID(userID).Build(), nil)
}

func (s *UserUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Update success")
	userID := domain.NewID()
	param := NewUpdateUserParamBuilder().WithName("name").Build()
	UserUpdateSuccessRepositoryMock(s.userRepository, userID)
	_, err := s.userService.Update(context.Background(), userID, param)
	t.Assert().Nil(err)
}

func UserUpdateFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("FindByID", context.Background(), userID).
		Return(domain.User{}, errs.ErrNotExist)
}

func (s *UserUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Update failure")
	userID := domain.NewID()
	param := NewUpdateUserParamBuilder().WithName("name").Build()
	UserUpdateFailureRepositoryMock(s.userRepository, userID)
	_, err := s.userService.Update(context.Background(), userID, param)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestUserUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Update", new(UserUpdateSuite))
}

// Delete Suite
type UserDeleteSuite struct {
	UserSuite
}

func UserDeleteSuccessRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("Delete", context.Background(), userID).
		Return(nil)
}

func (s *UserDeleteSuite) TestDelete_Success(t provider.T) {
	t.Parallel()
	t.Title("Delete success")
	userID := domain.NewID()
	UserDeleteSuccessRepositoryMock(s.userRepository, userID)
	err := s.userService.Delete(context.Background(), userID)
	t.Assert().Nil(err)
}

func UserDeleteFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("Delete", context.Background(), userID).
		Return(errs.ErrDeleteFailed)
}

func (s *UserDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Parallel()
	t.Title("Delete failure")
	userID := domain.NewID()
	UserDeleteFailureRepositoryMock(s.userRepository, userID)
	err := s.userService.Delete(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestUserDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Delete", new(UserDeleteSuite))
}
