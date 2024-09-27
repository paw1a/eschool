package test

import (
	"context"
	"github.com/guregu/null"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"go.uber.org/zap"
)

type SchoolSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *SchoolSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// FindAll Suite
type SchoolFindAllSuite struct {
	SchoolSuite
}

func SchoolFindAllSuccessRepositoryMock(repository *mocks.SchoolRepository) {
	repository.
		On("FindAll", context.Background()).
		Return([]domain.School{NewSchoolBuilder().Build()}, nil)
}

func (s *SchoolFindAllSuite) TestFindAll_Success(t provider.T) {
	t.Parallel()
	t.Title("Find all schools success")
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolFindAllSuccessRepositoryMock(schoolRepository)
	_, err := schoolService.FindAll(context.Background())
	t.Assert().Nil(err)
}

func SchoolFindAllFailureRepositoryMock(repository *mocks.SchoolRepository) {
	repository.
		On("FindAll", context.Background()).
		Return(nil, errs.ErrNotExist)
}

func (s *SchoolFindAllSuite) TestFindAll_Failure(t provider.T) {
	t.Parallel()
	t.Title("Find all schools failure")
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolFindAllFailureRepositoryMock(schoolRepository)
	_, err := schoolService.FindAll(context.Background())
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindAll", new(SchoolFindAllSuite))
}

// FindByID Suite
type SchoolFindByIDSuite struct {
	SchoolSuite
}

func SchoolFindByIDSuccessRepositoryMock(repository *mocks.SchoolRepository, schoolID domain.ID) {
	repository.
		On("FindByID", context.Background(), schoolID).
		Return(NewSchoolBuilder().WithID(schoolID).Build(), nil)
}

func (s *SchoolFindByIDSuite) TestFindByID_Success(t provider.T) {
	t.Title("Find school by id success")
	schoolID := domain.NewID()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolFindByIDSuccessRepositoryMock(schoolRepository, schoolID)
	school, err := schoolService.FindByID(context.Background(), schoolID)
	t.Assert().Nil(err)
	t.Assert().Equal(schoolID, school.ID)
}

func SchoolFindByIDFailureRepositoryMock(repository *mocks.SchoolRepository, schoolID domain.ID) {
	repository.
		On("FindByID", context.Background(), schoolID).
		Return(domain.School{}, errs.ErrNotExist)
}

func (s *SchoolFindByIDSuite) TestFindByID_Failure(t provider.T) {
	t.Title("Find school by id failure")
	schoolID := domain.NewID()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolFindByIDFailureRepositoryMock(schoolRepository, schoolID)
	_, err := schoolService.FindByID(context.Background(), schoolID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindByIDSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindByID", new(SchoolFindByIDSuite))
}

// FindUserSchools Suite
type SchoolFindUserSchoolsSuite struct {
	SchoolSuite
}

func SchoolFindUserSchoolsSuccessRepositoryMock(repository *mocks.SchoolRepository, userID domain.ID) {
	repository.
		On("FindUserSchools", context.Background(), userID).
		Return([]domain.School{NewSchoolBuilder().Build()}, nil)
}

func (s *SchoolFindUserSchoolsSuite) TestFindUserSchools_Success(t provider.T) {
	t.Title("Find user schools success")
	userID := domain.NewID()
	school := NewSchoolBuilder().Build()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolFindUserSchoolsSuccessRepositoryMock(schoolRepository, userID)
	schools, err := schoolService.FindUserSchools(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(schools[0].Name, school.Name)
}

func SchoolFindUserSchoolsFailureRepositoryMock(repository *mocks.SchoolRepository, userID domain.ID) {
	repository.
		On("FindUserSchools", context.Background(), userID).
		Return([]domain.School{{}}, errs.ErrNotExist)
}

func (s *SchoolFindUserSchoolsSuite) TestFindUserSchools_Failure(t provider.T) {
	t.Title("Find user schools failure")
	userID := domain.NewID()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolFindUserSchoolsFailureRepositoryMock(schoolRepository, userID)
	_, err := schoolService.FindUserSchools(context.Background(), userID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolFindUserSchoolsSuite(t *testing.T) {
	suite.RunNamedSuite(t, "FindUserSchools", new(SchoolFindUserSchoolsSuite))
}

// Create Suite
type SchoolCreateSuite struct {
	SchoolSuite
}

func SchoolCreateSuccessRepositoryMock(repository *mocks.SchoolRepository, name string) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewSchoolBuilder().WithName(name).Build(), nil)
}

func (s *SchoolCreateSuite) TestCreate_Success(t provider.T) {
	t.Title("Create school success")
	userID := domain.NewID()
	name := "school name"
	param := NewCreateSchoolParamBuilder().WithName(name).Build()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolCreateSuccessRepositoryMock(schoolRepository, name)
	school, err := schoolService.CreateUserSchool(context.Background(), userID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Name, school.Name)
}

func SchoolCreateFailureRepositoryMock(repository *mocks.SchoolRepository) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.School{}, errors.New("error"))
}

