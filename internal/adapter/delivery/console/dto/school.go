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

type CreateSchoolDTO struct {
	Name        string
	Description null.String
}

func InputCreateSchoolDTO(d *CreateSchoolDTO) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("School name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("empty school name")
	}
	d.Name = name

	fmt.Print("School description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	if description != "" {
		d.Description = null.StringFrom(description)
	}

	fmt.Println()
	return nil
}

type UpdateSchoolDTO struct {
	Description null.String
}

type SchoolDTO struct {
	ID          string
	OwnerID     string
	Name        string
	Description string
}

func PrintSchoolDTO(d SchoolDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Owner ID: %s\n", d.OwnerID)
	fmt.Printf("Name: %s\n", d.Name)
	fmt.Printf("Description: %s\n", d.Description)
}

func NewSchoolDTO(school domain.School) SchoolDTO {
	return SchoolDTO{
		ID:          school.ID.String(),
		OwnerID:     school.OwnerID.String(),
		Name:        school.Name,
		Description: school.Description,
	}
}
