package main

import (
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram for the duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
	requestsPerSecond = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests processed, labeled by status and path",
		},
		[]string{"path", "method"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestsPerSecond)
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(c.FullPath(), c.Request.Method).Observe(duration)
		requestsPerSecond.WithLabelValues(c.FullPath(), c.Request.Method).Inc()
	}
}

func main() {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	logger, _ := loggerBuilder.Build()
	cfg := config.GetConfig()
	db, err := postgres.NewPostgresDB(&cfg.Postgres, logger)
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo, logger)

	r := gin.Default()
	r.Use(prometheusMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/users", func(c *gin.Context) {
		users, err := userService.FindAll(c.Request.Context())
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		userDTOs := make([]dto.UserDTO, len(users))
		for i, user := range users {
			userDTOs[i] = dto.NewUserDTO(user)
		}

		c.JSON(http.StatusOK, userDTOs)
	})
	r.Run(":8080")
}
