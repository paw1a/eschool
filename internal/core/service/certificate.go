package service

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"time"
)

type CertificateService struct {
	repo       port.ICertificateRepository
	courseRepo port.ICourseRepository
	lessonRepo port.ILessonRepository
	statRepo   port.IStatRepository
}

func NewCertificateService(repo port.ICertificateRepository, courseRepo port.ICourseRepository,
	statRepo port.IStatRepository, lessonRepo port.ILessonRepository) *CertificateService {
	return &CertificateService{
		repo:       repo,
		courseRepo: courseRepo,
		statRepo:   statRepo,
		lessonRepo: lessonRepo,
	}
}

func (c *CertificateService) FindAll(ctx context.Context) ([]domain.Certificate, error) {
	return c.repo.FindAll(ctx)
}

func (c *CertificateService) FindByID(ctx context.Context, certificateID domain.ID) (domain.Certificate, error) {
	return c.repo.FindByID(ctx, certificateID)
}

func (c *CertificateService) FindUserCertificates(ctx context.Context,
	userID domain.ID) ([]domain.Certificate, error) {
	return c.repo.FindUserCertificates(ctx, userID)
}

func (c *CertificateService) FindCourseCertificate(ctx context.Context,
	courseID, userID domain.ID) (domain.Certificate, error) {
	return c.repo.FindUserCourseCertificate(ctx, courseID, userID)
}

func (c *CertificateService) calculateLessonsScores(ctx context.Context,
	userID domain.ID, lessons []domain.Lesson) (int, int, error) {
	var maxScore, score int
	for _, lesson := range lessons {
		lessonStat, err := c.statRepo.FindLessonStat(ctx, userID, lesson.ID)
		if err != nil {
			return 0, 0, err
		}

		switch lesson.Type {
		case domain.PracticeLesson:
			var maxLessonMark, lessonMark int
			for _, test := range lesson.Tests {
				for _, testStat := range lessonStat.TestStats {
					if test.ID == testStat.TestID {
						maxLessonMark += test.Score * test.Level
						lessonMark += testStat.Score * test.Level
					}
				}
			}
			maxScore += maxLessonMark + lesson.Score
			score += lessonMark + lesson.Score
		case domain.VideoLesson:
			fallthrough
		case domain.TheoryLesson:
			maxScore += lesson.Score
			score += lessonStat.Score
		}
	}

	return score, maxScore, nil
}

func (c *CertificateService) CreateCourseCertificate(ctx context.Context,
	userID, courseID domain.ID) (domain.Certificate, error) {
	course, err := c.courseRepo.FindByID(ctx, courseID)
	if err != nil {
		return domain.Certificate{}, err
	}

	lessons, err := c.lessonRepo.FindCourseLessons(ctx, courseID)
	if err != nil {
		return domain.Certificate{}, err
	}

	score, maxScore, err := c.calculateLessonsScores(ctx, userID, lessons)
	if err != nil {
		return domain.Certificate{}, err
	}

	percentage := float64(score) / float64(maxScore)
	var grade domain.CertificateGrade
	if percentage > 0.9 {
		grade = domain.GoldCertificate
	} else if percentage > 0.7 {
		grade = domain.SilverCertificate
	} else if percentage > 0.5 {
		grade = domain.BronzeCertificate
	} else {
		return domain.Certificate{}, errs.ErrCertificateCourseNotPassed
	}

	return c.repo.Create(ctx, domain.Certificate{
		ID:        domain.NewID(),
		UserID:    userID,
		CourseID:  courseID,
		Name:      fmt.Sprintf("Course \"%s\" certificate", course.Name),
		CreatedAt: time.Now(),
		Grade:     grade,
		Score:     score,
	})
}
