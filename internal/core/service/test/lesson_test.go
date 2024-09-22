package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	repositoryMocks "github.com/paw1a/eschool/internal/core/service/mocks"
	storageMocks "github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

var courseID = domain.NewID()
var testParams = []port.CreateTestParam{
	{
		Task:    "question",
		Options: []string{"opt1", "opt2", "opt3"},
		Answer:  "answer",
		Level:   3,
		Score:   10,
	},
	{
		Task:    "question",
		Options: nil,
		Answer:  "answer",
		Level:   10,
		Score:   5,
	},
	{
		Task:    "question",
		Options: []string{"opt1", "opt2", "opt3"},
		Answer:  "answer",
		Level:   -1,
		Score:   10,
	},
	{
		Task:    "question",
		Options: []string{"opt1", "opt2", "opt3"},
		Answer:  "answer",
		Level:   2,
		Score:   -1,
	},
}

func TestLessonService_CreatePracticeLesson(t *testing.T) {
	testTable := []struct {
		name           string
		initLessonRepo func(userRepo *repositoryMocks.LessonRepository)
		initStorage    func(storage *storageMocks.ObjectStorage)
		param          port.CreatePracticeParam
		hasError       bool
	}{
		{
			name: "tests are correct, ok",
			param: port.CreatePracticeParam{
				Title: "lesson",
				Score: 120,
				Tests: []port.CreateTestParam{testParams[0]},
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("Create", context.Background(), mock.AnythingOfType("domain.Lesson")).
					Return(domain.Lesson{}, nil)
			},
			initStorage: func(storage *storageMocks.ObjectStorage) {
				storage.On("SaveFile", context.Background(), mock.AnythingOfType("domain.File")).
					Return(domain.Url("url"), nil)
			},
			hasError: false,
		},
		{
			name: "tests options is empty, error",
			param: port.CreatePracticeParam{
				Title: "lesson",
				Score: 120,
				Tests: []port.CreateTestParam{testParams[1]},
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
			},
			initStorage: func(storage *storageMocks.ObjectStorage) {
			},
			hasError: true,
		},
		{
			name: "tests level is invalid, error",
			param: port.CreatePracticeParam{
				Title: "lesson",
				Score: 120,
				Tests: []port.CreateTestParam{testParams[2]},
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
			},
			initStorage: func(storage *storageMocks.ObjectStorage) {

			},
			hasError: true,
		},
		{
			name: "tests mark is invalid, error",
			param: port.CreatePracticeParam{
				Title: "lesson",
				Score: 120,
				Tests: []port.CreateTestParam{testParams[3]},
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
			},
			initStorage: func(storage *storageMocks.ObjectStorage) {

			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			lessonRepo := repositoryMocks.NewLessonRepository(t)
			objectStorage := storageMocks.NewObjectStorage(t)
			lessonService := service.NewLessonService(lessonRepo, objectStorage, zap.NewNop())
			test.initLessonRepo(lessonRepo)
			test.initStorage(objectStorage)
			_, err := lessonService.CreatePracticeLesson(context.Background(), courseID, test.param)
			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, err, nil)
			}
		})
	}
}
