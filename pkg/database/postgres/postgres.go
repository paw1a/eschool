package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type Config struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type DB struct {
	Root          *sqlx.DB
	Guest         *sqlx.DB
	Authenticated *sqlx.DB
}

const (
	maxConn         = 100
	maxConnIdleTime = 1 * time.Minute
	maxConnLifetime = 3 * time.Minute
)

func NewPostgresDB(cfg *Config, logger *zap.Logger) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Database,
		cfg.Password,
	)

	db, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		logger.Fatal("failed to connect postgres db: %s", zap.String("conn string", connectionString))
		return nil, err
	}

	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetConnMaxIdleTime(maxConnIdleTime)

	err = db.Ping()
	if err != nil {
		logger.Fatal("failed to ping postgres db: %s", zap.String("conn string", connectionString))
		return nil, err
	}

	return db, nil
}
