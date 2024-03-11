package app

import (
	"github.com/paw1a/eschool/internal/config"
	delivery "github.com/paw1a/eschool/internal/delivery/http"
	v1 "github.com/paw1a/eschool/internal/delivery/http/v1"
	"github.com/paw1a/eschool/internal/repository"
	pgRepo "github.com/paw1a/eschool/internal/repository/postgres"
	"github.com/paw1a/eschool/internal/service"
	"github.com/paw1a/eschool/pkg/auth"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/paw1a/eschool/pkg/database/redis"
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
			postgres.NewPgxPool,
			redis.NewClient,
			fx.Annotate(
				auth.NewTokenProvider,
				fx.As(new(auth.TokenProvider)),
			),
			fx.Annotate(
				pgRepo.NewUsersRepo,
				fx.As(new(repository.Users)),
			),
			fx.Annotate(
				service.NewUsersService,
				fx.As(new(service.Users)),
			),
		),
		fx.Supply(cfg, &cfg.Redis, &cfg.Postgres, &cfg.JWT),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
