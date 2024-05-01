package dto

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type CreateSchoolDTO struct {
	Name        string
	Description string
}

type UpdateSchoolDTO struct {
	Description null.String
}

type AddTeacherDTO struct {
	TeacherID string `json:"teacher_id" binding:"required"`
}

type SchoolDTO struct {
	ID          string `json:"id" binding:"required"`
	OwnerID     string `json:"owner_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func NewSchoolDTO(school domain.School) SchoolDTO {
	return SchoolDTO{
		ID:          school.ID.String(),
		OwnerID:     school.OwnerID.String(),
		Name:        school.Name,
		Description: school.Description,
	}
}
