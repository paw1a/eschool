package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
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

func (u *UserService) FindByID(ctx context.Context, userID domain.ID) (domain.User, error) {
	return u.repo.FindByID(ctx, userID)
}

func (u *UserService) FindByCredentials(ctx context.Context,
	credentials port.UserCredentials) (domain.User, error) {
	return u.repo.FindByCredentials(ctx, credentials.Email, credentials.Password)
}

func (u *UserService) FindUserInfo(ctx context.Context, userID domain.ID) (port.UserInfo, error) {
	return u.repo.FindUserInfo(ctx, userID)
}

func (u *UserService) Create(ctx context.Context, param port.CreateUserParam) (domain.User, error) {
	return u.repo.Create(ctx, domain.User{
		ID:        domain.NewID(),
		Name:      param.Name,
		Surname:   param.Surname,
		Phone:     param.Phone,
		City:      param.City,
		AvatarUrl: param.AvatarUrl,
		Email:     param.Email,
		Password:  param.Password,
	})
}

func (u *UserService) Update(ctx context.Context, userID domain.ID,
	param port.UpdateUserParam) (domain.User, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	if param.Name.Valid {
		user.Name = param.Name.String
	}
	if param.Surname.Valid {
		user.Surname = param.Surname.String
	}
	if param.City.Valid {
		user.City = param.City
	}
	if param.Phone.Valid {
		user.Phone = param.Phone
	}
	if param.AvatarUrl.Valid {
		user.AvatarUrl = param.AvatarUrl
	}

	return u.repo.Update(ctx, user)
}

func (u *UserService) Delete(ctx context.Context, userID domain.ID) error {
	return u.repo.Delete(ctx, userID)
}
