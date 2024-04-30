package port

import (
	"github.com/guregu/null"
)

type CreateUserParam struct {
	Name      string
	Surname   string
	Email     string
	Password  string
	Phone     null.String
	City      null.String
	AvatarUrl null.String
}

type UpdateUserParam struct {
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
