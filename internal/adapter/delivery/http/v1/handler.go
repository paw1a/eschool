package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"math"
	"net/http"
	"time"
)

type Config struct {
	Host string
	Port string
}

type Handler struct {
	config         *Config
	logger         *zap.Logger
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
	Config         *Config
	Logger         *zap.Logger
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

func NewHandler(params HandlerParams, router *gin.Engine) *Handler {
	handler := &Handler{
		config:         params.Config,
		logger:         params.Logger,
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

	v1 := router.Group("/api/v1")
	v1.Use(LoggerMiddleware(params.Logger))
	{
		handler.initAuthRoutes(v1)
		handler.initUsersRoutes(v1)
		handler.initCourseRoutes(v1)
		handler.initSchoolRoutes(v1)
		handler.initPaymentRoutes(v1)
	}

	return handler
}

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()

		if len(c.Errors) > 0 {
			logger.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%s %d] %s (%dms)", c.Request.Method, statusCode, path, latency)
			if statusCode >= http.StatusInternalServerError {
				logger.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				logger.Warn(msg)
			} else {
				logger.Info(msg)
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
