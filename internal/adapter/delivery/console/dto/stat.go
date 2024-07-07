package dto

import (
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
)

type LessonStatDTO struct {
	ID        string        `json:"id"`
	LessonID  string        `json:"lesson_id"`
	UserID    string        `json:"user_id"`
	Score     int           `json:"score"`
	TestStats []TestStatDTO `json:"tests"`
}

func PrintLessonStatDTO(d LessonStatDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Lesson ID: %s\n", d.LessonID)
	fmt.Printf("User ID: %s\n", d.UserID)
	fmt.Printf("Score: %d\n", d.Score)
	fmt.Println()
	fmt.Println("Tests progress")
	for _, test := range d.TestStats {
		PrintTestStatDTO(test)
		fmt.Println()
	}
}

type TestStatDTO struct {
	ID     string `json:"id"`
	TestID string `json:"test_id"`
	UserID string `json:"user_id"`
	Score  int    `json:"score"`
}

func PrintTestStatDTO(d TestStatDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Test ID: %s\n", d.TestID)
	fmt.Printf("User ID: %s\n", d.UserID)
	fmt.Printf("Score: %d\n", d.Score)
}

func NewLessonStatDTO(lessonStat domain.LessonStat) LessonStatDTO {
	testStats := make([]TestStatDTO, len(lessonStat.TestStats))
	for i, testStat := range lessonStat.TestStats {
		testStats[i] = NewTestStatDTO(testStat)
	}

	return LessonStatDTO{
		ID:        lessonStat.ID.String(),
		UserID:    lessonStat.UserID.String(),
		LessonID:  lessonStat.LessonID.String(),
		Score:     lessonStat.Score,
		TestStats: testStats,
	}
}

func NewTestStatDTO(testStat domain.TestStat) TestStatDTO {
	return TestStatDTO{
		ID:     testStat.ID.String(),
		UserID: testStat.UserID.String(),
		TestID: testStat.TestID.String(),
		Score:  testStat.Score,
	}
}
