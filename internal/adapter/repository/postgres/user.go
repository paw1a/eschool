package repository

import (
	"context"
	"database/sql"
	"github.com/jackc/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/eschool/internal/adapter/repository/postgres/entity"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/pkg/errors"
)

type PostgresUserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

const (
	UserFindAllQuery           = "SELECT * FROM public.user"
	UserFindByIDQuery          = "SELECT * FROM public.user WHERE id = $1"
	UserFindByEmailQuery       = "SELECT * FROM public.user WHERE email = $1"
	UserFindByCredentialsQuery = "SELECT * FROM public.user WHERE email = $1 AND password = $2"
	UserFindUserInfoQuery      = "SELECT name, surname FROM public.user WHERE id = $1"
	UserDeleteQuery            = "DELETE FROM public.user WHERE id = $1"
)

func (u *PostgresUserRepo) FindAll(ctx context.Context) ([]domain.User, error) {
	var pgUsers []entity.PgUser
	if err := u.db.SelectContext(ctx, &pgUsers, UserFindAllQuery); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	users := make([]domain.User, len(pgUsers))
	for i, user := range pgUsers {
		users[i] = user.ToDomain()
	}
	return users, nil
}

func (u *PostgresUserRepo) FindByID(ctx context.Context, userID domain.ID) (domain.User, error) {
	var pgUser entity.PgUser
	if err := u.db.GetContext(ctx, &pgUser, UserFindByIDQuery, userID); err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return pgUser.ToDomain(), nil
}

func (u *PostgresUserRepo) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var pgUser entity.PgUser
	if err := u.db.GetContext(ctx, &pgUser, UserFindByEmailQuery, email); err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return pgUser.ToDomain(), nil
}

func (u *PostgresUserRepo) FindByCredentials(ctx context.Context, email string, password string) (domain.User, error) {
	var pgUser entity.PgUser
	err := u.db.GetContext(ctx, &pgUser, UserFindByCredentialsQuery, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return pgUser.ToDomain(), nil
}

func (u *PostgresUserRepo) FindUserInfo(ctx context.Context, userID domain.ID) (port.UserInfo, error) {
	var pgUser entity.PgUser
	err := u.db.GetContext(ctx, &pgUser, UserFindUserInfoQuery, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return port.UserInfo{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return port.UserInfo{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return port.UserInfo{
		Name:    pgUser.Name,
		Surname: pgUser.Surname,
	}, nil
}

func (u *PostgresUserRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	var pgUser = entity.NewPgUser(user)
	queryString := entity.InsertQueryString(pgUser, "user")
	_, err := u.db.NamedExecContext(ctx, queryString, pgUser)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.User{}, errors.Wrap(errs.ErrDuplicate, err.Error())
			} else {
				return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	var createdUser entity.PgUser
	err = u.db.GetContext(ctx, &createdUser, UserFindByIDQuery, pgUser.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	return createdUser.ToDomain(), nil
}

func (u *PostgresUserRepo) Update(ctx context.Context, user domain.User) (domain.User, error) {
	var pgUser = entity.NewPgUser(user)
	queryString := entity.UpdateQueryString(pgUser, "user")
	_, err := u.db.NamedExecContext(ctx, queryString, pgUser)
	if err != nil {
		return domain.User{}, errors.Wrap(errs.ErrUpdateFailed, err.Error())
	}

	var updatedUser entity.PgUser
	err = u.db.GetContext(ctx, &updatedUser, UserFindByIDQuery, pgUser.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.User{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return updatedUser.ToDomain(), nil
}

func (u *PostgresUserRepo) Delete(ctx context.Context, userID domain.ID) error {
	_, err := u.db.ExecContext(ctx, UserDeleteQuery, userID)
	if err != nil {
		return errors.Wrap(errs.ErrDeleteFailed, err.Error())
	}
	return nil
}
