package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/zap"
)

type SchoolService struct {
	repo   port.ISchoolRepository
	logger *zap.Logger
}

func NewSchoolService(repo port.ISchoolRepository, logger *zap.Logger) *SchoolService {
	return &SchoolService{
		repo:   repo,
		logger: logger,
	}
}

func (s *SchoolService) FindAll(ctx context.Context) ([]domain.School, error) {
	schools, err := s.repo.FindAll(ctx)
	if err != nil {
		s.logger.Error("failed to get all schools", zap.Error(err))
		return nil, err
	}
	return schools, nil
}

func (s *SchoolService) FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error) {
	school, err := s.repo.FindByID(ctx, schoolID)
	if err != nil {
		s.logger.Error("failed to get school by id", zap.Error(err),
			zap.String("schoolID", schoolID.String()))
		return domain.School{}, err
	}
	return school, nil
}

func (s *SchoolService) FindSchoolCourses(ctx context.Context, schoolID domain.ID) ([]domain.Course, error) {
	courses, err := s.repo.FindSchoolCourses(ctx, schoolID)
	if err != nil {
		s.logger.Error("failed to get school courses", zap.Error(err),
			zap.String("schoolID", schoolID.String()))
		return nil, err
	}
	return courses, nil
}

func (s *SchoolService) FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error) {
	schools, err := s.repo.FindUserSchools(ctx, userID)
	if err != nil {
		s.logger.Error("failed to get school by id", zap.Error(err),
			zap.String("userID", userID.String()))
		return nil, err
	}
	return schools, nil
}

func (s *SchoolService) FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error) {
	teachers, err := s.repo.FindSchoolTeachers(ctx, schoolID)
	if err != nil {
		s.logger.Error("failed to get school teachers", zap.Error(err),
			zap.String("schoolID", schoolID.String()))
		return nil, err
	}
	return teachers, nil
}

func (s *SchoolService) IsSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) (bool, error) {
	flag, err := s.repo.IsSchoolTeacher(ctx, schoolID, teacherID)
	if err != nil {
		s.logger.Error("failed to check if user is school teacher", zap.Error(err),
			zap.String("schoolID", schoolID.String()), zap.String("userID", teacherID.String()))
		return false, err
	}
	return flag, nil
}

func (s *SchoolService) AddSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) error {
	err := s.repo.AddSchoolTeacher(ctx, schoolID, teacherID)
	if err != nil {
		s.logger.Error("failed to add school teacher", zap.Error(err),
			zap.String("schoolID", schoolID.String()), zap.String("userID", teacherID.String()))
		return err
	}

	s.logger.Info("school teacher is added",
		zap.String("schoolID", schoolID.String()), zap.String("userID", teacherID.String()))
	return nil
}

func (s *SchoolService) CreateUserSchool(ctx context.Context, userID domain.ID,
	param port.CreateSchoolParam) (domain.School, error) {
	school, err := s.repo.Create(ctx, domain.School{
		ID:          domain.NewID(),
		OwnerID:     userID,
		Name:        param.Name,
		Description: param.Description,
	})
	if err != nil {
		s.logger.Error("failed to create user school", zap.Error(err),
			zap.String("userID", userID.String()))
		return domain.School{}, err
	}

	s.logger.Info("school is successfully created",
		zap.String("schoolID", school.ID.String()), zap.String("userID", userID.String()))
	return school, nil
}

func (s *SchoolService) Update(ctx context.Context, schoolID domain.ID,
	param port.UpdateSchoolParam) (domain.School, error) {
	school, err := s.repo.FindByID(ctx, schoolID)
	if err != nil {
		s.logger.Error("failed to find school by id", zap.Error(err),
			zap.String("schoolID", schoolID.String()))
		return domain.School{}, err
	}

	if param.Description.Valid {
		school.Description = param.Description.String
	}

	school, err = s.repo.Update(ctx, school)
	if err != nil {
		s.logger.Error("failed to update school", zap.Error(err),
			zap.String("schoolID", schoolID.String()))
		return domain.School{}, err
	}

	s.logger.Info("school is successfully updated",
		zap.String("schoolID", school.ID.String()), zap.String("userID", school.OwnerID.String()))
	return school, nil
}

func (s *SchoolService) Delete(ctx context.Context, schoolID domain.ID) error {
	err := s.repo.Delete(ctx, schoolID)
	if err != nil {
		s.logger.Error("failed to delete school", zap.Error(err),
			zap.String("schoolID", schoolID.String()))
		return err
	}

	s.logger.Info("school is successfully deleted",
		zap.String("schoolID", schoolID.String()))
	return nil
}
