package test

import (
	"context"
	repository "github.com/paw1a/eschool/internal/adapter/repository/gorm"
	"testing"
)

func Benchmark_FindAll(b *testing.B) {
	container, err := newPostgresContainer(context.Background())
	if err != nil {
		return
	}

	url, err := container.ConnectionString(context.Background())
	if err != nil {
		return
	}

	db, err := newGormDB(url)
	if err != nil {
		return
	}
	userRepo := repository.NewGormUserRepository(db)
	for range b.N {
		_, err := userRepo.FindAll(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
