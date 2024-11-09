package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
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

func prometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(c.Path(), c.Request().Method).Observe(duration)
		requestsPerSecond.WithLabelValues(c.Path(), c.Request().Method).Inc()
		return err
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

	e := echo.New()
	e.Use(prometheusMiddleware)
	e.Use(middleware.Logger())
	e.Use(echoprometheus.NewMiddleware("echo"))
	e.GET("/metrics", echoprometheus.NewHandler())
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/users", func(c echo.Context) error {
		users, err := userService.FindAll(c.Request().Context())
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return err
		}

		userDTOs := make([]dto.UserDTO, len(users))
		for i, user := range users {
			userDTOs[i] = dto.NewUserDTO(user)
		}

		c.JSON(http.StatusOK, userDTOs)
		return nil
	})
	e.Logger.Fatal(e.Start(":8080"))
}
