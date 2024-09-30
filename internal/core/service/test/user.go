package test

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type CreateUserParamBuilder struct {
	param port.CreateUserParam
}

func NewCreateUserParamBuilder() *CreateUserParamBuilder {
	return &CreateUserParamBuilder{
		param: port.CreateUserParam{
			Name:      "name",
			Surname:   "surname",
			Phone:     null.StringFrom("+1234567890"),
			City:      null.String{},
			AvatarUrl: null.String{},
			Email:     "email@gmail.com",
			Password:  "password123",
		},
	}
}

func (b *CreateUserParamBuilder) WithName(name string) *CreateUserParamBuilder {
	b.param.Name = name
	return b
}

func (b *CreateUserParamBuilder) WithSurname(surname string) *CreateUserParamBuilder {
	b.param.Surname = surname
	return b
}

func (b *CreateUserParamBuilder) WithPhone(phone string) *CreateUserParamBuilder {
	b.param.Phone = null.StringFrom(phone)
	return b
}

func (b *CreateUserParamBuilder) WithCity(city string) *CreateUserParamBuilder {
	b.param.City = null.StringFrom(city)
	return b
}

func (b *CreateUserParamBuilder) WithAvatarUrl(avatarUrl string) *CreateUserParamBuilder {
	b.param.AvatarUrl = null.StringFrom(avatarUrl)
	return b
}

func (b *CreateUserParamBuilder) WithEmail(email string) *CreateUserParamBuilder {
	b.param.Email = email
	return b
}

func (b *CreateUserParamBuilder) WithPassword(password string) *CreateUserParamBuilder {
	b.param.Password = password
	return b
}

func (b *CreateUserParamBuilder) Build() port.CreateUserParam {
	return b.param
}

type UpdateUserParamBuilder struct {
	param port.UpdateUserParam
}

func NewUpdateUserParamBuilder() *UpdateUserParamBuilder {
	return &UpdateUserParamBuilder{
		param: port.UpdateUserParam{
			Name:      null.String{},
			Surname:   null.String{},
			Phone:     null.String{},
			City:      null.String{},
			AvatarUrl: null.String{},
		},
	}
}

func (b *UpdateUserParamBuilder) WithName(name string) *UpdateUserParamBuilder {
	b.param.Name = null.StringFrom(name)
	return b
}

func (b *UpdateUserParamBuilder) WithSurname(surname string) *UpdateUserParamBuilder {
	b.param.Surname = null.StringFrom(surname)
	return b
}

func (b *UpdateUserParamBuilder) WithPhone(phone string) *UpdateUserParamBuilder {
	b.param.Phone = null.StringFrom(phone)
	return b
}

func (b *UpdateUserParamBuilder) WithCity(city string) *UpdateUserParamBuilder {
	b.param.City = null.StringFrom(city)
	return b
}

func (b *UpdateUserParamBuilder) WithAvatarUrl(avatarUrl string) *UpdateUserParamBuilder {
	b.param.AvatarUrl = null.StringFrom(avatarUrl)
	return b
}

func (b *UpdateUserParamBuilder) Build() port.UpdateUserParam {
	return b.param
}

type UserBuilder struct {
	user domain.User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: domain.User{
			ID:        domain.NewID(),
			Name:      "John",
			Surname:   "Doe",
			Phone:     null.StringFrom("+123456789"),
			City:      null.StringFrom("Berlin"),
			AvatarUrl: null.StringFrom("https://example.com/avatar"),
			Email:     "john.doe@example.com",
			Password:  "password",
		},
	}
}

func (b *UserBuilder) WithID(id domain.ID) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilder) Build() domain.User {
	return b.user
}
