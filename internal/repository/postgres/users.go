package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/eschool/internal/domain"
	"github.com/paw1a/eschool/internal/domain/dto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type PostgresUsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *PostgresUsersRepo {
	return &PostgresUsersRepo{
		db: db,
	}
}

const (
	usersFindAllQuery           = "SELECT * FROM public.user ORDER BY id"
	usersFindByIDQuery          = "SELECT * FROM public.user WHERE id = $1"
	usersFindByCredentialsQuery = "SELECT * FROM public.user WHERE email = $1 AND password = $2"
	usersFindUserInfoQuery      = "SELECT email, name FROM public.user WHERE id = $1"
	usersCreateQuery            = "INSERT INTO public.user (email, password, name, surname, phone, city, avatar_url) " +
		"VALUES ($1, $2, $3, $4, NULL, NULL, NULL) RETURNING *"
	usersUpdateQuery = "UPDATE public.user SET name = $1 WHERE id = $2"
	usersDeleteQuery = "DELETE FROM public.user WHERE id = $1"
)

func (u *PostgresUsersRepo) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := u.db.SelectContext(ctx, &users, usersFindAllQuery)
	if err != nil {
		return nil, errors.Wrap(err, "user repo find all")
	}
	return users, nil
}

func (u *PostgresUsersRepo) FindByID(ctx context.Context, userID int64) (domain.User, error) {
	var user domain.User
	err := u.db.GetContext(ctx, &user, usersFindByIDQuery, userID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user repo find by id")
	}
	return user, nil
}

func (u *PostgresUsersRepo) FindByCredentials(ctx context.Context, email string, password string) (domain.User, error) {
	var user domain.User
	err := u.db.GetContext(ctx, &user, usersFindByCredentialsQuery, email, password)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user repo find by credentials")
	}
	return user, nil
}

func (u *PostgresUsersRepo) FindUserInfo(ctx context.Context, userID int64) (domain.UserInfo, error) {
	var userInfo domain.UserInfo
	err := u.db.GetContext(ctx, &userInfo, usersFindUserInfoQuery, userID)
	if err != nil {
		return domain.UserInfo{}, errors.Wrap(err, "user repo find user info")
	}
	return userInfo, nil
}

func (u *PostgresUsersRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	var createdUser domain.User
	err := u.db.QueryRowxContext(ctx, usersCreateQuery,
		user.Email, user.Password, user.Name, user.Surname).StructScan(&createdUser)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user repo create")
	}
	return createdUser, nil
}

func (u *PostgresUsersRepo) Update(ctx context.Context, userInput dto.UpdateUserInput,
	userID int64) (domain.User, error) {
	var updatedUser domain.User
	err := u.db.QueryRowxContext(ctx, usersUpdateQuery, userInput.Name, userID).Scan(&updatedUser)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "user repo update")
	}
	log.Debugf("updated user: %v\n", updatedUser)
	return updatedUser, nil
}

func (u *PostgresUsersRepo) Delete(ctx context.Context, userID int64) error {
	_, err := u.db.ExecContext(ctx, usersDeleteQuery, userID)
	if err != nil {
		return errors.Wrap(err, "user repo delete")
	}
	return nil
}
