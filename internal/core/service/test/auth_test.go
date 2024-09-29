package test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

type AuthSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *AuthSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// SignIn Suite
type AuthSignInSuite struct {
	UserSuite
}

func AuthSignInSuccessRepositoryMock(repository *mocks.UserRepository, provider *mocks.AuthProvider) {
	repository.
		On("FindByCredentials", context.Background(), mock.Anything, mock.Anything).
		Return(NewUserBuilder().Build(), nil)
	provider.
		On("CreateJWTSession", mock.Anything, mock.Anything).
		Return(domain.AuthDetails{}, nil)
}

func (s *AuthSignInSuite) TestSignIn_Success(t provider.T) {
	t.Parallel()
	t.Title("Sign in success")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthSignInSuccessRepositoryMock(userRepository, provider)
	_, err := authService.SignIn(context.Background(),
		port.SignInParam{Email: "email", Password: "password"})
	t.Assert().Nil(err)
}

func AuthSignInFailureRepositoryMock(repository *mocks.UserRepository, provider *mocks.AuthProvider) {
	repository.
		On("FindByCredentials", context.Background(), mock.Anything, mock.Anything).
		Return(domain.User{}, errs.ErrInvalidCredentials)
}

func (s *AuthSignInSuite) TestSignIn_Failure(t provider.T) {
	t.Parallel()
	t.Title("Sign in failure")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthSignInFailureRepositoryMock(userRepository, provider)
	_, err := authService.SignIn(context.Background(),
		port.SignInParam{Email: "email", Password: "password"})
	t.Assert().ErrorIs(err, errs.ErrInvalidCredentials)
}

func TestAuthSignInSuite(t *testing.T) {
	suite.RunNamedSuite(t, "SignIn", new(AuthSignInSuite))
}

// SignUp Suite
type AuthSignUpSuite struct {
	UserSuite
}

func AuthSignUpSuccessRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("FindByEmail", context.Background(), mock.Anything).
		Return(domain.User{}, errs.ErrNotUniqueEmail)
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewUserBuilder().Build(), nil)
}

func (s *AuthSignUpSuite) TestSignUp_Success(t provider.T) {
	t.Parallel()
	t.Title("Sign up success")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthSignUpSuccessRepositoryMock(userRepository)
	err := authService.SignUp(context.Background(),
		port.SignUpParam{Email: "email", Password: "password"})
	t.Assert().Nil(err)
}

func AuthSignUpFailureRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("FindByEmail", context.Background(), mock.Anything).
		Return(NewUserBuilder().Build(), nil)
}

func (s *AuthSignUpSuite) TestSignUp_Failure(t provider.T) {
	t.Parallel()
	t.Title("Sign up failure")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthSignUpFailureRepositoryMock(userRepository)
	err := authService.SignUp(context.Background(),
		port.SignUpParam{Email: "email", Password: "password"})
	t.Assert().ErrorIs(err, errs.ErrNotUniqueEmail)
}

func TestAuthSignUpSuite(t *testing.T) {
	suite.RunNamedSuite(t, "SignUp", new(AuthSignUpSuite))
}

// LogOut Suite
type AuthLogOutSuite struct {
	UserSuite
}

func AuthLogOutSuccessRepositoryMock(provider *mocks.AuthProvider) {
	provider.
		On("DeleteJWTSession", mock.Anything).
		Return(nil)
}

func (s *AuthLogOutSuite) TestLogOut_Success(t provider.T) {
	t.Parallel()
	t.Title("Log out success")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthLogOutSuccessRepositoryMock(provider)
	err := authService.LogOut(context.Background(), "token")
	t.Assert().Nil(err)
}

func AuthLogOutFailureRepositoryMock(repository *mocks.UserRepository, provider *mocks.AuthProvider) {
	provider.
		On("DeleteJWTSession", mock.Anything).
		Return(errs.ErrDeleteFailed)
}

func (s *AuthLogOutSuite) TestLogOut_Failure(t provider.T) {
	t.Parallel()
	t.Title("Sign in failure")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthLogOutFailureRepositoryMock(userRepository, provider)
	err := authService.LogOut(context.Background(), "token")
	t.Assert().ErrorIs(err, errs.ErrDeleteFailed)
}

