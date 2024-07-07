package service_test

import (
	"context"
	"errors"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
	"github.com/paw1a/eschool/internal/core/service"
	mocks "github.com/paw1a/eschool/internal/core/service/mocks/repository"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

var testUsers = []domain.User{
	{
		ID:        domain.NewID(),
		Name:      "Username1",
		Surname:   "Surname1",
		Phone:     null.StringFrom("+79261414820"),
		City:      null.String{},
		AvatarUrl: null.String{},
		Email:     "user1@gmail.com",
		Password:  "123456",
	},
	{
		ID:        domain.NewID(),
		Name:      "Username2",
		Surname:   "Surname2",
		Phone:     null.String{},
		City:      null.String{},
		AvatarUrl: null.String{},
		Email:     "user2@gmail.com",
		Password:  "qwerty",
	},
}

var unknownID = domain.RandomID()

func TestUserService_FindByID(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(userRepo *mocks.UserRepository)
		user         domain.User
		hasError     bool
	}{
		{
			name: "user found, ok",
			user: testUsers[0],
			initRepoMock: func(userRepo *mocks.UserRepository) {
				userRepo.On("FindByID", context.Background(), testUsers[0].ID).
					Return(testUsers[0], nil)
			},
			hasError: false,
		},
		{
			name: "user not found, error",
			user: domain.User{ID: unknownID},
			initRepoMock: func(userRepo *mocks.UserRepository) {
				userRepo.On("FindByID", context.Background(), unknownID).
					Return(domain.User{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			userRepo := mocks.NewUserRepository(t)
			userService := service.NewUserService(userRepo, zap.NewNop())

			test.initRepoMock(userRepo)

			user, err := userService.FindByID(context.Background(), test.user.ID)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.user.ID, user.ID)
			}
		})
	}
}

func TestUserService_FindByCredentials(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(userRepo *mocks.UserRepository)
		user         domain.User
		hasError     bool
	}{
		{
			name: "user found, ok",
			user: testUsers[0],
			initRepoMock: func(userRepo *mocks.UserRepository) {
				userRepo.On("FindByCredentials", context.Background(),
					testUsers[0].Email, testUsers[0].Password).
					Return(testUsers[0], nil)
			},
			hasError: false,
		},
		{
			name: "user no such email, error",
			user: domain.User{
				Email:    "unknown@gmail.com",
				Password: testUsers[0].Password,
			},
			initRepoMock: func(userRepo *mocks.UserRepository) {
				userRepo.On("FindByCredentials", context.Background(),
					"unknown@gmail.com", testUsers[0].Password).
					Return(domain.User{}, errors.New("error"))
			},
			hasError: true,
		},
		{
			name: "user invalid password, error",
			user: domain.User{
				Email:    testUsers[0].Email,
				Password: "invalid password",
			},
			initRepoMock: func(userRepo *mocks.UserRepository) {
				userRepo.On("FindByCredentials", context.Background(),
					testUsers[0].Email, "invalid password").
					Return(domain.User{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			userRepo := mocks.NewUserRepository(t)
			userService := service.NewUserService(userRepo, zap.NewNop())

			test.initRepoMock(userRepo)

			user, err := userService.FindByCredentials(context.Background(), port.UserCredentials{
				Email:    test.user.Email,
				Password: test.user.Password,
			})

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.user.ID, user.ID)
			}
		})
	}
}

func TestUserService_FindUserInfo(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(userRepo *mocks.UserRepository)
		user         domain.User
		expectedInfo port.UserInfo
		hasError     bool
	}{
		{
			name: "user info found, ok",
			user: testUsers[0],
			expectedInfo: port.UserInfo{
				Name:    testUsers[0].Name,
				Surname: testUsers[0].Surname,
			},
			initRepoMock: func(userRepo *mocks.UserRepository) {
				userRepo.On("FindUserInfo", context.Background(), testUsers[0].ID).
					Return(port.UserInfo{
						Name:    testUsers[0].Name,
						Surname: testUsers[0].Surname,
					}, nil)
			},
			hasError: false,
		},
		{
			name:         "user info not found, error",
			user:         domain.User{ID: unknownID},
			expectedInfo: port.UserInfo{},
			initRepoMock: func(userRepo *mocks.UserRepository) {
				userRepo.On("FindUserInfo", context.Background(), unknownID).
					Return(port.UserInfo{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			userRepo := mocks.NewUserRepository(t)
			userService := service.NewUserService(userRepo, zap.NewNop())

			test.initRepoMock(userRepo)

			userInfo, err := userService.FindUserInfo(context.Background(), test.user.ID)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.expectedInfo, userInfo)
			}
		})
	}
}
