package domain

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
	"strings"
)

type LessonType int

const (
	TheoryLesson LessonType = iota
	PracticeLesson
	VideoLesson
)

type Lesson struct {
	ID       ID
	CourseID ID
	Title    string
	Score    int
	Type     LessonType

	TheoryUrl null.String
	VideoUrl  null.String
	Tests     []Test
}

type Test struct {
	ID       ID
	LessonID ID
	TaskUrl  string
	Options  []string
	Answer   string
	Level    int
	Score    int
}

type LessonStat struct {
	ID        ID
	LessonID  ID
	UserID    ID
	Score     int
	TestStats []TestStat
}

type TestStat struct {
	ID     ID
	TestID ID
	UserID ID
	Score  int
}

func (l *Lesson) Validate() error {
	var errList []error
	if l.Score <= 0 {
		errList = append(errList, errs.ErrCourseLessonInvalidScore)
	}
	switch l.Type {
	case PracticeLesson:
		if len(l.Tests) == 0 {
			errList = append(errList, errs.ErrCoursePracticeLessonEmptyTests)
		}

		for _, test := range l.Tests {
			if test.TaskUrl == "" {
				errList = append(errList, errs.ErrCoursePracticeLessonEmptyTestTaskUrl)
			}
			if len(test.Options) == 0 {
				errList = append(errList, errs.ErrCoursePracticeLessonEmptyTestOptions)
			}
			if test.Level < 0 {
				errList = append(errList, errs.ErrLessonTestInvalidLevel)
			}
			if test.Score <= 0 {
				errList = append(errList, errs.ErrCoursePracticeLessonInvalidTestScore)
			}
		}
	case TheoryLesson:
		if !l.TheoryUrl.Valid {
			errList = append(errList, errs.ErrCourseTheoryLessonEmptyUrl)
		}
	case VideoLesson:
		if !l.VideoUrl.Valid {
			errList = append(errList, errs.ErrCourseVideoLessonEmptyUrl)
		}
	}

	if errList != nil {
		errStrings := make([]string, len(errList))
		for i := range errList {
			errStrings[i] = errList[i].Error()
		}
		return errors.New(strings.Join(errStrings, "\n"))
	}

	return nil
}