func TestAuthLogOutSuite(t *testing.T) {
	suite.RunNamedSuite(t, "LogOut", new(AuthLogOutSuite))
}

// Refresh Suite
type AuthRefreshSuite struct {
	UserSuite
}

func AuthRefreshSuccessRepositoryMock(provider *mocks.AuthProvider) {
	provider.
		On("RefreshJWTSession", mock.Anything, mock.Anything).
		Return(domain.AuthDetails{}, nil)
}

func (s *AuthRefreshSuite) TestRefresh_Success(t provider.T) {
	t.Parallel()
	t.Title("Refresh token success")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthRefreshSuccessRepositoryMock(provider)
	_, err := authService.Refresh(context.Background(), "token", "fingerprint")
	t.Assert().Nil(err)
}

func AuthRefreshFailureRepositoryMock(repository *mocks.UserRepository, provider *mocks.AuthProvider) {
	provider.
		On("RefreshJWTSession", mock.Anything, mock.Anything).
		Return(domain.AuthDetails{}, errs.ErrAuthSessionIsNotPresent)
}

func (s *AuthRefreshSuite) TestRefresh_Failure(t provider.T) {
	t.Parallel()
	t.Title("Refresh token failure")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthRefreshFailureRepositoryMock(userRepository, provider)
	_, err := authService.Refresh(context.Background(), "token", "fingerprint")
	t.Assert().ErrorIs(err, errs.ErrAuthSessionIsNotPresent)
}

func TestAuthRefreshSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Refresh", new(AuthRefreshSuite))
}

// Verify Suite
type AuthVerifySuite struct {
	UserSuite
}

func AuthVerifySuccessRepositoryMock(provider *mocks.AuthProvider) {
	provider.
		On("VerifyJWTToken", mock.Anything).
		Return(domain.AuthPayload{}, nil)
}

func (s *AuthVerifySuite) TestVerify_Success(t provider.T) {
	t.Parallel()
	t.Title("Verify token success")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthVerifySuccessRepositoryMock(provider)
	err := authService.Verify(context.Background(), "token")
	t.Assert().Nil(err)
}

func AuthVerifyFailureRepositoryMock(repository *mocks.UserRepository, provider *mocks.AuthProvider) {
	provider.
		On("VerifyJWTToken", mock.Anything).
		Return(domain.AuthPayload{}, errs.ErrInvalidTokenSignMethod)
}

func (s *AuthVerifySuite) TestVerify_Failure(t provider.T) {
	t.Parallel()
	t.Title("Verify token failure")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthVerifyFailureRepositoryMock(userRepository, provider)
	err := authService.Verify(context.Background(), "token")
	t.Assert().ErrorIs(err, errs.ErrInvalidTokenSignMethod)
}

func TestAuthVerifySuite(t *testing.T) {
	suite.RunNamedSuite(t, "Verify", new(AuthVerifySuite))
}

// Payload Suite
type AuthPayloadSuite struct {
	UserSuite
}

func AuthPayloadSuccessRepositoryMock(provider *mocks.AuthProvider) {
	provider.
		On("VerifyJWTToken", mock.Anything).
		Return(domain.AuthPayload{}, nil)
}

func (s *AuthPayloadSuite) TestPayload_Success(t provider.T) {
	t.Parallel()
	t.Title("Payload token success")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthPayloadSuccessRepositoryMock(provider)
	_, err := authService.Payload(context.Background(), "token")
	t.Assert().Nil(err)
}

func AuthPayloadFailureRepositoryMock(repository *mocks.UserRepository, provider *mocks.AuthProvider) {
	provider.
		On("VerifyJWTToken", mock.Anything).
		Return(domain.AuthPayload{}, errs.ErrInvalidTokenSignMethod)
}

func (s *AuthPayloadSuite) TestPayload_Failure(t provider.T) {
	t.Parallel()
	t.Title("Payload token failure")
	userRepository := mocks.NewUserRepository(t)
	provider := mocks.NewAuthProvider(t)
	authService := service.NewAuthTokenService(provider, userRepository, s.logger)
	AuthPayloadFailureRepositoryMock(userRepository, provider)
	_, err := authService.Payload(context.Background(), "token")
	t.Assert().ErrorIs(err, errs.ErrInvalidTokenSignMethod)
}

func TestAuthPayloadSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Payload", new(AuthPayloadSuite))
}
