package dto

import "github.com/guregu/null"

type CreateSchoolDTO struct {
	Description string
}

type UpdateSchoolDTO struct {
	Description null.String
}
