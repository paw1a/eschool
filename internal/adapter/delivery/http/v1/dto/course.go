package dto

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

const (
	CourseDTODraft     = "draft"
	CourseDTOReady     = "ready"
	CourseDTOPublished = "published"
)

type CreateCourseDTO struct {
	Name     string   `json:"name" binding:"required" example:"Course name"`
	Level    null.Int `json:"level" binding:"required" swaggertype:"string" example:"5"`
	Price    null.Int `json:"price" binding:"required" swaggertype:"string" example:"3990"`
	Language string   `json:"language" binding:"required" example:"english"`
}

type UpdateCourseDTO struct {
	Name     null.String `json:"name" binding:"omitempty" swaggertype:"string" example:"Updated name"`
	Level    null.Int    `json:"level" binding:"omitempty" swaggertype:"string" example:"5"`
	Price    null.Int    `json:"price" binding:"omitempty" swaggertype:"string" example:"3990"`
	Language null.String `json:"language" binding:"omitempty" swaggertype:"string" example:"english"`
}

type CourseDTO struct {
	ID       string `json:"id" example:"30e18bc1-4354-4937-9a4d-03cf0b7027ca"`
	SchoolID string `json:"school_id" example:"30e18bc1-4354-4937-9a3b-03cf0b7034cc"`
	Name     string `json:"name" example:"Course name"`
	Level    int    `json:"level" example:"5"`
	Price    int64  `json:"price" example:"3990"`
	Language string `json:"language" example:"english"`
	Status   string `json:"status" example:"published"`
}

func NewCourseDTO(course domain.Course) CourseDTO {
	var status string
	switch course.Status {
	case domain.CourseDraft:
		status = CourseDTODraft
	case domain.CourseReady:
		status = CourseDTOReady
	case domain.CoursePublished:
		status = CourseDTOPublished
	}

	return CourseDTO{
		ID:       course.ID.String(),
		SchoolID: course.SchoolID.String(),
		Name:     course.Name,
		Level:    course.Level,
		Price:    course.Price,
		Language: course.Language,
		Status:   status,
	}
}
