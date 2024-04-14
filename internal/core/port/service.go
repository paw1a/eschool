package port

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"io"
)

type IUserService interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID domain.ID) (domain.User, error)
	FindByCredentials(ctx context.Context, credentials UserCredentials) (domain.User, error)
	FindUserInfo(ctx context.Context, userID domain.ID) (UserInfo, error)
	Create(ctx context.Context, param CreateUserParam) (domain.User, error)
	Update(ctx context.Context, userID domain.ID, param UpdateUserParam) (domain.User, error)
	Delete(ctx context.Context, userID domain.ID) error
}

type ICourseService interface {
	FindAll(ctx context.Context) ([]domain.Course, error)
	FindByID(ctx context.Context, courseID domain.ID) (domain.Course, error)
	FindStudentCourses(ctx context.Context, studentID domain.ID) ([]domain.Course, error)
	FindTeacherCourses(ctx context.Context, teacherID domain.ID) ([]domain.Course, error)
	AddCourseStudent(ctx context.Context, studentID, courseID domain.ID) error
	AddCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) error
	ConfirmDraftCourse(ctx context.Context, courseID domain.ID) []error
	PublishReadyCourse(ctx context.Context, courseID domain.ID) error
	CreateSchoolCourse(ctx context.Context, schoolID domain.ID,
		param CreateCourseParam) (domain.Course, error)
	Update(ctx context.Context, courseID domain.ID,
		param UpdateCourseParam) (domain.Course, error)
	Delete(ctx context.Context, courseID domain.ID) error
}

type ILessonService interface {
	FindAll(ctx context.Context) ([]domain.Lesson, error)
	FindByID(ctx context.Context, lessonID domain.ID) (domain.Lesson, error)
	FindCourseLessons(ctx context.Context, courseID domain.ID) ([]domain.Lesson, error)
	CreateCourseLesson(ctx context.Context, courseID domain.ID,
		param CreateLessonParam) (domain.Lesson, error)
	AddLessonTests(ctx context.Context, lessonID domain.ID, tests []CreateTestParam) error
	UpdateLessonTest(ctx context.Context, testID domain.ID, param UpdateTestParam) (domain.Test, error)
	UpdateLessonTheory(ctx context.Context, lessonID domain.ID, param UpdateTheoryParam) error
	UpdateLessonVideo(ctx context.Context, lessonID domain.ID, param UpdateVideoParam) error
	Delete(ctx context.Context, lessonID domain.ID) error
	DeleteLessonTest(ctx context.Context, testID domain.ID) error
}

type ISchoolService interface {
	FindAll(ctx context.Context) ([]domain.School, error)
	FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error)
	FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error)
	FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error)
	AddSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) error
	CreateUserSchool(ctx context.Context, userID domain.ID, param CreateSchoolParam) (domain.School, error)
	Update(ctx context.Context, schoolID domain.ID, param UpdateSchoolParam) (domain.School, error)
	Delete(ctx context.Context, schoolID domain.ID) error
}

type IReviewService interface {
	FindAll(ctx context.Context) ([]domain.Review, error)
	FindByID(ctx context.Context, reviewID domain.ID) (domain.Review, error)
	FindUserReviews(ctx context.Context, userID domain.ID) ([]domain.Review, error)
	FindCourseReviews(ctx context.Context, courseID domain.ID) ([]domain.Review, error)
	CreateCourseReview(ctx context.Context, courseID, userID domain.ID,
		param CreateReviewParam) (domain.Review, error)
	Delete(ctx context.Context, reviewID domain.ID) error
}

type ICertificateService interface {
	FindAll(ctx context.Context) ([]domain.Certificate, error)
	FindByID(ctx context.Context, certificateID domain.ID) (domain.Certificate, error)
	FindUserCertificates(ctx context.Context, userID domain.ID) ([]domain.Certificate, error)
	CreateCourseCertificate(ctx context.Context, userID, courseID domain.ID) (domain.Certificate, error)
}

type IStatisticsService interface {
	FindUserLessonStat(ctx context.Context, userID, lessonID domain.ID) (domain.LessonStat, error)
	FindUserTestStat(ctx context.Context, userID, testID domain.ID) (domain.TestStat, error)
	CreateUserLessonStat(ctx context.Context, userID, lessonID domain.ID) error
	CreateUserTestStat(ctx context.Context, userID, testID domain.ID) error
	UpdateUserLessonStat(ctx context.Context, userID, lessonID domain.ID,
		param UpdateLessonStatParam) error
	UpdateUserTestStat(ctx context.Context, userID, testID domain.ID,
		param UpdateTestStatParam) error
}

type IPaymentService interface {
	PayCourse(ctx context.Context, userID, courseID domain.ID) error
}

type IMediaService interface {
	SaveMediaFile(ctx context.Context, file domain.File) (domain.Url, error)
	SaveUserAvatar(ctx context.Context, userID domain.ID, file domain.File) (domain.Url, error)
	SaveLessonTheory(ctx context.Context, lessonID domain.ID, reader io.Reader) (domain.Url, error)
	SaveTestQuestion(ctx context.Context, testID domain.ID, reader io.Reader) (domain.Url, error)
}
