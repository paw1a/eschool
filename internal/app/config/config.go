package config

import (
	"github.com/paw1a/eschool/pkg/auth"
	"github.com/paw1a/eschool/pkg/database/postgres"
	"github.com/paw1a/eschool/pkg/database/redis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Postgres postgres.Config
	JWT      auth.Config
	Redis    redis.Config
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
	bindings["server.host"] = "HOST"
	bindings["server.port"] = "PORT"
	bindings["jwt.secret"] = "JWT_SECRET"
	bindings["postgres.database"] = "DB_NAME"
	bindings["postgres.user"] = "DB_USER"
	bindings["postgres.password"] = "DB_PASSWORD"
	bindings["postgres.host"] = "DB_HOST"
	bindings["postgres.port"] = "DB_PORT"
	bindings["redis.uri"] = "REDIS_URI"
	
	for name, binding := range bindings {
		if err := viper.BindEnv(name, binding); err != nil {
			return err
		}
	}

	return nil
}
