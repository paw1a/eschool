package dto

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type UpdateUserDTO struct {
	Name      null.String `json:"name" binding:"omitempty"`
	Surname   null.String `json:"surname" binding:"omitempty"`
	Phone     null.String `json:"phone" binding:"omitempty"`
	City      null.String `json:"city" binding:"omitempty"`
	AvatarUrl null.String `json:"avatar_url" binding:"omitempty"`
}

type UserInfoDTO struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	AvatarUrl string `json:"avatar_url"`
}

func NewUserDTO(user domain.User) UserDTO {
	return UserDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Surname:   user.Surname,
		Phone:     user.Phone.String,
		City:      user.City.String,
		AvatarUrl: user.AvatarUrl.String,
		Email:     user.Email,
	}
}
