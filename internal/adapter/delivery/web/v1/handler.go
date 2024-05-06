package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"math"
	"net/http"
	"time"
)

type Handler struct {
	config             *config.Config
	userService        port.IUserService
	schoolService      port.ISchoolService
	lessonService      port.ILessonService
	reviewService      port.IReviewService
	courseService      port.ICourseService
	certificateService port.ICertificateService
	mediaService       port.IMediaService
	statService        port.IStatService
	authService        port.IAuthTokenService
	paymentService     port.IPaymentService
}

type HandlerParams struct {
	fx.In
	Config             *config.Config
	UserService        port.IUserService
	SchoolService      port.ISchoolService
	LessonService      port.ILessonService
	ReviewService      port.IReviewService
	CourseService      port.ICourseService
	CertificateService port.ICertificateService
	MediaService       port.IMediaService
	StatService        port.IStatService
	AuthService        port.IAuthTokenService
	PaymentService     port.IPaymentService
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		config:             params.Config,
		userService:        params.UserService,
		schoolService:      params.SchoolService,
		lessonService:      params.LessonService,
		reviewService:      params.ReviewService,
		courseService:      params.CourseService,
		certificateService: params.CertificateService,
		mediaService:       params.MediaService,
		statService:        params.StatService,
		authService:        params.AuthService,
		paymentService:     params.PaymentService,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	v1.Use(LoggerMiddleware())
	{
		h.initAuthRoutes(v1)
		h.initUsersRoutes(v1)
		h.initCourseRoutes(v1)
		h.initSchoolRoutes(v1)
		h.initPaymentRoutes(v1)
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()

		if len(c.Errors) > 0 {
			log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%s %d] %s (%dms)", c.Request.Method, statusCode, path, latency)
			if statusCode >= http.StatusInternalServerError {
				log.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				log.Warn(msg)
			} else {
				log.Info(msg)
			}
		}
	}
}

func getIdFromPath(c *gin.Context, paramName string) (domain.ID, error) {
	idString := c.Param(paramName)
	if idString == "" {
		return domain.RandomID(), PathIdParamIsEmptyError
	}

	if _, err := uuid.Parse(idString); err != nil {
		return domain.RandomID(), PathIdParamIsInvalidUUID
	}
	return domain.ID(idString), nil
}

func getIdFromRequestContext(context *gin.Context) (domain.ID, error) {
	id, ok := context.Get("userID")
	if !ok {
		return domain.RandomID(), UnauthorizedError
	}
	return domain.ID(id.(string)), nil
}
