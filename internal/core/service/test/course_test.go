package service

import (
	"context"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/service"
	repositoryMocks "github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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
		ID:        domain.NewID(),
		CourseID:  courses[0].ID,
		Title:     "Lesson1",
		Score:     5,
		Type:      domain.TheoryLesson,
		TheoryUrl: null.StringFrom("ddd"),
		VideoUrl:  null.String{},
	},
	domain.Lesson{
		ID:        domain.NewID(),
		CourseID:  courses[0].ID,
		Title:     "Lesson2",
		Score:     4,
		Type:      domain.PracticeLesson,
		TheoryUrl: null.String{},
		VideoUrl:  null.String{},
	},
	domain.Lesson{
		ID:        domain.NewID(),
		CourseID:  courses[0].ID,
		Title:     "Lesson3",
		Score:     -1,
		Type:      domain.TheoryLesson,
		TheoryUrl: null.StringFrom("ddd"),
		VideoUrl:  null.String{},
	},
	domain.Lesson{
		ID:        domain.NewID(),
		CourseID:  courses[0].ID,
		Title:     "Lesson4",
		Score:     5,
		Type:      domain.VideoLesson,
		TheoryUrl: null.String{},
		VideoUrl:  null.StringFrom("ddd"),
	},
}

var tests = []domain.Test{
	domain.Test{
		ID:       domain.NewID(),
		LessonID: lessons[1].ID,
		TaskUrl:  "url",
		Options:  []string{"opt1", "opt2", "opt3"},
		Answer:   "answer",
		Level:    3,
		Score:    10,
	},
	domain.Test{
		ID:       domain.NewID(),
		LessonID: lessons[1].ID,
		TaskUrl:  "",
		Options:  []string{"opt1"},
		Answer:   "answer",
		Level:    10,
		Score:    5,
	},
	domain.Test{
		ID:       domain.NewID(),
		LessonID: lessons[1].ID,
		TaskUrl:  "url",
		Options:  nil,
		Answer:   "answer",
		Level:    10,
		Score:    5,
	},
}

func TestCourseService_ConfirmDraftCourse(t *testing.T) {
	testTable := []struct {
		name           string
		initCourseRepo func(userRepo *repositoryMocks.CourseRepository)
		initLessonRepo func(userRepo *repositoryMocks.LessonRepository)
		course         domain.Course
		hasError       bool
	}{
		{
			name:   "course is correct, ok",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
				courseRepo.On("UpdateStatus", context.Background(), courses[0].ID, domain.CourseReady).
					Return(nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessons[1].Tests = []domain.Test{tests[0]}
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
				lessons[1].Tests = nil
			},
			hasError: false,
		},
		{
			name:   "course is already ready, error",
			course: courses[1],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[1].ID).
					Return(courses[1], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
			},
			hasError: true,
		},
		{
			name:   "course has only theory, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[0]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course has only practice, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[1], lessons[1]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course invalid lesson mark, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[1], lessons[2]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course lesson without test, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course video lesson without url, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[3], lessons[1]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course lesson test question in empty, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course lesson test options in empty, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
			},
			hasError: true,
		},
		{
			name:   "course cannot be updated, error",
			course: courses[0],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			initLessonRepo: func(lessonRepo *repositoryMocks.LessonRepository) {
				lessonRepo.On("FindCourseLessons", context.Background(), courses[0].ID).
					Return([]domain.Lesson{lessons[0], lessons[1]}, nil)
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			courseRepo := repositoryMocks.NewCourseRepository(t)
			lessonRepo := repositoryMocks.NewLessonRepository(t)
			schoolRepo := repositoryMocks.NewSchoolRepository(t)
			statRepo := repositoryMocks.NewStatRepository(t)
			courseService := service.NewCourseService(courseRepo, lessonRepo, schoolRepo, statRepo, zap.NewNop())

			test.initCourseRepo(courseRepo)
			test.initLessonRepo(lessonRepo)

			errList := courseService.ConfirmDraftCourse(context.Background(), test.course.ID)

			if test.hasError {
				require.NotEqual(t, len(errList), 0)
			} else {
				require.Equal(t, len(errList), 0)
			}
		})
	}
}

func TestCourseService_PublishReadyCourse(t *testing.T) {
	testTable := []struct {
		name           string
		initCourseRepo func(userRepo *repositoryMocks.CourseRepository)
		course         domain.Course
		hasError       bool
	}{
		{
			name:   "course is ready to be published, ok",
			course: courses[1],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
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
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[0].ID).
					Return(courses[0], nil)
			},
			hasError: true,
		},
		{
			name:   "course is already published, error",
			course: courses[2],
			initCourseRepo: func(courseRepo *repositoryMocks.CourseRepository) {
				courseRepo.On("FindByID", context.Background(), courses[2].ID).
					Return(courses[2], nil)
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			courseRepo := repositoryMocks.NewCourseRepository(t)
			lessonRepo := repositoryMocks.NewLessonRepository(t)
			schoolRepo := repositoryMocks.NewSchoolRepository(t)
			statRepo := repositoryMocks.NewStatRepository(t)
			courseService := service.NewCourseService(courseRepo, lessonRepo, schoolRepo, statRepo, zap.NewNop())

			test.initCourseRepo(courseRepo)
			err := courseService.PublishReadyCourse(context.Background(), test.course.ID)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, err, nil)
			}
		})
	}
}
