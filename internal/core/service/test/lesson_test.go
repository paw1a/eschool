package service

import (
	"context"
	"github.com/paw1a/eschool/internal/adapter/repository/mocks"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var lessonID = domain.NewID()
var testParams = []port.CreateTestParam{
	{
		QuestionString: "question",
		Options:        []string{"opt1", "opt2", "opt3"},
		Answer:         "answer",
		Level:          3,
		Score:          10,
	},
	{
		QuestionString: "",
		Options:        []string{"opt1"},
		Answer:         "answer",
		Level:          10,
		Score:          5,
	},
	{
		QuestionString: "question",
		Options:        nil,
		Answer:         "answer",
		Level:          10,
		Score:          5,
	},
	{
		QuestionString: "question",
		Options:        []string{"opt1", "opt2", "opt3"},
		Answer:         "answer",
		Level:          -1,
		Score:          10,
	},
	{
		QuestionString: "question",
		Options:        []string{"opt1", "opt2", "opt3"},
		Answer:         "answer",
		Level:          2,
		Score:          -1,
	},
}

func TestLessonService_AddLessonTests(t *testing.T) {
	testTable := []struct {
		name           string
		initLessonRepo func(userRepo *mocks.LessonRepository)
		tests          []port.CreateTestParam
		hasError       bool
	}{
		{
			name:  "tests are correct, ok",
			tests: []port.CreateTestParam{testParams[0]},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("AddLessonTests", context.Background(), mock.AnythingOfType("[]domain.Test")).
					Return(nil)
			},
			hasError: false,
		},
		{
			name:  "tests question string is empty, error",
			tests: []port.CreateTestParam{testParams[1]},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
			},
			hasError: true,
		},
		{
			name:  "tests options is empty, error",
			tests: []port.CreateTestParam{testParams[2]},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
			},
			hasError: true,
		},
		{
			name:  "tests level is invalid, error",
			tests: []port.CreateTestParam{testParams[3]},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
			},
			hasError: true,
		},
		{
			name:  "tests mark is invalid, error",
			tests: []port.CreateTestParam{testParams[4]},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Logf("Test: %s", test.name)

		lessonRepo := mocks.NewLessonRepository(t)
		lessonService := service.NewLessonService(lessonRepo)
		test.initLessonRepo(lessonRepo)
		err := lessonService.AddLessonTests(context.Background(), lessonID, test.tests)

		if test.hasError {
			require.Error(t, err)
		} else {
			require.Equal(t, err, nil)
		}
	}
}
