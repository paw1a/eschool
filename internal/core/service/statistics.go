package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type StatisticsService struct {
	repo port.IStatisticsRepository
}

func NewStatisticsService(repo port.IStatisticsRepository) *StatisticsService {
	return &StatisticsService{
		repo: repo,
	}
}

func (s *StatisticsService) FindUserLessonStat(ctx context.Context, userID, lessonID domain.ID) (domain.LessonStat, error) {
	return s.repo.FindUserLessonStat(ctx, userID, lessonID)
}

func (s *StatisticsService) FindUserTestStat(ctx context.Context, userID, testID domain.ID) (domain.TestStat, error) {
	return s.repo.FindUserTestStat(ctx, userID, testID)
}

func (s *StatisticsService) CreateUserLessonStat(ctx context.Context, userID, lessonID domain.ID) error {
	return s.repo.CreateLessonStat(ctx, domain.LessonStat{
		ID:       domain.NewID(),
		LessonID: lessonID,
		UserID:   userID,
		Score:    0,
	})
}

func (s *StatisticsService) CreateUserTestStat(ctx context.Context, userID, testID domain.ID) error {
	return s.repo.CreateTestStat(ctx, domain.TestStat{
		ID:     domain.NewID(),
		TestID: testID,
		UserID: userID,
		Score:  0,
	})
}

func (s *StatisticsService) UpdateUserLessonStat(ctx context.Context, userID, lessonID domain.ID,
	param port.UpdateLessonStatParam) error {
	stat, err := s.repo.FindUserLessonStat(ctx, userID, lessonID)
	if err != nil {
		return err
	}

	if param.Mark.Valid {
		stat.Score = int(param.Mark.Int64)
	}

	return s.repo.UpdateUserLessonStat(ctx, stat)
}

func (s *StatisticsService) UpdateUserTestStat(ctx context.Context, userID, testID domain.ID,
	param port.UpdateTestStatParam) error {
	stat, err := s.repo.FindUserTestStat(ctx, userID, testID)
	if err != nil {
		return err
	}

	if param.Mark.Valid {
		stat.Score = int(param.Mark.Int64)
	}

	return s.repo.UpdateUserTestStat(ctx, stat)
}
