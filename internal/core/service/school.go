package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

type SchoolService struct {
	repo port.ISchoolRepository
}

func NewSchoolService(repo port.ISchoolRepository) *SchoolService {
	return &SchoolService{
		repo: repo,
	}
}

func (s *SchoolService) FindAll(ctx context.Context) ([]domain.School, error) {
	return s.repo.FindAll(ctx)
}

func (s *SchoolService) FindByID(ctx context.Context, schoolID int64) (domain.School, error) {
	return s.repo.FindByID(ctx, schoolID)
}

func (s *SchoolService) AddSchoolTeacher(ctx context.Context, schoolID int64, teacherID int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolService) CreateUserSchool(ctx context.Context, schoolDTO dto.CreateSchoolDTO,
	userID int64) (domain.School, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolService) UpdateUserSchool(ctx context.Context, schoolID int64,
	schoolDTO dto.UpdateSchoolDTO, userID int64) (domain.School, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SchoolService) Delete(ctx context.Context, schoolID int64) error {
	//TODO implement me
	panic("implement me")
}
