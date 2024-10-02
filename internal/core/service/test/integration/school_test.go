package integration

import (
	"context"
	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/zap"
	"testing"
)

var schools = []domain.School{
	domain.School{
		ID:          domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
		OwnerID:     domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027ca"),
		Name:        "school1",
		Description: "desc1",
	},
	domain.School{
		ID:          domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
		OwnerID:     domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cb"),
		Name:        "school2",
		Description: "desc2",
	},
}

var newTeacherID = domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc")

var teachers = []domain.User{
	domain.User{
		ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027ca"),
		Name:      "Pavel",
		Surname:   "Shpakovskiy",
		Phone:     null.StringFrom("+79992233444"),
		City:      null.String{},
		AvatarUrl: null.String{},
		Email:     "paw1a@yandex.ru",
		Password:  "123",
	},
	domain.User{
		ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cb"),
		Name:      "Timur",
		Surname:   "Musin",
		Phone:     null.String{},
		City:      null.StringFrom("Moscow"),
		AvatarUrl: null.String{},
		Email:     "hanoys@mail.ru",
		Password:  "qwerty",
	},
}

var createdSchool = port.CreateSchoolParam{
	Name:        "school3",
	Description: "desc3",
}

var updatedSchool = port.UpdateSchoolParam{
	Description: null.StringFrom("updated desc"),
}

type SchoolSuite struct {
	suite.Suite
	logger    *zap.Logger
	container *postgres.PostgresContainer
	db        *sqlx.DB
}

func (s *SchoolSuite) BeforeAll(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

func (s *SchoolSuite) BeforeEach(t provider.T) {
	var err error
	s.container, err = newPostgresContainer(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	url, err := s.container.ConnectionString(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	s.db, err = newPostgresDB(url)
	if err != nil {
		t.Fatal(err)
	}
}

func (s *SchoolSuite) AfterAll(t provider.T) {
	if err := s.container.Terminate(context.Background()); err != nil {
		t.Fatalf("failed to terminate container: %s", err)
	}
}

func (s *SchoolSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *SchoolSuite) TestSchoolService_FindAll(t provider.T) {
	repo := repository.NewSchoolRepo(s.db)
	schoolService := service.NewSchoolService(repo, s.logger)
	found, err := schoolService.FindAll(context.Background())
	if err != nil {
		t.Errorf("failed to find all schools: %v", err)
	}
	t.Assert().Equal(len(found), len(schools))
	for i := range schools {
		t.Assert().Equal(schools[i], found[i])
	}
}

func (s *SchoolSuite) TestSchoolService_FindById(t provider.T) {
	repo := repository.NewSchoolRepo(s.db)
	schoolService := service.NewSchoolService(repo, s.logger)
	school, err := schoolService.FindByID(context.Background(), schools[0].ID)
	if err != nil {
		t.Errorf("failed to find school with id: %v", err)
	}
	t.Assert().Equal(school, schools[0])
}

func (s *SchoolSuite) TestSchoolService_FindSchoolTeachers(t provider.T) {
	repo := repository.NewSchoolRepo(s.db)
	schoolService := service.NewSchoolService(repo, s.logger)
	found, err := schoolService.FindSchoolTeachers(context.Background(), schools[0].ID)
	if err != nil {
		t.Errorf("failed to find school teachers: %v", err)
	}
	t.Assert().Equal(len(found), len(teachers))
	for i := range teachers {
		t.Assert().Equal(teachers[i], found[i])
	}
}

func (s *SchoolSuite) TestSchoolService_IsSchoolTeacher(t provider.T) {
	repo := repository.NewSchoolRepo(s.db)
	schoolService := service.NewSchoolService(repo, s.logger)
	ok, err := schoolService.IsSchoolTeacher(context.Background(), schools[0].ID, teachers[0].ID)
	if err != nil {
		t.Errorf("failed to find school teacher: %v", err)
	}
	t.Assert().Equal(ok, true)

	ok, err = repo.IsSchoolTeacher(context.Background(), schools[0].ID, newTeacherID)
	if err != nil {
		t.Errorf("failed to find school teacher: %v", err)
	}
	t.Assert().Equal(ok, false)
}

func (s *SchoolSuite) TestSchoolService_Create(t provider.T) {
	repo := repository.NewSchoolRepo(s.db)
	schoolService := service.NewSchoolService(repo, s.logger)
	school, err := schoolService.CreateUserSchool(context.Background(), users[0].ID, createdSchool)
	if err != nil {
		t.Errorf("failed to create school: %v", err)
	}
	t.Assert().Equal(school.Name, createdSchool.Name)
}

func (s *SchoolSuite) TestSchoolService_Update(t provider.T) {
	repo := repository.NewSchoolRepo(s.db)
	schoolService := service.NewSchoolService(repo, s.logger)
	school, err := schoolService.Update(context.Background(), schools[0].ID, updatedSchool)
	if err != nil {
		t.Errorf("failed to create school: %v", err)
	}
	require.Equal(t, school.Description, updatedSchool.Description.String)
}

func (s *SchoolSuite) TestSchoolService_Delete(t provider.T) {
	repo := repository.NewSchoolRepo(s.db)
	schoolService := service.NewSchoolService(repo, s.logger)
	err := schoolService.Delete(context.Background(), schools[0].ID)
	if err != nil {
		t.Errorf("failed to delete school: %v", err)
	}
}

func TestSchoolSuite(t *testing.T) {
	suite.RunNamedSuite(t, "School service suite", new(SchoolSuite))
}