func (s *SchoolCreateSuite) TestCreate_Failure(t provider.T) {
	t.Title("Create school failure")
	userID := domain.NewID()
	param := NewCreateSchoolParamBuilder().Build()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolCreateFailureRepositoryMock(schoolRepository)
	_, err := schoolService.CreateUserSchool(context.Background(), userID, param)
	t.Assert().NotNil(err)
}

func TestSchoolCreateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Create", new(SchoolCreateSuite))
}

// Update Suite
type SchoolUpdateSuite struct {
	SchoolSuite
}

func SchoolUpdateSuccessRepositoryMock(repository *mocks.SchoolRepository,
	schoolID domain.ID, description string) {
	repository.
		On("FindByID", context.Background(), schoolID).
		Return(NewSchoolBuilder().WithID(schoolID).Build(), nil)
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewSchoolBuilder().WithDescription(description).Build(), nil)
}

func (s *SchoolUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Title("Update school success")
	schoolID := domain.NewID()
	description := "school description"
	param := NewUpdateSchoolParamBuilder().WithDescription(null.StringFrom(description)).Build()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolUpdateSuccessRepositoryMock(schoolRepository, schoolID, description)
	school, err := schoolService.Update(context.Background(), schoolID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Description.String, school.Description)
}

func SchoolUpdateFailureRepositoryMock(repository *mocks.SchoolRepository, schoolID domain.ID) {
	repository.
		On("FindByID", context.Background(), schoolID).
		Return(domain.School{}, errs.ErrNotExist)
}

func (s *SchoolUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Title("Update school failure")
	schoolID := domain.NewID()
	param := NewUpdateSchoolParamBuilder().Build()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolUpdateFailureRepositoryMock(schoolRepository, schoolID)
	_, err := schoolService.Update(context.Background(), schoolID, param)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolUpdateSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Update", new(SchoolUpdateSuite))
}

// Delete Suite
type SchoolDeleteSuite struct {
	SchoolSuite
}

func SchoolDeleteSuccessRepositoryMock(repository *mocks.SchoolRepository, schoolID domain.ID) {
	repository.
		On("Delete", context.Background(), schoolID).
		Return(nil)
}

func (s *SchoolDeleteSuite) TestDelete_Success(t provider.T) {
	t.Title("Delete school success")
	schoolID := domain.NewID()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolDeleteSuccessRepositoryMock(schoolRepository, schoolID)
	err := schoolService.Delete(context.Background(), schoolID)
	t.Assert().Nil(err)
}

func SchoolDeleteFailureRepositoryMock(repository *mocks.SchoolRepository, schoolID domain.ID) {
	repository.
		On("Delete", context.Background(), schoolID).
		Return(errs.ErrNotExist)
}

func (s *SchoolDeleteSuite) TestDelete_Failure(t provider.T) {
	t.Title("Delete school failure")
	schoolID := domain.NewID()
	schoolRepository := mocks.NewSchoolRepository(t)
	schoolService := service.NewSchoolService(schoolRepository, s.logger)
	SchoolDeleteFailureRepositoryMock(schoolRepository, schoolID)
	err := schoolService.Delete(context.Background(), schoolID)
	t.Assert().ErrorIs(err, errs.ErrNotExist)
}

func TestSchoolDeleteSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Delete", new(SchoolDeleteSuite))
}
