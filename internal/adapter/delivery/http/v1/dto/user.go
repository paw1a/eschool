package dto

import "github.com/guregu/null"

type UpdateUserDTO struct {
	Name      null.String `json:"name"`
	Surname   null.String `json:"surname"`
	Phone     null.String `json:"phone"`
	City      null.String `json:"city"`
	AvatarUrl null.String `json:"avatar_url"`
}

type UserInfo struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}
