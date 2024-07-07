package dto

import (
	"fmt"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

type UpdateUserDTO struct {
	Name      null.String
	Surname   null.String
	Phone     null.String
	City      null.String
	AvatarUrl null.String
}

func InputUpdateUserDTO(d *UpdateUserDTO) error {
	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	if name != "" {
		d.Name = null.StringFrom(name)
	}

	var surname string
	fmt.Print("Surname: ")
	fmt.Scanln(&surname)
	if surname != "" {
		d.Surname = null.StringFrom(surname)
	}

	fmt.Print("Phone: ")
	var phone string
	fmt.Scanln(&phone)
	if phone != "" {
		e164Regex := `^\+[1-9]\d{1,14}$`
		re := regexp.MustCompile(e164Regex)
		phone = strings.ReplaceAll(phone, " ", "")
		if re.Find([]byte(phone)) == nil {
			return errors.New("invalid phone number format")
		}
		d.Phone = null.StringFrom(phone)
	}

	fmt.Print("City: ")
	var city string
	fmt.Scanln(&city)
	if city != "" {
		d.City = null.StringFrom(city)
	}

	fmt.Println()

	return nil
}

type UserInfoDTO struct {
	Name    string
	Surname string
}

type UserDTO struct {
	ID        string
	Name      string
	Surname   string
	Email     string
	Phone     string
	City      string
	AvatarUrl string
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

func PrintUserDTO(d UserDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Name: %s\n", d.Name)
	fmt.Printf("Surname: %s\n", d.Surname)
	fmt.Printf("Email: %s\n", d.Email)
	fmt.Printf("Phone: %s\n", d.Phone)
	fmt.Printf("City: %s\n", d.City)
}
