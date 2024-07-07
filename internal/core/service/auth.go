package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/zap"
)

type AuthTokenService struct {
	authProvider port.IAuthProvider
	userRepo     port.IUserRepository
	logger       *zap.Logger
}

func NewAuthTokenService(authProvider port.IAuthProvider, userRepo port.IUserRepository,
	logger *zap.Logger) *AuthTokenService {
	return &AuthTokenService{
		authProvider: authProvider,
		userRepo:     userRepo,
		logger:       logger,
	}
}

func (a *AuthTokenService) SignIn(ctx context.Context, param port.SignInParam) (domain.AuthDetails, error) {
	user, err := a.userRepo.FindByCredentials(ctx, param.Email, param.Password)
	if err != nil {
		a.logger.Error("failed to verify sign in credentials", zap.Error(err))
		return domain.AuthDetails{}, errs.ErrInvalidCredentials
	}
	details, err := a.authProvider.CreateJWTSession(domain.AuthPayload{UserID: user.ID}, param.Fingerprint)
	if err != nil {
		a.logger.Error("failed to create new jwt session", zap.Error(err))
		return domain.AuthDetails{}, err
	}
	a.logger.Info("user successfully signed in", zap.String("userID", user.ID.String()))
	return details, nil
}

func (a *AuthTokenService) SignUp(ctx context.Context, param port.SignUpParam) error {
	_, err := a.userRepo.FindByEmail(ctx, param.Email)
	if err == nil {
		a.logger.Error("failed to find user by email", zap.Error(err), zap.String("email", param.Email))
		return errs.ErrNotUniqueEmail
	}

	user, err := a.userRepo.Create(ctx, domain.User{
		ID:        domain.NewID(),
		Name:      param.Name,
		Surname:   param.Surname,
		Email:     param.Email,
		Password:  param.Password,
		Phone:     param.Phone,
		City:      param.City,
		AvatarUrl: param.AvatarUrl,
	})
	if err != nil {
		a.logger.Error("failed to sign up user", zap.Error(err))
		return err
	}

	a.logger.Info("user successfully signed up", zap.String("userID", user.ID.String()))
	return nil
}

func (a *AuthTokenService) LogOut(ctx context.Context, refreshToken domain.Token) error {
	err := a.authProvider.DeleteJWTSession(refreshToken)
	if err != nil {
		a.logger.Error("failed to log out user", zap.Error(err), zap.String("refreshToken", refreshToken.String()))
		return err
	}

	a.logger.Info("user successfully logged out", zap.String("refreshToken", refreshToken.String()))
	return nil
}

func (a *AuthTokenService) Refresh(ctx context.Context, refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	details, err := a.authProvider.RefreshJWTSession(refreshToken, fingerprint)
	if err != nil {
		a.logger.Error("failed to refresh user session", zap.Error(err),
			zap.String("refreshToken", refreshToken.String()))
		return domain.AuthDetails{}, err
	}
	return details, nil
}

func (a *AuthTokenService) Verify(ctx context.Context, accessToken domain.Token) error {
	_, err := a.authProvider.VerifyJWTToken(accessToken)
	if err != nil {
		a.logger.Error("failed to verify user", zap.Error(err),
			zap.String("accessToken", accessToken.String()))
		return err
	}

	a.logger.Info("user token is verified", zap.String("accessToken", accessToken.String()))
	return nil
}

func (a *AuthTokenService) Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error) {
	return a.authProvider.VerifyJWTToken(accessToken)
}
