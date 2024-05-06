package dto

import (
	"fmt"
	"github.com/guregu/null"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/pkg/errors"
)

const (
	LessonDTOTheory   = "theory"
	LessonDTOVideo    = "video"
	LessonDTOPractice = "practice"
)

type CreateLessonDTO struct {
	Title    string
	Type     string
	Score    null.Int
	Theory   null.String
	VideoUrl null.String
	Tests    []CreateTestDTO
}

func InputCreateLessonDTO(d *CreateLessonDTO) error {
	fmt.Print("Title: ")
	fmt.Scanln(&d.Title)

	var score int64
	fmt.Print("Score: ")
	fmt.Scanf("%d", &score)
	d.Score = null.IntFrom(score)

	fmt.Print("Type: ")
	fmt.Scanln(&d.Type)

	switch d.Type {
	case LessonDTOTheory:
		fmt.Print("Theory Markdown: ")
		var theory string
		fmt.Scanln(&theory)
		d.Theory = null.StringFrom(theory)
	case LessonDTOVideo:
		fmt.Print("Video URL: ")
		var url string
		fmt.Scanln(&url)
		d.VideoUrl = null.StringFrom(url)
	case LessonDTOPractice:
		var count int
		fmt.Print("Test count: ")
		fmt.Scanf("%d", &count)
		testDTOs := make([]CreateTestDTO, count)
		for i := 0; i < count; i++ {
			err := InputCreateTestDTO(&testDTOs[i])
			if err != nil {
				return err
			}
		}
		d.Tests = testDTOs
	default:
		return errors.New("invalid lesson type (theory, video, practice)")
	}

	fmt.Println()
	return nil
}

type CreateTestDTO struct {
	Task    string
	Options []string
	Answer  string
	Level   null.Int
	Score   null.Int
}

func InputCreateTestDTO(d *CreateTestDTO) error {
	fmt.Print("Task Markdown: ")
	fmt.Scanln(&d.Task)

	var count int
	fmt.Print("Test options count: ")
	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		return errors.New("invalid number")
	}

	options := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Option %d: ", i)
		fmt.Scanln(&options[i])
	}
	d.Options = options

	fmt.Print("Answer: ")
	fmt.Scanln(&d.Answer)

	var level int64
	fmt.Print("Level: ")
	_, err = fmt.Scanf("%d", &level)
	if err != nil {
		return errors.New("invalid number")
	}
	d.Level = null.IntFrom(level)

	var score int64
	fmt.Print("Score: ")
	_, err = fmt.Scanf("%d", &score)
	if err != nil {
		return errors.New("invalid number")
	}
	d.Score = null.IntFrom(score)

	fmt.Println()
	return nil
}

type UpdateLessonDTO struct {
	Title    null.String
	Score    null.Int
	Theory   null.String
	VideoUrl null.String
	Tests    []CreateTestDTO
}

type PassLessonDTO struct {
	LessonID  string
	PassTests []PassTestDTO
}

type PassTestDTO struct {
	TestID string
	Answer string
}

type LessonDTO struct {
	ID       string
	CourseID string
	Title    string
	Score    int
	Type     string

	TheoryUrl null.String
	VideoUrl  null.String
	Tests     []TestDTO
}

func PrintLessonDTO(d LessonDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Course ID: %s\n", d.CourseID)
	fmt.Printf("Title: %s\n", d.Title)
	fmt.Printf("Score: %d\n", d.Score)
	fmt.Printf("Type: %s\n", d.Type)
	switch d.Type {
	case LessonDTOTheory:
		fmt.Printf("TheoryUrl: %s\n", d.TheoryUrl.String)
	case LessonDTOVideo:
		fmt.Printf("VideoUrl: %s\n", d.VideoUrl.String)
	case LessonDTOPractice:
		fmt.Printf("Tests: %s\n", d.VideoUrl.String)
		fmt.Println()
		for _, test := range d.Tests {
			PrintTestDTO(test)
			fmt.Println()
		}
	}
}

type TestDTO struct {
	ID       string
	LessonID string
	TaskUrl  string
	Options  []string
	Answer   string
	Level    int
	Score    int
}

func PrintTestDTO(d TestDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Lesson ID: %s\n", d.LessonID)
	fmt.Printf("Task URL: %s\n", d.TaskUrl)
	fmt.Printf("Price: %v\n", d.Options)
	fmt.Printf("Answer: %s\n", d.Answer)
	fmt.Printf("Level: %d\n", d.Level)
	fmt.Printf("Score: %d\n", d.Score)
}

func NewLessonDTO(lesson domain.Lesson) LessonDTO {
	var lessonType string
	switch lesson.Type {
	case domain.TheoryLesson:
		lessonType = LessonDTOTheory
	case domain.VideoLesson:
		lessonType = LessonDTOVideo
	case domain.PracticeLesson:
		lessonType = LessonDTOPractice
	}

	tests := make([]TestDTO, len(lesson.Tests))
	if lesson.Type == domain.PracticeLesson {
		for i, test := range lesson.Tests {
			tests[i] = NewTestDTO(test)
		}
	}

	return LessonDTO{
		ID:        lesson.ID.String(),
		CourseID:  lesson.CourseID.String(),
		Title:     lesson.Title,
		Score:     lesson.Score,
		Type:      lessonType,
		TheoryUrl: lesson.TheoryUrl,
		VideoUrl:  lesson.VideoUrl,
		Tests:     tests,
	}
}

func NewTestDTO(test domain.Test) TestDTO {
	return TestDTO{
		ID:       test.ID.String(),
		LessonID: test.LessonID.String(),
		TaskUrl:  test.TaskUrl,
		Options:  test.Options,
		Answer:   test.Answer,
		Level:    test.Level,
		Score:    test.Score,
	}
}
