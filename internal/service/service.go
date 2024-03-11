package service

import (
	"context"
	"github.com/paw1a/eschool/internal/domain"
	"github.com/paw1a/eschool/internal/domain/dto"
)

type Users interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID int64) (domain.User, error)
	FindByCredentials(ctx context.Context, signInDTO dto.SignInDTO) (domain.User, error)
	FindUserInfo(ctx context.Context, userID int64) (domain.UserInfo, error)
	Create(ctx context.Context, userDTO dto.CreateUserDTO) (domain.User, error)
	Update(ctx context.Context, userDTO dto.UpdateUserDTO,
		userID int64) (domain.User, error)
	Delete(ctx context.Context, userID int64) error
}
