package e2e

import (
	"context"
	"github.com/jmoiron/sqlx"
	minio2 "github.com/minio/minio-go/v7"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	storage "github.com/paw1a/eschool/internal/adapter/storage/minio"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/testcontainers/testcontainers-go/modules/minio"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/zap"
	"testing"
)

var userID = domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027ca")

var createSchoolParam = port.CreateSchoolParam{
	Name:        "e2e school",
	Description: "e2e school description",
}

var createCourseParam = port.CreateCourseParam{
	Name:     "e2e course",
	Level:    4,
	Price:    1200,
	Language: "russian",
}

var createTheoryLessonParam = port.CreateTheoryParam{
	Title:  "e2e theory lesson",
	Score:  20,
	Theory: "# Lesson theory markdown",
}

var createPracticeLessonParam = port.CreatePracticeParam{
	Title: "e2e practice param",
	Score: 50,
	Tests: []port.CreateTestParam{
		port.CreateTestParam{
			Task:    "1 + 1 = ?",
			Options: []string{"1", "2", "3"},
			Answer:  "3",
			Level:   2,
			Score:   20,
		},
		port.CreateTestParam{
			Task:    "4 * 4 = ?",
			Options: []string{"4", "10", "16"},
			Answer:  "16",
			Level:   3,
			Score:   40,
		},
	},
}

type EndToEndSuite struct {
	suite.Suite
	logger         *zap.Logger
	container      *postgres.PostgresContainer
	minioContainer *minio.MinioContainer
	db             *sqlx.DB
	minioClient    *minio2.Client
}

func (s *EndToEndSuite) BeforeAll(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

func (s *EndToEndSuite) BeforeEach(t provider.T) {
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

	s.minioContainer, err = newMinioContainer(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	url, err = s.minioContainer.ConnectionString(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	s.minioClient, err = newMinioClient(url)
	if err != nil {
		t.Fatal(err)
	}
}

func (s *EndToEndSuite) AfterAll(t provider.T) {
	if err := s.container.Terminate(context.Background()); err != nil {
		t.Fatalf("failed to terminate container: %s", err)
	}
}

func (s *EndToEndSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *EndToEndSuite) Test_EndToEnd(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	lessonRepo := repository.NewLessonRepo(s.db)
	schoolRepo := repository.NewSchoolRepo(s.db)
	courseRepo := repository.NewCourseRepo(s.db)
	statRepo := repository.NewStatRepo(s.db)
	objectStorage := storage.NewObjectStorage(s.minioClient, &minioConfig)
	lessonService := service.NewLessonService(lessonRepo, objectStorage, s.logger)
	schoolService := service.NewSchoolService(schoolRepo, s.logger)
	courseService := service.NewCourseService(courseRepo, lessonRepo, schoolRepo, statRepo, s.logger)

	school, err := schoolService.CreateUserSchool(context.Background(), userID, createSchoolParam)
	if err != nil {
		t.Errorf("failed to create school: %v", err)
	}
	t.Assert().Equal(school.Name, createSchoolParam.Name)

	err = schoolService.AddSchoolTeacher(context.Background(), school.ID, userID)
	if err != nil {
		t.Errorf("failed to add school teacher: %v", err)
	}

	course, err := courseService.CreateSchoolCourse(context.Background(), school.ID, createCourseParam)
	if err != nil {
		t.Errorf("failed create school course: %v", err)
	}
	t.Assert().Equal(course.Name, createCourseParam.Name)

	err = courseService.AddCourseTeacher(context.Background(), userID, course.ID)
	if err != nil {
		t.Errorf("failed to add course teacher: %v", err)
	}

	lessonTheory, err := lessonService.CreateTheoryLesson(context.Background(), course.ID, createTheoryLessonParam)
	if err != nil {
		t.Errorf("failed to create theory lesson: %v", err)
	}
	t.Assert().Equal(lessonTheory.Title, createTheoryLessonParam.Title)
	t.Assert().Equal(lessonTheory.Score, createTheoryLessonParam.Score)

	lessonPractice, err := lessonService.CreatePracticeLesson(context.Background(), course.ID, createPracticeLessonParam)
	if err != nil {
		t.Errorf("failed to create theory lesson: %v", err)
	}
	t.Assert().Equal(lessonPractice.Title, createPracticeLessonParam.Title)
	t.Assert().Equal(lessonPractice.Score, createPracticeLessonParam.Score)
	for i, test := range lessonPractice.Tests {
		t.Assert().Equal(test.Answer, createPracticeLessonParam.Tests[i].Answer)
	}

	errs := courseService.ConfirmDraftCourse(context.Background(), course.ID)
	t.Assert().Nil(errs)

	err = courseService.PublishReadyCourse(context.Background(), course.ID)
	t.Assert().Nil(err)

	course, err = courseService.FindByID(context.Background(), course.ID)
	if err != nil {
		t.Errorf("failed to find course by id: %v", err)
	}
	t.Assert().Equal(course.Status, domain.CoursePublished)
}

func TestEndToEndSuite(t *testing.T) {
	suite.RunNamedSuite(t, "End to end suite", new(EndToEndSuite))
}
