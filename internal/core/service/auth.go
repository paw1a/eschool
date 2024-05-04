package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
)

type AuthTokenService struct {
	authProvider port.IAuthProvider
	userRepo     port.IUserRepository
}

func NewAuthTokenService(authProvider port.IAuthProvider, userRepo port.IUserRepository) *AuthTokenService {
	return &AuthTokenService{
		authProvider: authProvider,
		userRepo:     userRepo,
	}
}

func (a *AuthTokenService) SignIn(ctx context.Context, param port.SignInParam) (domain.AuthDetails, error) {
	user, err := a.userRepo.FindByCredentials(ctx, param.Email, param.Password)
	if err != nil {
		return domain.AuthDetails{}, errs.ErrInvalidCredentials
	}
	return a.authProvider.CreateJWTSession(domain.AuthPayload{UserID: user.ID}, param.Fingerprint)
}

func (a *AuthTokenService) SignUp(ctx context.Context, param port.SignUpParam) error {
	_, err := a.userRepo.FindByEmail(ctx, param.Email)
	if err == nil {
		return errs.ErrNotUniqueEmail
	}

	_, err = a.userRepo.Create(ctx, domain.User{
		ID:        domain.NewID(),
		Name:      param.Name,
		Surname:   param.Surname,
		Email:     param.Email,
		Password:  param.Password,
		Phone:     param.Phone,
		City:      param.City,
		AvatarUrl: param.AvatarUrl,
	})
	return err
}

func (a *AuthTokenService) LogOut(ctx context.Context, refreshToken domain.Token) error {
	return a.authProvider.DeleteJWTSession(refreshToken)
}

func (a *AuthTokenService) Refresh(ctx context.Context, refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	return a.authProvider.RefreshJWTSession(refreshToken, fingerprint)
}

func (a *AuthTokenService) Verify(ctx context.Context, accessToken domain.Token) error {
	_, err := a.authProvider.VerifyJWTToken(accessToken)
	return err
}

func (a *AuthTokenService) Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error) {
	return a.authProvider.VerifyJWTToken(accessToken)
}
