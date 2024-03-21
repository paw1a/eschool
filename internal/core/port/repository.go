package port

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
)

type IUserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID int64) (domain.User, error)
	FindByCredentials(ctx context.Context, email string, password string) (domain.User, error)
	FindUserInfo(ctx context.Context, userID int64) (domain.UserInfo, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, userInput dto.UpdateUserInput,
		userID int64) (domain.User, error)
	Delete(ctx context.Context, userID int64) error
}
