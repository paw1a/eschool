package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/pkg/auth"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"math"
	"net/http"
	"time"
)

type Handler struct {
	config        *config.Config
	tokenProvider auth.TokenProvider
	userService   service.IUserService
}

type HandlerParams struct {
	fx.In
	Config        *config.Config
	TokenProvider auth.TokenProvider
	UserService   service.IUserService
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

func getIdFromRequestContext(context *gin.Context, paramName string) (int64, error) {
	id, ok := context.Get(paramName)
	if !ok {
		return 0, errors.New("not authenticated")
	}
	tempID, ok := id.(float64)
	if !ok {
		return 0, errors.New("not authenticated")
	}
	return int64(tempID), nil
}
