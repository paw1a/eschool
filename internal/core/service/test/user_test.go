package test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type UserSuite struct {
	suite.Suite
	logger         *zap.Logger
	userRepository *mocks.UserRepository
	userService    *service.UserService
}

func (s *UserSuite) SetupTest() {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()

	s.userRepository = mocks.NewUserRepository(s.T())
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

func (s *UserFindAllSuite) TestFindAll_Success() {
	UserFindAllSuccessRepositoryMock(s.userRepository)
	_, err := s.userService.FindAll(context.Background())
	assert.Nil(s.T(), err)
}

func UserFindAllFailureRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("FindAll", context.Background()).
		Return(nil, errs.ErrNotExist)
}

func (s *UserFindAllSuite) TestFindAll_Failure() {
	UserFindAllFailureRepositoryMock(s.userRepository)
	_, err := s.userService.FindAll(context.Background())
	assert.ErrorIs(s.T(), err, errs.ErrNotExist)
}

func TestUserFindAllSuite(t *testing.T) {
	suite.Run(t, new(UserFindAllSuite))
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

func (s *UserFindByIDSuite) TestFindByID_Success() {
	userID := domain.NewID()
	UserFindByIDSuccessRepositoryMock(s.userRepository, userID)
	user, err := s.userService.FindByID(context.Background(), userID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), userID, user.ID)
}

func UserFindByIDFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("FindByID", context.Background(), userID).
		Return(domain.User{}, errs.ErrNotExist)
}

func (s *UserFindByIDSuite) TestFindByID_Failure() {
	userID := domain.NewID()
	UserFindByIDFailureRepositoryMock(s.userRepository, userID)
	_, err := s.userService.FindByID(context.Background(), userID)
	assert.ErrorIs(s.T(), err, errs.ErrNotExist)
}

func TestUserFindByIDSuite(t *testing.T) {
	suite.Run(t, new(UserFindByIDSuite))
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

func (s *UserFindByCredentialsSuite) TestFindByCredentials_Success() {
	email := "test@example.com"
	password := "password"
	UserFindByCredentialsSuccessRepositoryMock(s.userRepository, email, password)
	user, err := s.userService.FindByCredentials(context.Background(),
		port.UserCredentials{Email: email, Password: password})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), email, user.Email)
}

func UserFindByCredentialsFailureRepositoryMock(repository *mocks.UserRepository, email, password string) {
	repository.
		On("FindByCredentials", context.Background(), email, password).
		Return(domain.User{}, errs.ErrInvalidCredentials)
}

func (s *UserFindByCredentialsSuite) TestFindByCredentials_Failure() {
	email := "test@example.com"
	password := "password"
	UserFindByCredentialsFailureRepositoryMock(s.userRepository, email, password)
	_, err := s.userService.FindByCredentials(context.Background(),
		port.UserCredentials{Email: email, Password: password})
	assert.ErrorIs(s.T(), err, errs.ErrInvalidCredentials)
}

func TestUserFindByCredentialsSuite(t *testing.T) {
	suite.Run(t, new(UserFindByCredentialsSuite))
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

func (s *UserCreateSuite) TestCreate_Success() {
	param := NewCreateUserParamBuilder().Build()
	UserCreateSuccessRepositoryMock(s.userRepository, param.Email)
	user, err := s.userService.Create(context.Background(), param)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), param.Email, user.Email)
}

func UserCreateFailureRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.User{}, errs.ErrNotUniqueEmail)
}

func (s *UserCreateSuite) TestCreate_Failure() {
	param := NewCreateUserParamBuilder().Build()
	UserCreateFailureRepositoryMock(s.userRepository)
	_, err := s.userService.Create(context.Background(), param)
	assert.ErrorIs(s.T(), err, errs.ErrNotUniqueEmail)
}

func TestUserCreateSuite(t *testing.T) {
	suite.Run(t, new(UserCreateSuite))
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

func (s *UserUpdateSuite) TestUpdate_Success() {
	userID := domain.NewID()
	param := NewUpdateUserParamBuilder().WithName("name").Build()
	UserUpdateSuccessRepositoryMock(s.userRepository, userID)
	_, err := s.userService.Update(context.Background(), userID, param)
	assert.Nil(s.T(), err)
}

func UserUpdateFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("FindByID", context.Background(), userID).
		Return(domain.User{}, errs.ErrNotExist)
}

func (s *UserUpdateSuite) TestUpdate_Failure() {
	userID := domain.NewID()
	param := NewUpdateUserParamBuilder().WithName("name").Build()
	UserUpdateFailureRepositoryMock(s.userRepository, userID)
	_, err := s.userService.Update(context.Background(), userID, param)
	assert.ErrorIs(s.T(), err, errs.ErrNotExist)
}

func TestUserUpdateSuite(t *testing.T) {
	suite.Run(t, new(UserUpdateSuite))
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

func (s *UserDeleteSuite) TestDelete_Success() {
	userID := domain.NewID()
	UserDeleteSuccessRepositoryMock(s.userRepository, userID)
	err := s.userService.Delete(context.Background(), userID)
	assert.Nil(s.T(), err)
}

func UserDeleteFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("Delete", context.Background(), userID).
		Return(errs.ErrDeleteFailed)
}

func (s *UserDeleteSuite) TestDelete_Failure() {
	userID := domain.NewID()
	UserDeleteFailureRepositoryMock(s.userRepository, userID)
	err := s.userService.Delete(context.Background(), userID)
	assert.ErrorIs(s.T(), err, errs.ErrDeleteFailed)
}

func TestUserDeleteSuite(t *testing.T) {
	suite.Run(t, new(UserDeleteSuite))
}
