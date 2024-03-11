package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/paw1a/eschool/internal/domain"
	"github.com/paw1a/eschool/internal/domain/dto"
)

type PostgresUsersRepo struct {
	db *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool) *PostgresUsersRepo {
	return &PostgresUsersRepo{
		db: db,
	}
}

func (u *PostgresUsersRepo) FindAll(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}

func (u *PostgresUsersRepo) FindByID(ctx context.Context, userID int64) (domain.User, error) {
	return domain.User{}, nil
}

func (u *PostgresUsersRepo) FindByCredentials(ctx context.Context, email string, password string) (domain.User, error) {
	return domain.User{}, nil
}

func (u *PostgresUsersRepo) FindUserInfo(ctx context.Context, userID int64) (domain.UserInfo, error) {
	return domain.UserInfo{}, nil
}

func (u *PostgresUsersRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (u *PostgresUsersRepo) Update(ctx context.Context, userInput dto.UpdateUserInput,
	userID int64) (domain.User, error) {
	return domain.User{}, nil
}

func (u *PostgresUsersRepo) Delete(ctx context.Context, userID int64) error {
	return nil
}
