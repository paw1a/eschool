package console

import (
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/fx"
)

type Handler struct {
	userService    port.IUserService
	schoolService  port.ISchoolService
	lessonService  port.ILessonService
	reviewService  port.IReviewService
	courseService  port.ICourseService
	mediaService   port.IMediaService
	statService    port.IStatService
	authService    port.IAuthTokenService
	paymentService port.IPaymentService
}

type HandlerParams struct {
	fx.In
	UserService    port.IUserService
	SchoolService  port.ISchoolService
	LessonService  port.ILessonService
	ReviewService  port.IReviewService
	CourseService  port.ICourseService
	MediaService   port.IMediaService
	StatService    port.IStatService
	AuthService    port.IAuthTokenService
	PaymentService port.IPaymentService
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		userService:    params.UserService,
		schoolService:  params.SchoolService,
		lessonService:  params.LessonService,
		reviewService:  params.ReviewService,
		courseService:  params.CourseService,
		mediaService:   params.MediaService,
		statService:    params.StatService,
		authService:    params.AuthService,
		paymentService: params.PaymentService,
	}
}

func getIdFromConsole(c *Console) (domain.ID, error) {
	return domain.RandomID(), nil
}
