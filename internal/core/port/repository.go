package port

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/domain/dto"
)

//go:generate mockery --dir . --name UserRepository --output ./mocks --filename user.go
type IUserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID int64) (domain.User, error)
	FindByCredentials(ctx context.Context, email string, password string) (domain.User, error)
	FindUserInfo(ctx context.Context, userID int64) (dto.UserInfo, error)
	Create(ctx context.Context, userDTO dto.CreateUserDTO) (domain.User, error)
	Update(ctx context.Context, userID int64, userDTO dto.UpdateUserDTO) (domain.User, error)
	Delete(ctx context.Context, userID int64) error
}

//go:generate mockery --dir . --name CourseRepository --output ./mocks --filename course.go
type ICourseRepository interface {
	FindAll(ctx context.Context) ([]domain.Course, error)
	FindByID(ctx context.Context, courseID int64) (domain.Course, error)
	FindCourseInfo(ctx context.Context, courseID int64) (dto.CourseInfo, error)
	FindStudentCourses(ctx context.Context, studentID int64) ([]domain.Course, error)
	FindTeacherCourses(ctx context.Context, teacherID int64) ([]domain.Course, error)
	AddCourseStudent(ctx context.Context, studentID, courseID int64) error
	AddCourseTeacher(ctx context.Context, teacherID, courseID int64) error
	AddCourseLesson(ctx context.Context, courseID, lessonID int64) error
	DeleteCourseLesson(ctx context.Context, courseID, lessonID int64) error
	Create(ctx context.Context, courseDTO dto.CreateCourseDTO) (domain.Course, error)
	Update(ctx context.Context, courseID int64,
		courseDTO dto.UpdateCourseDTO) (domain.Course, error)
	Delete(ctx context.Context, courseID int64) error
}

//go:generate mockery --dir . --name LessonRepository --output ./mocks --filename lesson.go
type ILessonRepository interface {
	FindAll(ctx context.Context) ([]domain.Lesson, error)
	FindByID(ctx context.Context, lessonID int64) (domain.Lesson, error)
	FindCourseLessons(ctx context.Context, courseID int64) ([]domain.Lesson, error)
	Create(ctx context.Context, lessonDTO dto.CreateLessonDTO) (domain.Lesson, error)
	AddLessonTests(ctx context.Context, lessonID int64, tests []dto.CreateTestDTO) error
	DeleteLessonTest(ctx context.Context, lessonID, testID int64) error
	UpdateLessonTest(ctx context.Context, lessonID, testID int64,
		testDTO dto.UpdateTestDTO) (domain.Test, error)
	UpdateLessonTheory(ctx context.Context, lessonID int64, theoryDTO dto.UpdateTheoryDTO) error
	UpdateLessonVideo(ctx context.Context, lessonID int64, videoDTO dto.UpdateVideoDTO) error
	Delete(ctx context.Context, lessonID int64) error
}

//go:generate mockery --dir . --name SchoolRepository --output ./mocks --filename school.go
type ISchoolRepository interface {
	FindAll(ctx context.Context) ([]domain.School, error)
	FindByID(ctx context.Context, schoolID int64) (domain.School, error)
	AddSchoolTeacher(ctx context.Context, schoolID int64, teacherID int64) error
	CreateUserSchool(ctx context.Context, schoolDTO dto.CreateSchoolDTO,
		userID int64) (domain.School, error)
	UpdateUserSchool(ctx context.Context, schoolID int64,
		schoolDTO dto.UpdateSchoolDTO, userID int64) (domain.School, error)
	Delete(ctx context.Context, schoolID int64) error
}

//go:generate mockery --dir . --name ReviewRepository --output ./mocks --filename review.go
type IReviewRepository interface {
	FindAll(ctx context.Context) ([]domain.Review, error)
	FindByID(ctx context.Context, reviewID int64) (domain.Review, error)
	FindUserReviews(ctx context.Context, userID int64) ([]domain.Review, error)
	FindCourseReviews(ctx context.Context, courseID int64) ([]domain.Review, error)
	CreateCourseReview(ctx context.Context, courseID, userID int64,
		reviewDTO dto.CreateReviewDTO) (domain.Review, error)
	Delete(ctx context.Context, reviewID int64) error
}

//go:generate mockery --dir . --name CertificateRepository --output ./mocks --filename certificate.go
type ICertificateRepository interface {
	FindAll(ctx context.Context) ([]domain.Certificate, error)
	FindByID(ctx context.Context, certificateID int64) (domain.Certificate, error)
	FindUserCertificates(ctx context.Context, userID int64) ([]domain.Certificate, error)
	CreateCourseCertificate(ctx context.Context, userID, courseID int64) (domain.Certificate, error)
}

type IStatisticsRepository interface {
}
