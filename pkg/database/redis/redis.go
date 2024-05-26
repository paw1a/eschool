package redis

import (
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

type Config struct {
	URI string
}

func NewClient(cfg *Config, logger *zap.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.URI,
	})

	_, err := client.Ping().Result()
	if err != nil {
		logger.Error("failed to ping storage", zap.String("conn string", cfg.URI))
		return nil, err
	}

	return client, nil
}
