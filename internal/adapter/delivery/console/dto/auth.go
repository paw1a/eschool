package dto

import (
	"fmt"
	"github.com/guregu/null"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

type SignUpDTO struct {
	Name      string
	Surname   string
	Email     string
	Password  string
	Phone     null.String
	City      null.String
	AvatarUrl null.String
}

func InputSignUpDTO(d *SignUpDTO) error {
	fmt.Print("Name: ")
	fmt.Scanln(&d.Name)

	fmt.Print("Surname: ")
	fmt.Scanln(&d.Surname)

	err := InputEmail(&d.Email)
	if err != nil {
		return err
	}

	fmt.Print("Password: ")
	fmt.Scanln(&d.Password)

	fmt.Print("Phone: ")
	var phone string
	fmt.Scanln(&phone)
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	phone = strings.ReplaceAll(phone, " ", "")

	if re.Find([]byte(phone)) == nil {
		return errors.New("invalid phone number format")
	}
	d.Phone.String = phone

	fmt.Print("City: ")
	var city string
	fmt.Scanln(&city)
	d.City.String = city

	fmt.Println()

	return nil
}

type SignInDTO struct {
	Email       string
	Password    string
	Fingerprint string
}

func InputSignInDTO(d *SignInDTO) error {
	err := InputEmail(&d.Email)
	if err != nil {
		return err
	}
	fmt.Print("Password: ")
	fmt.Scanln(&d.Password)
	d.Fingerprint = "secret"
	return nil
}
