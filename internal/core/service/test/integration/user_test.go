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
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/zap"
	"testing"
)

var users = []domain.User{
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
	domain.User{
		ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc"),
		Name:      "Emir",
		Surname:   "Shimshir",
		Phone:     null.StringFrom("+79992233555"),
		City:      null.String{},
		AvatarUrl: null.String{},
		Email:     "emir@gmail.com",
		Password:  "12345",
	},
}

var createdUser = port.CreateUserParam{
	Name:      "createdName",
	Surname:   "createdSurname",
	Phone:     null.StringFrom("+77777777777"),
	City:      null.StringFrom("Test city"),
	AvatarUrl: null.String{},
	Email:     "user@mail.com",
	Password:  "password",
}

var updatedUser = port.UpdateUserParam{
	Name:      null.StringFrom("Maxim"),
	Surname:   null.StringFrom("Shpakovskiy"),
	Phone:     null.String{},
	City:      null.StringFrom("Sochi"),
	AvatarUrl: null.String{},
}

type UserSuite struct {
	suite.Suite
	logger    *zap.Logger
	container *postgres.PostgresContainer
	db        *sqlx.DB
}

func (s *UserSuite) BeforeAll(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

func (s *UserSuite) BeforeEach(t provider.T) {
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

func (s *UserSuite) AfterAll(t provider.T) {
	if err := s.container.Terminate(context.Background()); err != nil {
		t.Fatalf("failed to terminate container: %s", err)
	}
}

func (s *UserSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *UserSuite) TestUserService_FindAll(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	t.Title("User service find all")
	repo := repository.NewUserRepo(s.db)
	userService := service.NewUserService(repo, s.logger)
	found, err := userService.FindAll(context.Background())
	if err != nil {
		t.Errorf("failed to find all users: %v", err)
	}
	t.Assert().Equal(len(found), len(users))
	for i := range users {
		t.Assert().Equal(users[i], found[i])
	}
}

func (s *UserSuite) TestUserService_FindByID(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	t.Title("User service find by id")
	repo := repository.NewUserRepo(s.db)
	userService := service.NewUserService(repo, s.logger)
	user, err := userService.FindByID(context.Background(), users[0].ID)
	if err != nil {
		t.Errorf("failed to find user with id: %v", err)
	}
	t.Assert().Equal(user, users[0])
}

func (s *UserSuite) TestUserService_Create(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewUserRepo(s.db)
	userService := service.NewUserService(repo, s.logger)
	user, err := userService.Create(context.Background(), createdUser)
	if err != nil {
		t.Errorf("failed to create user: %v", err)
	}
	t.Assert().Equal(user.Name, createdUser.Name)
	t.Assert().Equal(user.Email, createdUser.Email)
}

func (s *UserSuite) TestUserService_Update(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewUserRepo(s.db)
	userService := service.NewUserService(repo, s.logger)
	user, err := userService.Update(context.Background(), users[0].ID, updatedUser)
	if err != nil {
		t.Errorf("failed to create user: %v", err)
	}
	t.Assert().Equal(users[0].Email, user.Email)
	t.Assert().Equal(updatedUser.Name.String, user.Name)
}

func (s *UserSuite) TestUserService_Delete(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewUserRepo(s.db)
	userService := service.NewUserService(repo, s.logger)
	err := userService.Delete(context.Background(), users[0].ID)
	if err != nil {
		t.Errorf("failed to delete user: %v", err)
	}
}

func TestUserSuite(t *testing.T) {
	suite.RunNamedSuite(t, "User service suite", new(UserSuite))
}
