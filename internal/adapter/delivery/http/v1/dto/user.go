package dto

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type UpdateUserDTO struct {
	Name      null.String `json:"name" binding:"omitempty" swaggertype:"string" example:"Pavel"`
	Surname   null.String `json:"surname" binding:"omitempty" swaggertype:"string" example:"Shpakovskiy"`
	Phone     null.String `json:"phone" binding:"omitempty" swaggertype:"string" example:"+79999999999"`
	City      null.String `json:"city" binding:"omitempty" swaggertype:"string" example:"Moscow"`
	AvatarUrl null.String `json:"avatar_url" binding:"omitempty" swaggertype:"string" example:"image.io/avatar.png"`
}

type UserInfoDTO struct {
	Name    string `json:"name" example:"Pavel"`
	Surname string `json:"surname" example:"Shpakovskiy"`
}

type UserDTO struct {
	ID        string `json:"id" example:"30e18bc1-4354-4937-9a3b-03cf0b7027ca"`
	Name      string `json:"name" example:"Pavel"`
	Surname   string `json:"surname" example:"Shpakovskiy"`
	Email     string `json:"email" example:"paw1a@yandex.ru"`
	Phone     string `json:"phone" example:"+79999999999"`
	City      string `json:"city" example:"Moscow"`
	AvatarUrl string `json:"avatar_url" example:"image.io/avatar.png"`
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
