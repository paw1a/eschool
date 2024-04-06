package dto

import "github.com/guregu/null"

type CreateUserDTO struct {
	Name      string
	Surname   string
	Email     string
	Password  string
	Phone     null.String
	City      null.String
	AvatarUrl null.String
}

type UpdateUserDTO struct {
	Name      null.String
	Surname   null.String
	Phone     null.String
	City      null.String
	AvatarUrl null.String
}

type UserCredentials struct {
	Email    string
	Password string
}

type UserInfo struct {
	Name    string
	Surname string
}
