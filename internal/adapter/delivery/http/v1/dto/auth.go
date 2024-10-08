package dto

import "github.com/guregu/null"

type SignUpDTO struct {
	Name      string      `json:"name" binding:"required" example:"Maxim"`
	Surname   string      `json:"surname" binding:"required" example:"Ivanov"`
	Email     string      `json:"email" binding:"required,email" example:"user@gmail.com"`
	Password  string      `json:"password" binding:"required" example:"123"`
	Phone     null.String `json:"phone" binding:"omitempty" swaggertype:"string" example:"+79999999999"`
	City      null.String `json:"city" binding:"omitempty" swaggertype:"string" example:"Moscow"`
	AvatarUrl null.String `json:"avatar_url" binding:"omitempty,url" swaggertype:"string" example:"image.io/avatar.png"`
}

type SignInDTO struct {
	Email       string `json:"email" binding:"required,email" example:"paw1a@yandex.ru"`
	Password    string `json:"password" binding:"required" example:"123"`
	Fingerprint string `json:"fingerprint" binding:"required" example:"fingerprint"`
}

type RefreshDTO struct {
	Fingerprint string `json:"fingerprint" binding:"required" example:"fingerprint"`
}
