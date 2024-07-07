package port

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
)

type UpdateLessonStatParam struct {
	Score     null.Int
	TestStats []UpdateTestStatParam
}

type UpdateTestStatParam struct {
	TestID domain.ID
	Score  int
}
