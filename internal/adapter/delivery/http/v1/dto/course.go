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
	Name     string `json:"name" binding:"required"`
	Level    int    `json:"level" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	Language string `json:"language" binding:"required"`
}

type UpdateCourseDTO struct {
	Name     null.String `json:"name" binding:"omitempty"`
	Level    null.Int    `json:"level" binding:"omitempty"`
	Price    null.Int    `json:"price" binding:"omitempty"`
	Language null.String `json:"language" binding:"omitempty"`
}

type CourseDTO struct {
	ID       string `json:"id" binding:"required"`
	SchoolID string `json:"school_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Level    int    `json:"level" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	Language string `json:"language" binding:"required"`
	Status   string `json:"status" binding:"required"`
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
