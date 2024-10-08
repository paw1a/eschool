package dto

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type CreateSchoolDTO struct {
	Name        string      `json:"name" binding:"required"`
	Description null.String `json:"description" binding:"required" swaggertype:"string" example:"School description"`
}

type UpdateSchoolDTO struct {
	Description null.String `json:"description" binding:"omitempty" swaggertype:"string" example:"Updated description"`
}

type SchoolDTO struct {
	ID          string `json:"id" example:"30e18bc1-4354-4937-9a3b-03cf0b7034cc"`
	OwnerID     string `json:"owner_id" example:"30e18bc1-4354-4937-9a3b-03cf0b7027ca"`
	Name        string `json:"name" example:"School name"`
	Description string `json:"description" example:"School description"`
}

func NewSchoolDTO(school domain.School) SchoolDTO {
	return SchoolDTO{
		ID:          school.ID.String(),
		OwnerID:     school.OwnerID.String(),
		Name:        school.Name,
		Description: school.Description,
	}
}
