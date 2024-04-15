package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type StatisticsService struct {
	repo       port.IStatisticsRepository
	lessonRepo port.ILessonRepository
}

func NewStatisticsService(repo port.IStatisticsRepository,
	lessonRepo port.ILessonRepository) *StatisticsService {
	return &StatisticsService{
		repo:       repo,
		lessonRepo: lessonRepo,
	}
}

func (s *StatisticsService) FindLessonStat(ctx context.Context,
	userID, lessonID domain.ID) (domain.LessonStat, error) {
	return s.repo.FindLessonStat(ctx, userID, lessonID)
}

func (s *StatisticsService) CreateLessonStat(ctx context.Context,
	userID, lessonID domain.ID) error {
	lesson, err := s.lessonRepo.FindByID(ctx, lessonID)
	if err != nil {
		return err
	}

	var testStats []domain.TestStat
	if lesson.Type == domain.PracticeLesson {
		testStats = make([]domain.TestStat, len(lesson.Tests))
		for i, test := range lesson.Tests {
			testStats[i] = domain.TestStat{
				ID:     domain.NewID(),
				TestID: test.ID,
				UserID: userID,
				Score:  0,
			}
		}
	}

	return s.repo.CreateLessonStat(ctx, domain.LessonStat{
		ID:        domain.NewID(),
		LessonID:  lessonID,
		UserID:    userID,
		Score:     0,
		TestStats: testStats,
	})
}

func (s *StatisticsService) UpdateLessonStat(ctx context.Context, userID, lessonID domain.ID,
	param port.UpdateLessonStatParam) error {
	lessonStat, err := s.repo.FindLessonStat(ctx, userID, lessonID)
	if err != nil {
		return err
	}

	if param.Score.Valid {
		lessonStat.Score = int(param.Score.Int64)
	}

	if len(param.TestStats) != 0 {
		paramStats := param.TestStats
		for i := range lessonStat.TestStats {
			for j := range paramStats {
				if lessonStat.TestStats[i].TestID == paramStats[j].TestID {
					lessonStat.TestStats[i].Score = paramStats[j].Score
				}
			}
		}
	}

	return s.repo.UpdateLessonStat(ctx, lessonStat)
}
