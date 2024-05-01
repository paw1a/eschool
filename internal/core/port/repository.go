package port

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
)

type IUserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID domain.ID) (domain.User, error)
	FindByCredentials(ctx context.Context, email string, password string) (domain.User, error)
	FindUserInfo(ctx context.Context, userID domain.ID) (UserInfo, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, userID domain.ID) error
}

type ICourseRepository interface {
	FindAll(ctx context.Context) ([]domain.Course, error)
	FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error)
	FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error)
	FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error)
	AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error
	AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error
	Create(ctx context.Context, course domain.Course) (domain.Course, error)
	Update(ctx context.Context, course domain.Course) (domain.Course, error)
	UpdateStatus(ctx context.Context, courseID domain.ID, status domain.CourseStatus) error
	Delete(ctx context.Context, courseID domain.ID) error
}

type ILessonRepository interface {
	FindAll(ctx context.Context) ([]domain.Lesson, error)
	FindByID(ctx context.Context, lessonID domain.ID) (domain.Lesson, error)
	FindCourseLessons(ctx context.Context, courseID domain.ID) ([]domain.Lesson, error)
	FindLessonTests(ctx context.Context, lessonID domain.ID) ([]domain.Test, error)
	Create(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error)
	Update(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error)
	Delete(ctx context.Context, lessonID domain.ID) error
}

type ISchoolRepository interface {
	FindAll(ctx context.Context) ([]domain.School, error)
	FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error)
	FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error)
	FindSchoolCourses(ctx context.Context, schoolID domain.ID) ([]domain.Course, error)
	FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error)
	AddSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) error
	Create(ctx context.Context, school domain.School) (domain.School, error)
	Update(ctx context.Context, school domain.School) (domain.School, error)
	Delete(ctx context.Context, schoolID domain.ID) error
}

type IReviewRepository interface {
	FindAll(ctx context.Context) ([]domain.Review, error)
	FindByID(ctx context.Context, reviewID domain.ID) (domain.Review, error)
	FindUserReviews(ctx context.Context, userID domain.ID) ([]domain.Review, error)
	FindCourseReviews(ctx context.Context, courseID domain.ID) ([]domain.Review, error)
	Create(ctx context.Context, review domain.Review) (domain.Review, error)
	Delete(ctx context.Context, reviewID domain.ID) error
}

type ICertificateRepository interface {
	FindAll(ctx context.Context) ([]domain.Certificate, error)
	FindByID(ctx context.Context, certID domain.ID) (domain.Certificate, error)
	FindUserCertificates(ctx context.Context, userID domain.ID) ([]domain.Certificate, error)
	Create(ctx context.Context, cert domain.Certificate) (domain.Certificate, error)
}

type IStatRepository interface {
	FindLessonStat(ctx context.Context, userID, lessonID domain.ID) (domain.LessonStat, error)
	CreateLessonStat(ctx context.Context, stat domain.LessonStat) error
	UpdateLessonStat(ctx context.Context, stat domain.LessonStat) error
}
