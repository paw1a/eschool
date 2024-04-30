package dto

import "github.com/guregu/null"

type UpdateUserDTO struct {
	Name      null.String `json:"name"`
	Surname   null.String `json:"surname"`
	Phone     null.String `json:"phone"`
	City      null.String `json:"city"`
	AvatarUrl null.String `json:"avatar_url"`
}

type UserInfoDTO struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

type UserDTO struct {
	ID        string `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"omitempty"`
	City      string `json:"city" binding:"omitempty"`
	AvatarUrl string `json:"avatar_url" binding:"omitempty"`
}
