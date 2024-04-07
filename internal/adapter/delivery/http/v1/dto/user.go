package dto

import "github.com/guregu/null"

type SignUpDTO struct {
	Name      string      `json:"name"`
	Surname   string      `json:"surname"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Phone     null.String `json:"phone"`
	City      null.String `json:"city"`
	AvatarUrl null.String `json:"avatar_url"`
}

type UpdateUserDTO struct {
	Name      null.String `json:"name"`
	Surname   null.String `json:"surname"`
	Phone     null.String `json:"phone"`
	City      null.String `json:"city"`
	AvatarUrl null.String `json:"avatar_url"`
}

type SignInDTO struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Fingerprint string `json:"fingerprint"`
}

type UserInfo struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
