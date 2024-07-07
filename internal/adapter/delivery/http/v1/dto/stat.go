package dto

import (
	"github.com/paw1a/eschool/internal/core/domain"
)

type LessonStatDTO struct {
	ID        string        `json:"id"`
	LessonID  string        `json:"lesson_id"`
	UserID    string        `json:"user_id"`
	Score     int           `json:"score"`
	TestStats []TestStatDTO `json:"tests"`
}

type TestStatDTO struct {
	ID     string `json:"id"`
	TestID string `json:"test_id"`
	UserID string `json:"user_id"`
	Score  int    `json:"score"`
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
