package repository

import (
	"context"
	"database/sql"
	"github.com/jackc/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/eschool/internal/adapter/repository/postgres/entity"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
)

type PostgresSchoolRepo struct {
	db *sqlx.DB
}

func NewSchoolRepo(db *sqlx.DB) *PostgresSchoolRepo {
	return &PostgresSchoolRepo{
		db: db,
	}
}

const (
	schoolFindAllQuery            = "SELECT * FROM public.school"
	schoolFindByIDQuery           = "SELECT * FROM public.school WHERE id = $1"
	schoolFindUserSchoolsQuery    = "SELECT * FROM public.school WHERE owner_id = $1"
	schoolFindSchoolTeachersQuery = "SELECT u.* FROM public.user u " +
		"JOIN public.school_teacher st on u.id = st.teacher_id " +
		"JOIN public.school s on st.school_id = s.id WHERE s.id = $1"
	schoolAddTeacherQuery = "INSERT INTO public.school_teacher (teacher_id, school_id) " +
		"VALUES ($1, $2)"
	schoolDeleteQuery = "DELETE FROM public.school WHERE id = $1"
)

func (s *PostgresSchoolRepo) FindAll(ctx context.Context) ([]domain.School, error) {
	var pgSchools []entity.PgSchool
	if err := s.db.SelectContext(ctx, &pgSchools, schoolFindAllQuery); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	schools := make([]domain.School, len(pgSchools))
	for i, school := range pgSchools {
		schools[i] = school.ToDomain()
	}
	return schools, nil
}

func (s *PostgresSchoolRepo) FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error) {
	var pgSchool entity.PgSchool
	if err := s.db.GetContext(ctx, &pgSchool, schoolFindByIDQuery, schoolID); err != nil {
		if err == sql.ErrNoRows {
			return domain.School{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.School{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return pgSchool.ToDomain(), nil
}

func (s *PostgresSchoolRepo) FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error) {
	var pgSchools []entity.PgSchool
	if err := s.db.SelectContext(ctx, &pgSchools, schoolFindUserSchoolsQuery, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	schools := make([]domain.School, len(pgSchools))
	for i, school := range pgSchools {
		schools[i] = school.ToDomain()
	}
	return schools, nil
}

func (s *PostgresSchoolRepo) FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error) {
	var pgUsers []entity.PgUser
	if err := s.db.SelectContext(ctx, &pgUsers, schoolFindSchoolTeachersQuery, schoolID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	teachers := make([]domain.User, len(pgUsers))
	for i, teacher := range pgUsers {
		teachers[i] = teacher.ToDomain()
	}
	return teachers, nil
}

func (s *PostgresSchoolRepo) AddSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) error {
	_, err := s.db.ExecContext(ctx, schoolAddTeacherQuery, teacherID, schoolID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return errors.Wrap(errs.ErrDuplicate, err.Error())
			} else {
				return errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		} else {
			return errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return nil
}

func (s *PostgresSchoolRepo) Create(ctx context.Context, school domain.School) (domain.School, error) {
	var pgSchool = entity.NewPgSchool(school)
	queryString := entity.InsertQueryString(pgSchool, "school")
	_, err := s.db.NamedExecContext(ctx, queryString, pgSchool)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.School{}, errors.Wrap(errs.ErrDuplicate, err.Error())
			} else {
				return domain.School{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.School{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	var createdSchool entity.PgSchool
	err = s.db.GetContext(ctx, &createdSchool, schoolFindByIDQuery, pgSchool.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.School{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.School{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	return createdSchool.ToDomain(), nil
}

func (s *PostgresSchoolRepo) Update(ctx context.Context, school domain.School) (domain.School, error) {
	var pgSchool = entity.NewPgSchool(school)
	queryString := entity.UpdateQueryString(pgSchool, "school")
	_, err := s.db.NamedExecContext(ctx, queryString, pgSchool)
	if err != nil {
		return domain.School{}, errors.Wrap(errs.ErrUpdateFailed, err.Error())
	}

	var updatedSchool entity.PgSchool
	err = s.db.GetContext(ctx, &updatedSchool, schoolFindByIDQuery, pgSchool.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.School{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.School{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	return updatedSchool.ToDomain(), nil
}

func (s *PostgresSchoolRepo) Delete(ctx context.Context, schoolID domain.ID) error {
	_, err := s.db.ExecContext(ctx, schoolDeleteQuery, schoolID)
	if err != nil {
		return errors.Wrap(errs.ErrDeleteFailed, err.Error())
	}
	return nil
}
