package domain

import "github.com/guregu/null"

type User struct {
	ID        ID
	Name      string
	Surname   string
	Phone     null.String
	City      null.String
	AvatarUrl null.String
	Email     string
	Password  string
}
