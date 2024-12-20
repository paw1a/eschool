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

type PostgresLessonRepo struct {
	db *sqlx.DB
}

func NewLessonRepo(db *sqlx.DB) *PostgresLessonRepo {
	return &PostgresLessonRepo{
		db: db,
	}
}

const (
	LessonFindAllQuery           = "SELECT * FROM public.lesson ORDER BY id"
	LessonFindByIDQuery          = "SELECT * FROM public.lesson WHERE id = $1"
	LessonFindCourseLessonsQuery = "SELECT * FROM public.lesson WHERE course_id = $1"
	LessonFindLessonTestsQuery   = "SELECT * FROM public.test WHERE lesson_id = $1"
	LessonDeleteQuery            = "DELETE FROM public.lesson WHERE id = $1"
	LessonDeleteLessonTestsQuery = "DELETE FROM public.test WHERE lesson_id = $1"
)

func (p *PostgresLessonRepo) FindAll(ctx context.Context) ([]domain.Lesson, error) {
	var pgLessons []entity.PgLesson
	if err := p.db.SelectContext(ctx, &pgLessons, LessonFindAllQuery); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	lessons := make([]domain.Lesson, len(pgLessons))
	for i, lesson := range pgLessons {
		lessons[i] = lesson.ToDomain()
		if lesson.Type == entity.PgLessonPractice {
			tests, err := p.FindLessonTests(ctx, lessons[i].ID)
			if err != nil {
				return nil, err
			}
			lessons[i].Tests = tests
		}
	}
	return lessons, nil
}

func (p *PostgresLessonRepo) FindByID(ctx context.Context, lessonID domain.ID) (domain.Lesson, error) {
	var pgLesson entity.PgLesson
	if err := p.db.GetContext(ctx, &pgLesson, LessonFindByIDQuery, lessonID); err != nil {
		if err == sql.ErrNoRows {
			return domain.Lesson{}, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return domain.Lesson{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}
	lesson := pgLesson.ToDomain()
	if pgLesson.Type == entity.PgLessonPractice {
		tests, err := p.FindLessonTests(ctx, lessonID)
		if err != nil {
			return domain.Lesson{}, err
		}
		lesson.Tests = tests
	}
	return lesson, nil
}

func (p *PostgresLessonRepo) FindCourseLessons(ctx context.Context,
	courseID domain.ID) ([]domain.Lesson, error) {
	var pgLessons []entity.PgLesson
	if err := p.db.SelectContext(ctx, &pgLessons, LessonFindCourseLessonsQuery, courseID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	lessons := make([]domain.Lesson, len(pgLessons))
	for i, lesson := range pgLessons {
		lessons[i] = lesson.ToDomain()
		if lesson.Type == entity.PgLessonPractice {
			tests, err := p.FindLessonTests(ctx, lessons[i].ID)
			if err != nil {
				return nil, err
			}
			lessons[i].Tests = tests
		}
	}
	return lessons, nil
}

func (p *PostgresLessonRepo) FindLessonTests(ctx context.Context, lessonID domain.ID) ([]domain.Test, error) {
	var pgTests []entity.PgTest
	if err := p.db.SelectContext(ctx, &pgTests, LessonFindLessonTestsQuery, lessonID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(errs.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	tests := make([]domain.Test, len(pgTests))
	for i, test := range pgTests {
		tests[i] = test.ToDomain()
	}
	return tests, nil
}

func (p *PostgresLessonRepo) Create(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return domain.Lesson{}, errors.Wrap(errs.ErrTransactionError, err.Error())
	}

	var pgLesson = entity.NewPgLesson(lesson)
	queryString := entity.InsertQueryString(pgLesson, "lesson")
	_, err = tx.NamedExecContext(ctx, queryString, pgLesson)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.Lesson{}, errors.Wrap(errs.ErrDuplicate, err.Error())
			} else {
				return domain.Lesson{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.Lesson{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
		}
	}

	if pgLesson.Type == entity.PgLessonPractice {
		for _, test := range lesson.Tests {
			var pgTest = entity.NewPgTest(test)
			queryString := entity.InsertQueryString(pgTest, "test")
			_, err = tx.NamedExecContext(ctx, queryString, pgTest)
			if err != nil {
				tx.Rollback()
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					if pgErr.Code == PgUniqueViolationCode {
						return domain.Lesson{}, errors.Wrap(errs.ErrDuplicate, err.Error())
					} else {
						return domain.Lesson{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
					}
				} else {
					return domain.Lesson{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
				}
			}
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return domain.Lesson{}, errors.Wrap(errs.ErrTransactionError, err.Error())
	}

	return p.FindByID(ctx, lesson.ID)
}

func (p *PostgresLessonRepo) Update(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return domain.Lesson{}, errors.Wrap(errs.ErrTransactionError, err.Error())
	}

	var pgLesson = entity.NewPgLesson(lesson)
	queryString := entity.UpdateQueryString(pgLesson, "lesson")
	_, err = tx.NamedExecContext(ctx, queryString, pgLesson)
	if err != nil {
		tx.Rollback()
		return domain.Lesson{}, errors.Wrap(errs.ErrUpdateFailed, err.Error())
	}

	if pgLesson.Type == entity.PgLessonPractice {
		_, err = tx.ExecContext(ctx, LessonDeleteLessonTestsQuery, lesson.ID)
		if err != nil {
			tx.Rollback()
			return domain.Lesson{}, errors.Wrap(errs.ErrUpdateFailed, err.Error())
		}

		for _, test := range lesson.Tests {
			var pgTest = entity.NewPgTest(test)
			queryString := entity.InsertQueryString(pgTest, "test")
			_, err = tx.NamedExecContext(ctx, queryString, pgTest)
			if err != nil {
				tx.Rollback()
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					if pgErr.Code == PgUniqueViolationCode {
						return domain.Lesson{}, errors.Wrap(errs.ErrDuplicate, err.Error())
					} else {
						return domain.Lesson{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
					}
				} else {
					return domain.Lesson{}, errors.Wrap(errs.ErrPersistenceFailed, err.Error())
				}
			}
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return domain.Lesson{}, errors.Wrap(errs.ErrTransactionError, err.Error())
	}

	return p.FindByID(ctx, lesson.ID)
}

func (p *PostgresLessonRepo) Delete(ctx context.Context, lessonID domain.ID) error {
	_, err := p.db.ExecContext(ctx, LessonDeleteQuery, lessonID)
	if err != nil {
		return errors.Wrap(errs.ErrDeleteFailed, err.Error())
	}
	return nil
}
