package config

import (
	"github.com/paw1a/eschool/internal/adapter/auth/jwt"
	v1 "github.com/paw1a/eschool/internal/adapter/delivery/http/v1"
	"github.com/paw1a/eschool/internal/adapter/payment/yoomoney"
	storage "github.com/paw1a/eschool/internal/adapter/storage/minio"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/paw1a/eschool/pkg/database/redis"
	"github.com/paw1a/eschool/pkg/logging"
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	Logging               logging.Config  `mapstructure:"logging"`
	Web                   v1.Config       `mapstructure:"web"`
	PostgresRoot          postgres.Config `mapstructure:"db_root"`
	PostgresGuest         postgres.Config `mapstructure:"db_guest"`
	PostgresAuthenticated postgres.Config `mapstructure:"db_authenticated"`
	JWT                   jwt.Config      `mapstructure:"jwt"`
	Redis                 redis.Config    `mapstructure:"redis"`
	Minio                 storage.Config  `mapstructure:"minio"`
	Yoomoney              yoomoney.Config `mapstructure:"yoomoney"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath("config")
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		if err := bindEnvConfig(); err != nil {
			log.Fatalf("error reading config file: %v", err)
		}

		log.Println("read config file: config/config.yml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error reading config file: %v", err)
		}

		instance = &Config{}
		if err := viper.Unmarshal(&instance); err != nil {
			log.Fatalf("error unmarshaling config file: %v", err)
		}
	})
	return instance
}

func bindEnvConfig() error {
	bindings := make(map[string]string)
	bindings["web.host"] = "HOST"
	bindings["web.port"] = "PORT"
	bindings["jwt.secret"] = "JWT_SECRET"

	bindings["db_root.database"] = "DB_NAME"
	bindings["db_root.host"] = "DB_HOST"
	bindings["db_root.port"] = "DB_PORT"
	bindings["db_root.user"] = "DB_ROOT_USER"
	bindings["db_root.password"] = "DB_ROOT_PASSWORD"

	bindings["db_guest.database"] = "DB_NAME"
	bindings["db_guest.host"] = "DB_HOST"
	bindings["db_guest.port"] = "DB_PORT"
	bindings["db_guest.user"] = "DB_GUEST_USER"
	bindings["db_guest.password"] = "DB_GUEST_PASSWORD"

	bindings["db_authenticated.database"] = "DB_NAME"
	bindings["db_authenticated.host"] = "DB_HOST"
	bindings["db_authenticated.port"] = "DB_PORT"
	bindings["db_authenticated.user"] = "DB_AUTHENTICATED_USER"
	bindings["db_authenticated.password"] = "DB_AUTHENTICATED_PASSWORD"

	bindings["redis.uri"] = "REDIS_URI"
	bindings["minio.endpoint"] = "MINIO_ENDPOINT"
	bindings["minio.user"] = "MINIO_ROOT_USER"
	bindings["minio.password"] = "MINIO_ROOT_PASSWORD"
	bindings["minio.bucketName"] = "MINIO_BUCKET_NAME"
	bindings["yoomoney.scheme"] = "PAYMENT_SCHEME"
	bindings["yoomoney.host"] = "PAYMENT_HOST"
	bindings["yoomoney.path"] = "PAYMENT_PATH"
	bindings["yoomoney.wallet"] = "PAYMENT_WALLET"

	for name, binding := range bindings {
		if err := viper.BindEnv(name, binding); err != nil {
			return err
		}
	}

	return nil
}
