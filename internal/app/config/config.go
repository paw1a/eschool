package config

import (
	"github.com/paw1a/eschool-auth/jwt"
	v1 "github.com/paw1a/eschool-delivery/http/v1"
	"github.com/paw1a/eschool-payment/yoomoney"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/paw1a/eschool/pkg/database/redis"
	"github.com/paw1a/eschool/pkg/minio"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Web      v1.Config
	Postgres postgres.Config
	JWT      jwt.Config
	Redis    redis.Config
	Minio    minio.Config
	Yoomoney yoomoney.Config
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

		log.Infof("read config file: config/config.yml")
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
	bindings["postgres.database"] = "DB_NAME"
	bindings["postgres.user"] = "DB_USER"
	bindings["postgres.password"] = "DB_PASSWORD"
	bindings["postgres.host"] = "DB_HOST"
	bindings["postgres.port"] = "DB_PORT"
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
