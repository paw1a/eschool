package repository

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
)

type SchoolRepository struct {
}

func (s *SchoolRepository) FindAll(ctx context.Context) ([]domain.School, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolRepository) FindByID(ctx context.Context, schoolID int64) (domain.School, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolRepository) AddSchoolTeacher(ctx context.Context, schoolID int64, teacherID int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolRepository) CreateUserSchool(ctx context.Context, schoolDTO dto.CreateSchoolDTO, userID int64) (domain.School, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolRepository) UpdateUserSchool(ctx context.Context, schoolID int64, schoolDTO dto.UpdateSchoolDTO, userID int64) (domain.School, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolRepository) Delete(ctx context.Context, schoolID int64) error {
	//TODO implement me
	panic("implement me")
}
