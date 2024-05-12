package port

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"io"
	"net/url"
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
	FindCourseTeachers(ctx context.Context, courseID domain.ID) ([]domain.User, error)
	IsCourseStudent(ctx context.Context, studentID, courseID domain.ID) (bool, error)
	IsCourseTeacher(ctx context.Context, teacherID, courseID domain.ID) (bool, error)
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
	CreateTheoryLesson(ctx context.Context, courseID domain.ID,
		param CreateTheoryParam) (domain.Lesson, error)
	CreateVideoLesson(ctx context.Context, courseID domain.ID,
		param CreateVideoParam) (domain.Lesson, error)
	CreatePracticeLesson(ctx context.Context, courseID domain.ID,
		param CreatePracticeParam) (domain.Lesson, error)
	UpdateTheoryLesson(ctx context.Context, lessonID domain.ID,
		param UpdateTheoryParam) (domain.Lesson, error)
	UpdateVideoLesson(ctx context.Context, lessonID domain.ID,
		param UpdateVideoParam) (domain.Lesson, error)
	UpdatePracticeLesson(ctx context.Context, lessonID domain.ID,
		param UpdatePracticeParam) (domain.Lesson, error)
	Delete(ctx context.Context, lessonID domain.ID) error
}

type ISchoolService interface {
	FindAll(ctx context.Context) ([]domain.School, error)
	FindByID(ctx context.Context, schoolID domain.ID) (domain.School, error)
	FindUserSchools(ctx context.Context, userID domain.ID) ([]domain.School, error)
	FindSchoolCourses(ctx context.Context, schoolID domain.ID) ([]domain.Course, error)
	FindSchoolTeachers(ctx context.Context, schoolID domain.ID) ([]domain.User, error)
	IsSchoolTeacher(ctx context.Context, schoolID, teacherID domain.ID) (bool, error)
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
	FindCourseCertificate(ctx context.Context, courseID, userID domain.ID) (domain.Certificate, error)
	FindUserCertificates(ctx context.Context, userID domain.ID) ([]domain.Certificate, error)
	CreateCourseCertificate(ctx context.Context, userID, courseID domain.ID) (domain.Certificate, error)
}

type IStatService interface {
	FindLessonStat(ctx context.Context, userID, lessonID domain.ID) (domain.LessonStat, error)
	CreateLessonStat(ctx context.Context, userID, lessonID domain.ID) error
	UpdateLessonStat(ctx context.Context, userID, lessonID domain.ID,
		param UpdateLessonStatParam) error
}

type IPaymentService interface {
	GetCoursePaymentUrl(ctx context.Context, userID, courseID domain.ID) (url.URL, error)
	ProcessCoursePayment(ctx context.Context, label string, paid int64) (domain.PaymentPayload, error)
}

type IMediaService interface {
	SaveMediaFile(ctx context.Context, file domain.File) (domain.Url, error)
	SaveUserAvatar(ctx context.Context, userID domain.ID, file domain.File) (domain.Url, error)
	SaveLessonTheory(ctx context.Context, lessonID domain.ID, reader io.Reader) (domain.Url, error)
	SaveTestQuestion(ctx context.Context, testID domain.ID, reader io.Reader) (domain.Url, error)
}

type IAuthTokenService interface {
	SignIn(ctx context.Context, param SignInParam) (domain.AuthDetails, error)
	SignUp(ctx context.Context, param SignUpParam) error
	LogOut(ctx context.Context, refreshToken domain.Token) error
	Refresh(ctx context.Context, refreshToken domain.Token, fingerprint string) (domain.AuthDetails, error)
	Verify(ctx context.Context, accessToken domain.Token) error
	Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error)
}
