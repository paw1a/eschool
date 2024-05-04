package port

import "github.com/guregu/null"

type CreateSchoolParam struct {
	Name        string
	Description string
}

type UpdateSchoolParam struct {
	Description null.String
}
