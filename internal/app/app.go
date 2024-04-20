package app

import (
	"github.com/paw1a/eschool/internal/adapter/auth/jwt"
	delivery "github.com/paw1a/eschool/internal/adapter/delivery/http"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1"
	pgRepo "github.com/paw1a/eschool/internal/adapter/repository/postgres"
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
			minio.NewMinioClient,
			fx.Annotate(
				jwt.NewTokenProvider,
				fx.As(new(jwt.TokenProvider)),
			),
			fx.Annotate(
				pgRepo.NewUserRepo,
				fx.As(new(port.IUserRepository)),
			),
			fx.Annotate(
				service.NewUserService,
				fx.As(new(port.IUserService)),
			),
		),
		fx.Supply(cfg, &cfg.Redis, &cfg.Postgres, &cfg.JWT, &cfg.Minio),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
