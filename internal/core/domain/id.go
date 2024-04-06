package domain

import "github.com/twinj/uuid"

type ID string

func (id ID) String() string {
	return string(id)
}

func NewID() ID {
	return ID(uuid.NewV4().String())
}

func RandomID() ID {
	return NewID()
}

func ParseID(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return ID(uuid.UUID{}.String()), err
	}
	return ID(id.String()), nil
}
