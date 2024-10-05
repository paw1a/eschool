package integration

import (
	"context"
	"github.com/guregu/null"
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
	"reflect"
	"testing"
)

var lessonCourseID = domain.ID("30e18bc1-4354-4937-9a4d-03cf0b7027ca")
var lessons = []domain.Lesson{
	domain.Lesson{
		ID:        domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022ca"),
		CourseID:  lessonCourseID,
		Title:     "lesson1",
		Score:     10,
		Type:      domain.TheoryLesson,
		TheoryUrl: null.StringFrom("url"),
		VideoUrl:  null.String{},
		Tests:     nil,
	},
	domain.Lesson{
		ID:        domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cb"),
		CourseID:  lessonCourseID,
		Title:     "lesson2",
		Score:     10,
		Type:      domain.VideoLesson,
		TheoryUrl: null.String{},
		VideoUrl:  null.StringFrom("url"),
		Tests:     nil,
	},
	domain.Lesson{
		ID:        domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cc"),
		CourseID:  lessonCourseID,
		Title:     "lesson3",
		Score:     10,
		Type:      domain.PracticeLesson,
		TheoryUrl: null.String{},
		VideoUrl:  null.String{},
		Tests:     tests,
	},
}

var tests = []domain.Test{
	domain.Test{
		ID:       domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7027ca"),
		LessonID: domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cc"),
		TaskUrl:  "url",
		Options:  []string{"opt1", "opt2", "opt3"},
		Answer:   "opt1",
		Level:    2,
		Score:    12,
	},
	domain.Test{
		ID:       domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7027cb"),
		LessonID: domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cc"),
		TaskUrl:  "url",
		Options:  []string{"opt1", "opt2"},
		Answer:   "opt2",
		Level:    2,
		Score:    12,
	},
	domain.Test{
		ID:       domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7027cc"),
		LessonID: domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cc"),
		TaskUrl:  "url",
		Options:  []string{"opt1"},
		Answer:   "opt1",
		Level:    2,
		Score:    12,
	},
}

var createdLesson = port.CreateTheoryParam{
	Title:  "created lesson 4",
	Score:  100,
	Theory: "theory text",
}

var updatedLesson = domain.Lesson{
	ID:        domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cc"),
	CourseID:  lessonCourseID,
	Title:     "updated lesson 3",
	Score:     20,
	Type:      domain.PracticeLesson,
	TheoryUrl: null.String{},
	VideoUrl:  null.String{},
	Tests:     tests,
}

var createdTests = []domain.Test{
	domain.Test{
		ID:       domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7025ca"),
		LessonID: domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cd"),
		TaskUrl:  "url",
		Options:  []string{"opt1", "opt2", "opt3"},
		Answer:   "opt1",
		Level:    2,
		Score:    12,
	},
	domain.Test{
		ID:       domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7025cb"),
		LessonID: domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cd"),
		TaskUrl:  "url",
		Options:  []string{"opt1", "opt2"},
		Answer:   "opt2",
		Level:    2,
		Score:    12,
	},
	domain.Test{
		ID:       domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7025cc"),
		LessonID: domain.ID("30e18bc1-4352-4937-9a3b-03cf0b7022cd"),
		TaskUrl:  "url",
		Options:  []string{"opt1"},
		Answer:   "opt1",
		Level:    2,
		Score:    12,
	},
}

type LessonSuite struct {
	suite.Suite
	logger         *zap.Logger
	container      *postgres.PostgresContainer
	minioContainer *minio.MinioContainer
	db             *sqlx.DB
	minioClient    *minio2.Client
}

func (s *LessonSuite) BeforeAll(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

func (s *LessonSuite) BeforeEach(t provider.T) {
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

func (s *LessonSuite) AfterAll(t provider.T) {
	if err := s.container.Terminate(context.Background()); err != nil {
		t.Fatalf("failed to terminate container: %s", err)
	}
}

func (s *LessonSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *LessonSuite) TestUserService_FindAll(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewLessonRepo(s.db)
	objectStorage := storage.NewObjectStorage(s.minioClient, &minioConfig)
	lessonService := service.NewLessonService(repo, objectStorage, s.logger)
	found, err := lessonService.FindAll(context.Background())
	if err != nil {
		t.Errorf("failed to find all lessons: %v", err)
	}
	t.Assert().Equal(len(found), len(lessons))
	for i, lesson := range found {
		if lesson.Type == domain.PracticeLesson {
			for j, test := range lesson.Tests {
				t.Assert().Equal(test, lessons[i].Tests[j])
				t.Assert().Equal(reflect.DeepEqual(test.Options,
					lessons[i].Tests[j].Options), true)
			}
		} else {
			t.Assert().Equal(lesson, lessons[i])
		}
	}
}

func (s *LessonSuite) TestUserService_FindByID(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewLessonRepo(s.db)
	objectStorage := storage.NewObjectStorage(s.minioClient, &minioConfig)
	lessonService := service.NewLessonService(repo, objectStorage, s.logger)
	lesson, err := lessonService.FindByID(context.Background(), lessons[2].ID)
	if err != nil {
		t.Errorf("failed to find course with id: %v", err)
	}

	for j, test := range lesson.Tests {
		t.Assert().Equal(test, lesson.Tests[j])
		t.Assert().Equal(reflect.DeepEqual(test.Options,
			lesson.Tests[j].Options), true)
	}
}

func (s *LessonSuite) TestUserService_FindCourseLessons(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewLessonRepo(s.db)
	objectStorage := storage.NewObjectStorage(s.minioClient, &minioConfig)
	lessonService := service.NewLessonService(repo, objectStorage, s.logger)
	found, err := lessonService.FindCourseLessons(context.Background(), lessonCourseID)
	if err != nil {
		t.Errorf("failed to find course lessons: %v", err)
	}
	t.Assert().Equal(len(found), len(lessons))
	for i, lesson := range found {
		if lesson.Type == domain.PracticeLesson {
			for j, test := range lesson.Tests {
				t.Assert().Equal(test, lessons[i].Tests[j])
				t.Assert().Equal(reflect.DeepEqual(test.Options,
					lessons[i].Tests[j].Options), true)
			}
		} else {
			t.Assert().Equal(lesson, lessons[i])
		}
	}
}

func (s *LessonSuite) TestUserService_CreateTheoryLesson(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewLessonRepo(s.db)
	objectStorage := storage.NewObjectStorage(s.minioClient, &minioConfig)
	lessonService := service.NewLessonService(repo, objectStorage, s.logger)
	lesson, err := lessonService.CreateTheoryLesson(context.Background(), courseID, createdLesson)
	if err != nil {
		t.Errorf("failed to create lesson: %v", err)
	}
	t.Assert().Equal(lesson, createdLesson)
}

func (s *LessonSuite) TestUserService_Delete(t provider.T) {
	if isPreviousTestsFailed() {
		t.Skip()
	}
	repo := repository.NewLessonRepo(s.db)
	objectStorage := storage.NewObjectStorage(s.minioClient, &minioConfig)
	lessonService := service.NewLessonService(repo, objectStorage, s.logger)
	err := lessonService.Delete(context.Background(), lessons[0].ID)
	if err != nil {
		t.Errorf("failed to delete lesson: %v", err)
	}
}

func TestLessonSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Lesson service suite", new(UserSuite))
}
