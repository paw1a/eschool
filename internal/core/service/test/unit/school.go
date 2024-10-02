package unit

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
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

type CreateSchoolParamBuilder struct {
	param port.CreateSchoolParam
}

func NewCreateSchoolParamBuilder() *CreateSchoolParamBuilder {
	return &CreateSchoolParamBuilder{
		param: port.CreateSchoolParam{
			Name:        "default name",
			Description: "default description",
		},
	}
}

func (b *CreateSchoolParamBuilder) WithName(name string) *CreateSchoolParamBuilder {
	b.param.Name = name
	return b
}

func (b *CreateSchoolParamBuilder) WithDescription(description string) *CreateSchoolParamBuilder {
	b.param.Description = description
	return b
}

func (b *CreateSchoolParamBuilder) Build() port.CreateSchoolParam {
	return b.param
}

type UpdateSchoolParamBuilder struct {
	param port.UpdateSchoolParam
}

func NewUpdateSchoolParamBuilder() *UpdateSchoolParamBuilder {
	return &UpdateSchoolParamBuilder{
		param: port.UpdateSchoolParam{
			Description: null.NewString("default description", true),
		},
	}
}

func (b *UpdateSchoolParamBuilder) WithDescription(description null.String) *UpdateSchoolParamBuilder {
	b.param.Description = description
	return b
}

func (b *UpdateSchoolParamBuilder) Build() port.UpdateSchoolParam {
	return b.param
}
