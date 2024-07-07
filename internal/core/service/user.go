package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/zap"
)

type UserService struct {
	repo   port.IUserRepository
	logger *zap.Logger
}

func NewUserService(repo port.IUserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserService) FindAll(ctx context.Context) ([]domain.User, error) {
	users, err := u.repo.FindAll(ctx)
	if err != nil {
		u.logger.Error("failed to find all users", zap.Error(err))
		return nil, err
	}
	return users, nil
}

func (u *UserService) FindByID(ctx context.Context, userID domain.ID) (domain.User, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil {
		u.logger.Error("failed to find user by id", zap.Error(err),
			zap.String("userID", userID.String()))
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserService) FindByCredentials(ctx context.Context,
	credentials port.UserCredentials) (domain.User, error) {
	user, err := u.repo.FindByCredentials(ctx, credentials.Email, credentials.Password)
	if err != nil {
		u.logger.Error("failed to find user by credentials", zap.Error(err),
			zap.String("email", credentials.Email))
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserService) FindUserInfo(ctx context.Context, userID domain.ID) (port.UserInfo, error) {
	info, err := u.repo.FindUserInfo(ctx, userID)
	if err != nil {
		u.logger.Error("failed to find user account", zap.Error(err),
			zap.String("userID", userID.String()))
		return port.UserInfo{}, err
	}
	return info, nil
}

func (u *UserService) Create(ctx context.Context, param port.CreateUserParam) (domain.User, error) {
	user, err := u.repo.Create(ctx, domain.User{
		ID:        domain.NewID(),
		Name:      param.Name,
		Surname:   param.Surname,
		Phone:     param.Phone,
		City:      param.City,
		AvatarUrl: param.AvatarUrl,
		Email:     param.Email,
		Password:  param.Password,
	})
	if err != nil {
		u.logger.Error("failed to create user", zap.Error(err))
		return domain.User{}, err
	}

	u.logger.Info("user is successfully created",
		zap.String("userID", user.ID.String()))
	return user, nil
}

func (u *UserService) Update(ctx context.Context, userID domain.ID,
	param port.UpdateUserParam) (domain.User, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil {
		u.logger.Error("failed to find user by id", zap.Error(err),
			zap.String("userID", userID.String()))
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

	user, err = u.repo.Update(ctx, user)
	if err != nil {
		u.logger.Error("failed to update user", zap.Error(err),
			zap.String("userID", userID.String()))
		return domain.User{}, err
	}

	u.logger.Info("user is successfully updated",
		zap.String("userID", user.ID.String()))
	return user, nil
}

func (u *UserService) Delete(ctx context.Context, userID domain.ID) error {
	err := u.repo.Delete(ctx, userID)
	if err != nil {
		u.logger.Error("failed to delete user", zap.Error(err),
			zap.String("userID", userID.String()))
		return err
	}

	u.logger.Info("user is successfully deleted",
		zap.String("userID", userID.String()))
	return nil
}
