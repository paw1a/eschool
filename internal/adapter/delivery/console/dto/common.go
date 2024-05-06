package dto

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/mail"
)

func InputEmail(email *string) error {
	fmt.Print("Email: ")
	var input string
	fmt.Scanln(&input)
	_, err := mail.ParseAddress(input)
	if err != nil {
		return errors.New("invalid email format")
	}
	*email = input
	return nil
}

func InputID(id *domain.ID, idOwner string) error {
	fmt.Printf("%s ID: ", cases.Title(language.Und, cases.NoLower).String(idOwner))
	var input string
	fmt.Scanln(&input)
	_, err := uuid.Parse(input)
	if err != nil {
		return errors.New("invalid uuid format")
	}
	*id = domain.ID(input)
	return nil
}
