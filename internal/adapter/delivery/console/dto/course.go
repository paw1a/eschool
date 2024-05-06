package dto

import (
	"fmt"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

const (
	CourseDTODraft     = "draft"
	CourseDTOReady     = "ready"
	CourseDTOPublished = "published"
)

type CreateCourseDTO struct {
	Name     string
	Level    null.Int
	Price    null.Int
	Language string
}

type UpdateCourseDTO struct {
	Name     null.String
	Level    null.Int
	Price    null.Int
	Language null.String
}

type CourseDTO struct {
	ID       string
	SchoolID string
	Name     string
	Level    int
	Price    int64
	Language string
	Status   string
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

func PrintCourseDTO(d CourseDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("School ID: %s\n", d.SchoolID)
	fmt.Printf("Name: %s\n", d.Name)
	fmt.Printf("Price: %d\n", d.Price)
	fmt.Printf("Level: %d\n", d.Level)
	fmt.Printf("Language: %s\n", d.Language)
	fmt.Printf("Status: %s\n", d.Status)
}
