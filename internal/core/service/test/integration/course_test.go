package integration

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/zap"
	"testing"
)

var courses = []domain.Course{
	domain.Course{
		ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027ca"),
		SchoolID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
		Name:     "course1",
		Level:    4,
		Price:    1200,
		Language: "russian",
		Status:   domain.CourseDraft,
	},
	domain.Course{
		ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027cb"),
		SchoolID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
		Name:     "course2",
		Level:    2,
		Price:    1500,
		Language: "english",
		Status:   domain.CoursePublished,
	},
	domain.Course{
		ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027cc"),
		SchoolID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
		Name:     "course3",
		Level:    3,
		Price:    12000,
		Language: "russian",
		Status:   domain.CourseReady,
	},
	domain.Course{
		ID:       domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027cd"),
		SchoolID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
		Name:     "course4",
		Level:    2,
		Price:    0,
		Language: "english",
		Status:   domain.CoursePublished,
	},
}

var studentCoursesID = domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027ca")
var teacherCoursesID = domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cb")
var newUserID = domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc")

type CourseSuite struct {
	suite.Suite
	logger    *zap.Logger
	container *postgres.PostgresContainer
	db        *sqlx.DB
}

func (s *CourseSuite) BeforeAll(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

func (s *CourseSuite) BeforeEach(t provider.T) {
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

func (s *CourseSuite) AfterAll(t provider.T) {
	if err := s.container.Terminate(context.Background()); err != nil {
		t.Fatalf("failed to terminate container: %s", err)
	}
}

func (s *CourseSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *CourseSuite) TestUserService_FindAll(t provider.T) {
	repo := repository.NewCourseRepo(s.db)
	lessonRepo := repository.NewLessonRepo(s.db)
	schoolRepo := repository.NewSchoolRepo(s.db)
	statRepo := repository.NewStatRepo(s.db)
	courseService := service.NewCourseService(repo, lessonRepo, schoolRepo, statRepo, s.logger)
	found, err := courseService.FindAll(context.Background())
	if err != nil {
		t.Errorf("failed to find all courses: %v", err)
	}
	t.Assert().Equal(len(found), len(courses))
}

func (s *CourseSuite) TestUserService_FindByID(t provider.T) {
	repo := repository.NewCourseRepo(s.db)
	lessonRepo := repository.NewLessonRepo(s.db)
	schoolRepo := repository.NewSchoolRepo(s.db)
	statRepo := repository.NewStatRepo(s.db)
	courseService := service.NewCourseService(repo, lessonRepo, schoolRepo, statRepo, s.logger)
	course, err := courseService.FindByID(context.Background(), courses[0].ID)
	if err != nil {
		t.Errorf("failed to find course with id: %v", err)
	}
	t.Assert().Equal(course, courses[0])
}

func (s *CourseSuite) TestUserService_FindStudentCourses(t provider.T) {
	repo := repository.NewCourseRepo(s.db)
	lessonRepo := repository.NewLessonRepo(s.db)
	schoolRepo := repository.NewSchoolRepo(s.db)
	statRepo := repository.NewStatRepo(s.db)
	courseService := service.NewCourseService(repo, lessonRepo, schoolRepo, statRepo, s.logger)
	found, err := courseService.FindStudentCourses(context.Background(), studentCoursesID)
	if err != nil {
		t.Errorf("failed to find student courses: %v", err)
	}
	t.Assert().Equal(len(found), 2)
	t.Assert().Equal(found[0], courses[0])
	t.Assert().Equal(found[1], courses[1])
}

func (s *CourseSuite) TestUserService_FindTeacherCourses(t provider.T) {
	repo := repository.NewCourseRepo(s.db)
	lessonRepo := repository.NewLessonRepo(s.db)
	schoolRepo := repository.NewSchoolRepo(s.db)
	statRepo := repository.NewStatRepo(s.db)
	courseService := service.NewCourseService(repo, lessonRepo, schoolRepo, statRepo, s.logger)
	found, err := courseService.FindTeacherCourses(context.Background(), teacherCoursesID)
	if err != nil {
		t.Errorf("failed to find teacher courses: %v", err)
	}
	t.Assert().Equal(len(found), 2)
	t.Assert().Equal(found[0], courses[0])
	t.Assert().Equal(found[1], courses[1])
}

func (s *CourseSuite) TestUserService_AddCourseStudent(t provider.T) {
	repo := repository.NewCourseRepo(s.db)
	lessonRepo := repository.NewLessonRepo(s.db)
	schoolRepo := repository.NewSchoolRepo(s.db)
	statRepo := repository.NewStatRepo(s.db)
	courseService := service.NewCourseService(repo, lessonRepo, schoolRepo, statRepo, s.logger)
	err := courseService.AddCourseStudent(context.Background(), newUserID, courses[0].ID)
	if err != nil {
		t.Errorf("failed to add course student: %v", err)
	}
}

func (s *CourseSuite) TestUserService_Delete(t provider.T) {
	repo := repository.NewCourseRepo(s.db)
	lessonRepo := repository.NewLessonRepo(s.db)
	schoolRepo := repository.NewSchoolRepo(s.db)
	statRepo := repository.NewStatRepo(s.db)
	courseService := service.NewCourseService(repo, lessonRepo, schoolRepo, statRepo, s.logger)
	err := courseService.Delete(context.Background(), courses[0].ID)
	if err != nil {
		t.Errorf("failed to delete course: %v", err)
	}
}

func TestCourseSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Course service suite", new(CourseSuite))
}
