package entity

import (
	"github.com/google/uuid"
	"github.com/paw1a/eschool/internal/core/domain"
)

type PgLessonStat struct {
	ID       uuid.UUID
	LessonID uuid.UUID
	UserID   uuid.UUID
	Score    int
}

type PgTestStat struct {
	ID     uuid.UUID
	TestID uuid.UUID
	UserID uuid.UUID
	Score  int
}

func (s *PgLessonStat) ToDomain() domain.LessonStat {
	return domain.LessonStat{
		ID:        domain.ID(s.ID.String()),
		LessonID:  domain.ID(s.LessonID.String()),
		UserID:    domain.ID(s.UserID.String()),
		Score:     s.Score,
		TestStats: nil,
	}
}

func NewPgLessonStat(stat domain.LessonStat) PgLessonStat {
	id, _ := uuid.Parse(stat.ID.String())
	lessonID, _ := uuid.Parse(stat.LessonID.String())
	userID, _ := uuid.Parse(stat.UserID.String())
	return PgLessonStat{
		ID:       id,
		LessonID: lessonID,
		UserID:   userID,
		Score:    stat.Score,
	}
}

func (s *PgTestStat) ToDomain() domain.TestStat {
	return domain.TestStat{
		ID:     domain.ID(s.ID.String()),
		TestID: domain.ID(s.TestID.String()),
		UserID: domain.ID(s.UserID.String()),
		Score:  s.Score,
	}
}

func NewPgTestStat(stat domain.TestStat) PgTestStat {
	id, _ := uuid.Parse(stat.ID.String())
	testID, _ := uuid.Parse(stat.TestID.String())
	userID, _ := uuid.Parse(stat.UserID.String())
	return PgTestStat{
		ID:     id,
		TestID: testID,
		UserID: userID,
		Score:  stat.Score,
	}
}
