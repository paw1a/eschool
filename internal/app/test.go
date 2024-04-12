package app

import (
	"context"
	"fmt"
	"github.com/guregu/null"
	repository "github.com/paw1a/eschool/internal/adapter/repository/postgres"
	"github.com/paw1a/eschool/internal/app/config"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/pkg/database/postgres"
	log "github.com/sirupsen/logrus"
)

func Test() {
	log.Info("application startup")
	log.Info("logger initialized")

	cfg := config.GetConfig()
	log.Info("config created")

	db, err := postgres.NewPostgresDB(&cfg.Postgres)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// find all users
	repo := repository.NewUserRepo(db)
	users, err := repo.FindAll(context.Background())
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%v\n", users)

	// find user by id
	user, err := repo.FindByID(context.Background(), "30e18bc1-4354-4937-9a3b-03cf0b7027ca")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%v\n", user)

	// create user
	user = domain.User{
		ID:        domain.NewID(),
		Name:      "testUser",
		Surname:   "surname",
		Phone:     null.StringFrom("+79999999999"),
		City:      null.String{},
		AvatarUrl: null.String{},
		Email:     "testuser@gmail.com",
		Password:  "123456",
	}
	user, err = repo.Create(context.Background(), user)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%v\n", user)

	// update user
	updatedUser := user
	updatedUser.City = null.StringFrom("Moscow")
	updatedUser.Name = "updatedName"
	user, err = repo.Update(context.Background(), updatedUser)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("%v\n", user)

	// delete user
	err = repo.Delete(context.Background(), user.ID)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
