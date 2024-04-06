package port

import "github.com/guregu/null"

type CreateSchoolParam struct {
	Description string
}

type UpdateSchoolParam struct {
	Description null.String
}
