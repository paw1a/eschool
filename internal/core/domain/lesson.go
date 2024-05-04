package domain

import (
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/errs"
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

func (l *Lesson) Validate() error {
	if l.Score <= 0 {
		return errs.ErrCourseLessonInvalidScore
	}
	switch l.Type {
	case PracticeLesson:
		if len(l.Tests) == 0 {
			return errs.ErrCoursePracticeLessonEmptyTests
		}

		for _, test := range l.Tests {
			if test.TaskUrl == "" {
				return errs.ErrCoursePracticeLessonEmptyTestTaskUrl
			}
			if len(test.Options) == 0 {
				return errs.ErrCoursePracticeLessonEmptyTestOptions
			}
			if test.Level < 0 {
				return errs.ErrCoursePracticeLessonInvalidTestLevel
			}
			if test.Score <= 0 {
				return errs.ErrCoursePracticeLessonInvalidTestScore
			}
		}
	case TheoryLesson:
		if !l.TheoryUrl.Valid {
			return errs.ErrCourseTheoryLessonEmptyUrl
		}
	case VideoLesson:
		if !l.VideoUrl.Valid {
			return errs.ErrCourseVideoLessonEmptyUrl
		}
	}

	return nil
}
