package test

import (
	"github.com/paw1a/eschool/internal/core/domain"
)

type SchoolBuilder struct {
	school domain.School
}

func NewSchoolBuilder() *SchoolBuilder {
	return &SchoolBuilder{
		school: domain.School{
			ID:          domain.NewID(),
			OwnerID:     domain.NewID(),
			Name:        "default school",
			Description: "default description",
		},
	}
}

func (b *SchoolBuilder) WithID(id domain.ID) *SchoolBuilder {
	b.school.ID = id
	return b
}

func (b *SchoolBuilder) WithOwnerID(ownerID domain.ID) *SchoolBuilder {
	b.school.OwnerID = ownerID
	return b
}

func (b *SchoolBuilder) WithName(name string) *SchoolBuilder {
	b.school.Name = name
	return b
}

func (b *SchoolBuilder) WithDescription(description string) *SchoolBuilder {
	b.school.Description = description
	return b
}

func (b *SchoolBuilder) Build() domain.School {
	return b.school
}
