package app

import (
	sessionStorage "github.com/paw1a/eschool/internal/adapter/auth/adapter/storage/redis"
	"github.com/paw1a/eschool/internal/adapter/auth/jwt"
	authPort "github.com/paw1a/eschool/internal/adapter/auth/port"
	delivery "github.com/paw1a/eschool/internal/adapter/delivery/http"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1"
	"github.com/paw1a/eschool/internal/adapter/payment/yoomoney"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	storage "github.com/paw1a/eschool/internal/adapter/storage/minio"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/paw1a/eschool/pkg/database/redis"
	"github.com/paw1a/eschool/pkg/minio"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

func Run() {
	log.Info("application startup")
	log.Info("logger initialized")

	cfg := config.GetConfig()
	log.Info("config created")

	fx.New(
		fx.Provide(
			NewServer,
			NewGinRouter,
			delivery.NewHandler,
			v1.NewHandler,
			postgres.NewPostgresDB,
			redis.NewClient,
			minio.NewClient,
			// repositories
			fx.Annotate(
				repository.NewUserRepo,
				fx.As(new(port.IUserRepository)),
			),
			fx.Annotate(
				repository.NewCourseRepo,
				fx.As(new(port.ICourseRepository)),
			),
			fx.Annotate(
				repository.NewSchoolRepo,
				fx.As(new(port.ISchoolRepository)),
			),
			fx.Annotate(
				repository.NewLessonRepo,
				fx.As(new(port.ILessonRepository)),
			),
			fx.Annotate(
				repository.NewReviewRepo,
				fx.As(new(port.IReviewRepository)),
			),
			fx.Annotate(
				repository.NewCertificateRepo,
				fx.As(new(port.ICertificateRepository)),
			),
			fx.Annotate(
				repository.NewStatRepo,
				fx.As(new(port.IStatRepository)),
			),
			fx.Annotate(
				storage.NewObjectStorage,
				fx.As(new(port.IObjectStorage)),
			),
			fx.Annotate(
				jwt.NewAuthProvider,
				fx.As(new(port.IAuthProvider)),
			),
			fx.Annotate(
				sessionStorage.NewSessionStorage,
				fx.As(new(authPort.ISessionStorage)),
			),
			fx.Annotate(
				yoomoney.NewPaymentGateway,
				fx.As(new(port.IPaymentGateway)),
			),
			// services
			fx.Annotate(
				service.NewUserService,
				fx.As(new(port.IUserService)),
			),
			fx.Annotate(
				service.NewCourseService,
				fx.As(new(port.ICourseService)),
			),
			fx.Annotate(
				service.NewSchoolService,
				fx.As(new(port.ISchoolService)),
			),
			fx.Annotate(
				service.NewLessonService,
				fx.As(new(port.ILessonService)),
			),
			fx.Annotate(
				service.NewReviewService,
				fx.As(new(port.IReviewService)),
			),
			fx.Annotate(
				service.NewCertificateService,
				fx.As(new(port.ICertificateService)),
			),
			fx.Annotate(
				service.NewMediaService,
				fx.As(new(port.IMediaService)),
			),
			fx.Annotate(
				service.NewStatService,
				fx.As(new(port.IStatService)),
			),
			fx.Annotate(
				service.NewPaymentService,
				fx.As(new(port.IPaymentService)),
			),
			fx.Annotate(
				service.NewAuthTokenService,
				fx.As(new(port.IAuthTokenService)),
			),
		),
		fx.Supply(cfg, &cfg.Redis, &cfg.Postgres, &cfg.JWT, &cfg.Minio, &cfg.Yoomoney),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
