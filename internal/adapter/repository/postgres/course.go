package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type PostgresCourseRepo struct {
	db *sqlx.DB
}

func NewCourseRepo(db *sqlx.DB) *PostgresCourseRepo {
	return &PostgresCourseRepo{
		db: db,
	}
}

const (
	courseFindAllQuery           = "SELECT * FROM public.user ORDER BY id"
	courseFindByIDQuery          = "SELECT * FROM public.user WHERE id = $1"
	courseFindByCredentialsQuery = "SELECT * FROM public.user WHERE email = $1 AND password = $2"
	courseFindUserInfoQuery      = "SELECT email, name, surname FROM public.user WHERE id = $1"
	courseCreateQuery            = "INSERT INTO public.user (id, email, password, name, surname, phone, city, avatar_url) " +
		"VALUES ($1, $2, $3, $4, $5, NULL, NULL, NULL) RETURNING *"
	courseUpdateQuery = "UPDATE public.user SET name = $1 WHERE id = $2"
	courseDeleteQuery = "DELETE FROM public.user WHERE id = $1"
)

func (p *PostgresCourseRepo) FindAll(ctx context.Context) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) FindCourseInfo(ctx context.Context, courseID domain.ID) (port.CourseInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) AddCourseLesson(ctx context.Context, courseID, lessonID domain.ID) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) DeleteCourseLesson(ctx context.Context, courseID, lessonID domain.ID) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) Create(ctx context.Context, course domain.Course) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) Update(ctx context.Context, courseID domain.ID,
	param port.UpdateCourseParam) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) UpdateStatus(ctx context.Context, courseID domain.ID,
	status domain.CourseStatus) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresCourseRepo) Delete(ctx context.Context, courseID domain.ID) error {
	//TODO implement me
	panic("implement me")
}
