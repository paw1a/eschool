package test

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate"
	migratepg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	testpg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"path/filepath"
	"runtime"
)

type Config struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

var (
	postgresConfig = Config{
		Database: "eschool",
		User:     "postgres",
		Password: "password",
	}
)

func newPostgresContainer(ctx context.Context) (*testpg.PostgresContainer, error) {
	container, err := testpg.Run(
		ctx,
		"docker.io/postgres:16-alpine",
		testpg.WithDatabase(postgresConfig.Database),
		testpg.WithUsername(postgresConfig.User),
		testpg.WithPassword(postgresConfig.Password),
		testpg.BasicWaitStrategies(),
		testpg.WithSQLDriver("pgx"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get caller path")
	}

	sourceUrl := "file://" + filepath.Dir(path) + "/migrations"
	url, err := container.ConnectionString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get postgres db url: %s", err)
	}

	db, err := newGormDB(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres db: %s", err)
	}

	gormDB, _ := db.DB()
	driver, err := migratepg.WithInstance(gormDB, &migratepg.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to get db driver from instance: %s", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(sourceUrl, postgresConfig.Database, driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrator driver: %s", err)
	}

	err = mig.Up()
	if err != nil {
		return nil, fmt.Errorf("failed to up migrations: %s", err)
	}

	return container, nil
}

func newGormDB(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect postgres db: %s", url)
		return nil, err
	}

	return db, nil
}
