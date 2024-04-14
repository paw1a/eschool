package port

import (
	"github.com/guregu/null"
)

type CreateTheoryParam struct {
	Title  string
	Score  int
	Theory string
}

type CreateVideoParam struct {
	Title    string
	Score    int
	VideoUrl string
}

type CreatePracticeParam struct {
	Title string
	Score int
	Tests []CreateTestParam
}

type CreateTestParam struct {
	Task    string
	Options []string
	Answer  string
	Level   int
	Score   int
}

type UpdateTheoryParam struct {
	Title  null.String
	Score  null.Int
	Theory null.String
}

type UpdateVideoParam struct {
	Title    null.String
	Score    null.Int
	VideoUrl null.String
}

type UpdatePracticeParam struct {
	Title null.String
	Score null.Int
	Tests []UpdateTestParam
}

type UpdateTestParam struct {
	Task    string
	Options []string
	Answer  string
	Level   int
	Score   int
}
