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
	Name        string   `json:"name" binding:"required" example:"Course name"`
	Description string   `json:"description" binding:"required" example:"Course description"`
	Level       null.Int `json:"level" binding:"required" swaggertype:"string" example:"5"`
	Price       null.Int `json:"price" binding:"required" swaggertype:"string" example:"3990"`
	Language    string   `json:"language" binding:"required" example:"english"`
	ImageUrl    string   `json:"image_url" swaggertype:"string" example:"image.io/image.png"`
}

type UpdateCourseDTO struct {
	Name        null.String `json:"name" binding:"omitempty" swaggertype:"string" example:"Updated name"`
	Description null.String `json:"description" binding:"omitempty" swaggertype:"string" example:"Updated description"`
	Level       null.Int    `json:"level" binding:"omitempty" swaggertype:"string" example:"5"`
	Price       null.Int    `json:"price" binding:"omitempty" swaggertype:"string" example:"3990"`
	Language    null.String `json:"language" binding:"omitempty" swaggertype:"string" example:"english"`
	ImageUrl    null.String `json:"image_url" swaggertype:"string" example:"image.io/image.png"`
}

type CourseDTO struct {
	ID          string `json:"id" example:"30e18bc1-4354-4937-9a4d-03cf0b7027ca"`
	SchoolID    string `json:"school_id" example:"30e18bc1-4354-4937-9a3b-03cf0b7034cc"`
	Name        string `json:"name" example:"Course name"`
	Description string `json:"description" example:"Course description"`
	Level       int    `json:"level" example:"5"`
	Price       int64  `json:"price" example:"3990"`
	Language    string `json:"language" example:"english"`
	ImageUrl    string `json:"image_url" swaggertype:"string" example:"image.io/image.png"`
	Status      string `json:"status" example:"published"`
}

type CourseBriefDTO struct {
	ID          string           `json:"id" example:"30e18bc1-4354-4937-9a4d-03cf0b7027ca"`
	Name        string           `json:"name" example:"Course name"`
	Description string           `json:"description" example:"Course description"`
	Price       int64            `json:"price" example:"3990"`
	Language    string           `json:"language" example:"english"`
	ImageUrl    string           `json:"image_url" swaggertype:"string" example:"image.io/image.png"`
	Lessons     []LessonBriefDTO `json:"lessons"`
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
		ID:          course.ID.String(),
		SchoolID:    course.SchoolID.String(),
		Name:        course.Name,
		Description: course.Description,
		Level:       course.Level,
		Price:       course.Price,
		Language:    course.Language,
		ImageUrl:    course.ImageUrl,
		Status:      status,
	}
}

func NewCourseBriefDTO(course domain.Course) CourseBriefDTO {
	return CourseBriefDTO{
		ID:          course.ID.String(),
		Name:        course.Name,
		Description: course.Description,
		Price:       course.Price,
		Language:    course.Language,
		ImageUrl:    course.ImageUrl,
	}
}
