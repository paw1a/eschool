package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
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

func (s *SchoolService) FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error) {
	return s.repo.FindByID(ctx, schoolID)
}

func (s *SchoolService) FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error) {
	return s.repo.FindUserSchools(ctx, userID)
}

func (s *SchoolService) FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error) {
	return s.repo.FindSchoolTeachers(ctx, schoolID)
}

func (s *SchoolService) AddSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) error {
	return s.repo.AddSchoolTeacher(ctx, schoolID, teacherID)
}

func (s *SchoolService) CreateUserSchool(ctx context.Context, userID domain.ID,
	param port.CreateSchoolParam) (domain.School, error) {
	return s.repo.Create(ctx, domain.School{
		ID:          domain.NewID(),
		OwnerID:     userID,
		Description: param.Description,
	})
}

func (s *SchoolService) Update(ctx context.Context, schoolID domain.ID,
	param port.UpdateSchoolParam) (domain.School, error) {
	return s.repo.Update(ctx, schoolID, param)
}

func (s *SchoolService) Delete(ctx context.Context, schoolID domain.ID) error {
	return s.repo.Delete(ctx, schoolID)
}
