package repository

import (
	"context"
	"github.com/paw1a/eschool/internal/adapter/repository/gorm/entity"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (ur *GormUserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	gormUser := entity.NewGormUser(user)
	result := ur.db.Table("user").WithContext(ctx).Create(gormUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return domain.User{}, errors.Wrap(errs.ErrDuplicate, result.Error.Error())
		}
		return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, result.Error.Error())
	}

	var createdUser entity.GormUser
	result = ur.db.Table("user").WithContext(ctx).Take(&createdUser, "id = ?", gormUser.ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, errors.Wrap(errs.ErrNotExist, result.Error.Error())
		}
		return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, result.Error.Error())
	}

	return createdUser.ToDomain(), nil
}

func (ur *GormUserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []entity.GormUser
	result := ur.db.Table("user").WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, errors.Wrap(errs.ErrNotExist, result.Error.Error())
	}

	domainUsers := make([]domain.User, len(users))
	for i, user := range users {
		domainUsers[i] = user.ToDomain()
	}

	return domainUsers, nil
}

func (ur *GormUserRepository) FindByID(ctx context.Context, userID domain.ID) (domain.User, error) {
	var foundUser entity.GormUser
	result := ur.db.Table("user").WithContext(ctx).Take(&foundUser, "id = ?", userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, errors.Wrap(errs.ErrNotExist, result.Error.Error())
		}
		return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, result.Error.Error())
	}

	return foundUser.ToDomain(), nil
}

func (ur *GormUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *GormUserRepository) FindByCredentials(ctx context.Context, email string, password string) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *GormUserRepository) FindUserInfo(ctx context.Context, userID domain.ID) (port.UserInfo, error) {
	return port.UserInfo{}, nil
}

func (ur *GormUserRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *GormUserRepository) Delete(ctx context.Context, userID domain.ID) error {
	return nil
}
