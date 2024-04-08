package service

import (
	"context"
	"errors"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/adapter/repository/mocks"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/stretchr/testify/require"
	"testing"
)

var courses = []domain.Course{
	domain.Course{
		ID:       domain.NewID(),
		SchoolID: domain.ID(""),
		Name:     "Course1",
		Level:    1,
		Price:    1200,
		Language: "russian",
		Status:   domain.CourseDraft,
	},
	domain.Course{
		ID:       domain.NewID(),
		SchoolID: domain.ID(""),
		Name:     "Course2",
		Level:    3,
		Price:    200,
		Language: "english",
		Status:   domain.CourseReady,
	},
	domain.Course{
		ID:       domain.NewID(),
		SchoolID: domain.ID(""),
		Name:     "Course3",
		Level:    6,
		Price:    2000,
		Language: "russian",
		Status:   domain.CoursePublished,
	},
}

var lessons = []domain.Lesson{
	domain.Lesson{
		ID:         domain.NewID(),
		CourseID:   courses[0].ID,
		Title:      "Lesson1",
		Mark:       5,
		Type:       domain.TheoryLesson,
		ContentUrl: null.String{},
	},
	domain.Lesson{
		ID:         domain.NewID(),
		CourseID:   courses[0].ID,
		Title:      "Lesson2",
		Mark:       4,
		Type:       domain.PracticeLesson,
		ContentUrl: null.String{},
	},
	domain.Lesson{
		ID:         domain.NewID(),
		CourseID:   courses[0].ID,
		Title:      "Lesson3",
		Mark:       -1,
		Type:       domain.TheoryLesson,
		ContentUrl: null.String{},
	},
	domain.Lesson{
		ID:         domain.NewID(),
		CourseID:   courses[0].ID,
		Title:      "Lesson4",
		Mark:       5,
		Type:       domain.VideoLesson,
		ContentUrl: null.String{},
	},
}

var tests = []domain.Test{
	domain.Test{
		ID:          domain.NewID(),
		LessonID:    lessons[1].ID,
		QuestionUrl: "url",
		Options:     []string{"opt1", "opt2", "opt3"},
		Answer:      "answer",
		Level:       3,
		Mark:        10,
	},
	domain.Test{
		ID:          domain.NewID(),
		LessonID:    lessons[1].ID,
		QuestionUrl: "",
		Options:     []string{"opt1"},
		Answer:      "answer",
		Level:       10,
		Mark:        5,
	},
	domain.Test{
		ID:          domain.NewID(),
		LessonID:    lessons[1].ID,
		QuestionUrl: "url",
		Options:     nil,
		Answer:      "answer",
		Level:       10,
		Mark:        5,
	},
}

func TestCourseService_ConfirmDraftCourse(t *testing.T) {
	testTable := []struct {
		name           string
		initCourseRepo func(userRepo *mocks.CourseRepository)
		initLessonRepo func(userRepo *mocks.LessonRepository)
		course         domain.Course
		hasError       bool
	}{
		{
			name:   "course is correct, ok",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
				courseRepo.On("UpdateStatus", context.Background(), courses[0].ID, domain.CourseReady).
					Return(nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[0]}, nil)
			},
			hasError: false,
		},
		{
			name:   "course is already ready, error",
			course: courses[1],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[1].ID).
					Return(courses[1], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
			},
			hasError: true,
		},
		{
			name:   "course has only theory, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[0]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course has only practice, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[1], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[0]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course invalid lesson mark, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[1], lessons[2]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[0]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course lesson without test, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{}, nil)
			},
			hasError: true,
		},
		{
			name:   "course video lesson without url, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[3], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[0]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course lesson test question in empty, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[1]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course lesson test options in empty, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[2]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course cannot be updated, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
				courseRepo.On("UpdateStatus", context.Background(), courses[0].ID, domain.CourseReady).
					Return(errors.New("error"))
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[0]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course cannot be updated, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
				courseRepo.On("UpdateStatus", context.Background(), courses[0].ID, domain.CourseReady).
					Return(errors.New("error"))
			},
			initLessonRepo: func(lessonRepo *mocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
				lessonRepo.On("FindLessonTests", context.Background(), lessons[1].ID).
					Return([]domain.Test{tests[0]}, nil)
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Logf("Test: %s", test.name)

		courseRepo := mocks.NewCourseRepository(t)
		lessonRepo := mocks.NewLessonRepository(t)
		courseService := service.NewCourseService(courseRepo, lessonRepo)

		test.initCourseRepo(courseRepo)
		test.initLessonRepo(lessonRepo)

		errList := courseService.ConfirmDraftCourse(context.Background(), test.course.ID)

		if test.hasError {
			require.NotEqual(t, len(errList), 0)
		} else {
			require.Equal(t, len(errList), 0)
		}
	}
}

func TestCourseService_PublishReadyCourse(t *testing.T) {
	testTable := []struct {
		name           string
		initCourseRepo func(userRepo *mocks.CourseRepository)
		course         domain.Course
		hasError       bool
	}{
		{
			name:   "course is ready to be published, ok",
			course: courses[1],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[1].ID).
					Return(courses[1], nil)
				courseRepo.On("UpdateStatus", context.Background(), courses[1].ID, domain.CoursePublished).
					Return(nil)
			},
			hasError: false,
		},
		{
			name:   "course draft state, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			hasError: true,
		},
		{
			name:   "course is already published, error",
			course: courses[2],
			initCourseRepo: func(courseRepo *mocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[2].ID).
					Return(courses[2], nil)
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Logf("Test: %s", test.name)

		courseRepo := mocks.NewCourseRepository(t)
		lessonRepo := mocks.NewLessonRepository(t)
		courseService := service.NewCourseService(courseRepo, lessonRepo)

		test.initCourseRepo(courseRepo)
		err := courseService.PublishReadyCourse(context.Background(), test.course.ID)

		if test.hasError {
			require.Error(t, err)
		} else {
			require.Equal(t, err, nil)
		}
	}
}
