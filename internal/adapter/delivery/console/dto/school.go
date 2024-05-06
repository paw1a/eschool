package dto

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type CreateSchoolDTO struct {
	Name        string      `json:"name" binding:"required"`
	Description null.String `json:"description" binding:"required"`
}

type UpdateSchoolDTO struct {
	Description null.String `json:"description" binding:"omitempty"`
}

type SchoolDTO struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewSchoolDTO(school domain.School) SchoolDTO {
	return SchoolDTO{
		ID:          school.ID.String(),
		OwnerID:     school.OwnerID.String(),
		Name:        school.Name,
		Description: school.Description,
	}
}
