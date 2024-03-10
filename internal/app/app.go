package app

import (
	"github.com/paw1a/eschool/internal/config"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/paw1a/eschool/pkg/database/redis"
	log "github.com/sirupsen/logrus"
)

func Run() {
	log.Info("application startup...")
	log.Info("logger initialized")

	cfg := config.GetConfig()
	log.Info("config created")

	_, err := postgres.NewPgxPool(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("postgres is connected")

	_, err = redis.NewClient(&cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("redis is connected")
}
