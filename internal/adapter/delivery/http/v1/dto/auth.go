package dto

import "github.com/guregu/null"

type SignUpDTO struct {
	Name      string      `json:"name" binding:"required"`
	Surname   string      `json:"surname" binding:"required"`
	Email     string      `json:"email" binding:"required"`
	Password  string      `json:"password" binding:"required"`
	Phone     null.String `json:"phone"`
	City      null.String `json:"city"`
	AvatarUrl null.String `json:"avatar_url"`
}

type SignInDTO struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Fingerprint string `json:"fingerprint" binding:"required"`
}

type RefreshDTO struct {
	Fingerprint string `json:"fingerprint" binding:"required"`
}
