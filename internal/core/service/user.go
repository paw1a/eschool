package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

type UserService struct {
	repo port.IUserRepository
}

func NewUserService(repo port.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) FindAll(ctx context.Context) ([]domain.User, error) {
	return u.repo.FindAll(ctx)
}

func (u *UserService) FindByID(ctx context.Context, userID int64) (domain.User, error) {
	return u.repo.FindByID(ctx, userID)
}

func (u *UserService) FindByCredentials(ctx context.Context, signInDTO dto.UserCredentials) (domain.User, error) {
	return u.repo.FindByCredentials(ctx, signInDTO.Email, signInDTO.Password)
}

func (u *UserService) FindUserInfo(ctx context.Context, userID int64) (dto.UserInfo, error) {
	return u.repo.FindUserInfo(ctx, userID)
}

func (u *UserService) Create(ctx context.Context, userDTO dto.CreateUserDTO) (domain.User, error) {
	return u.repo.Create(ctx, userDTO)
}

func (u *UserService) Update(ctx context.Context, userID int64, userDTO dto.UpdateUserDTO) (domain.User, error) {
	return u.repo.Update(ctx, userID, userDTO)
}

func (u *UserService) Delete(ctx context.Context, userID int64) error {
	return u.repo.Delete(ctx, userID)
}
