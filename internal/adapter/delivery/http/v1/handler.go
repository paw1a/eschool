package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/auth/jwt"
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
	config        *config.Config
	tokenProvider jwt.TokenProvider
	userService   port.IUserService
}

type HandlerParams struct {
	fx.In
	Config        *config.Config
	TokenProvider jwt.TokenProvider
	UserService   port.IUserService
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		config:        params.Config,
		tokenProvider: params.TokenProvider,
		userService:   params.UserService,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	v1.Use(LoggerMiddleware())
	{
		h.initAuthRoutes(v1)
		h.initUsersRoutes(v1)
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

func getIdFromRequestContext(context *gin.Context, paramName string) (domain.ID, error) {
	id, ok := context.Get(paramName)
	if !ok {
		return domain.RandomID(), errors.New("not authenticated")
	}
	return domain.ID(id.(string)), nil
}
