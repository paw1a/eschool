package dto

import (
	"bufio"
	"fmt"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/pkg/errors"
	"os"
	"strings"
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

func InputCreateCourseDTO(d *CreateCourseDTO) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Course name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("empty school name")
	}
	d.Name = name

	var level int64
	fmt.Print("Level: ")
	fmt.Scanf("%d", &level)
	d.Level = null.IntFrom(level)

	var price int64
	fmt.Print("Price: ")
	fmt.Scanf("%d", &price)
	d.Price = null.IntFrom(price)

	var language string
	fmt.Print("Language: ")
	fmt.Scanln(&language)
	if language == "" {
		return errors.New("empty language field")
	}
	d.Language = language

	fmt.Println()
	return nil
}

type UpdateCourseDTO struct {
	Name     null.String
	Level    null.Int
	Price    null.Int
	Language null.String
}

func InputUpdateCourseDTO(d *UpdateCourseDTO) error {
	var name string
	fmt.Print("Course name: ")
	fmt.Scanln(&name)
	if name != "" {
		d.Name = null.StringFrom(name)
	}

	var level int64
	fmt.Print("Level: ")
	_, err := fmt.Scanf("%d", &level)
	if err != nil {
		return errors.New("invalid number")
	}
	d.Level = null.IntFrom(level)

	var price int64
	fmt.Print("Price: ")
	_, err = fmt.Scanf("%d", &price)
	if err != nil {
		return errors.New("invalid number")
	}
	d.Price = null.IntFrom(price)

	var language string
	fmt.Print("Language: ")
	fmt.Scanln(&language)
	if language != "" {
		d.Language = null.StringFrom(language)
	}

	fmt.Println()
	return nil
}

type CourseDTO struct {
	ID       string
	SchoolID string
	Name     string
	Level    int
	Price    int64
	Language string
	Status   string
	Rating   float64
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
		Rating:   course.Rating,
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
	fmt.Printf("Avg rating: %.2f\n", d.Rating)
}
