package service

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/zap"
	"time"
)

type CertificateService struct {
	repo       port.ICertificateRepository
	courseRepo port.ICourseRepository
	lessonRepo port.ILessonRepository
	statRepo   port.IStatRepository
	logger     *zap.Logger
}

func NewCertificateService(repo port.ICertificateRepository, courseRepo port.ICourseRepository,
	statRepo port.IStatRepository, lessonRepo port.ILessonRepository, logger *zap.Logger) *CertificateService {
	return &CertificateService{
		repo:       repo,
		courseRepo: courseRepo,
		statRepo:   statRepo,
		lessonRepo: lessonRepo,
		logger:     logger,
	}
}

func (c *CertificateService) FindAll(ctx context.Context) ([]domain.Certificate, error) {
	certificates, err := c.repo.FindAll(ctx)
	if err != nil {
		c.logger.Error("failed to find all certificates", zap.Error(err))
		return nil, err
	}
	return certificates, nil
}

func (c *CertificateService) FindByID(ctx context.Context, certificateID domain.ID) (domain.Certificate, error) {
	certificate, err := c.repo.FindByID(ctx, certificateID)
	if err != nil {
		c.logger.Error("failed to find certificate by id", zap.Error(err),
			zap.String("certificateID", certificateID.String()))
		return domain.Certificate{}, err
	}
	return certificate, nil
}

func (c *CertificateService) FindUserCertificates(ctx context.Context,
	userID domain.ID) ([]domain.Certificate, error) {
	certificates, err := c.repo.FindUserCertificates(ctx, userID)
	if err != nil {
		c.logger.Error("failed to find certificates by user id", zap.Error(err),
			zap.String("userID", userID.String()))
		return nil, err
	}
	return certificates, nil
}

func (c *CertificateService) FindCourseCertificate(ctx context.Context,
	courseID, userID domain.ID) (domain.Certificate, error) {
	certificate, err := c.repo.FindUserCourseCertificate(ctx, courseID, userID)
	if err != nil {
		c.logger.Error("failed to find course certificate", zap.Error(err),
			zap.String("courseID", courseID.String()), zap.String("userID", userID.String()))
		return domain.Certificate{}, err
	}
	return certificate, err
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
		c.logger.Error("failed to create course certificate", zap.Error(err),
			zap.String("courseID", courseID.String()), zap.String("userID", userID.String()))
		return domain.Certificate{}, err
	}

	lessons, err := c.lessonRepo.FindCourseLessons(ctx, courseID)
	if err != nil {
		c.logger.Error("failed to find course lessons", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return domain.Certificate{}, err
	}

	score, maxScore, err := c.calculateLessonsScores(ctx, userID, lessons)
	if err != nil {
		c.logger.Error("failed to calculate certificate scores", zap.Error(err),
			zap.String("courseID", courseID.String()), zap.String("userID", userID.String()))
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

	certificate, err := c.repo.Create(ctx, domain.Certificate{
		ID:        domain.NewID(),
		UserID:    userID,
		CourseID:  courseID,
		Name:      fmt.Sprintf("Course \"%s\" certificate", course.Name),
		CreatedAt: time.Now(),
		Grade:     grade,
		Score:     score,
	})
	if err != nil {
		c.logger.Error("failed to create course certificate", zap.Error(err),
			zap.String("courseID", courseID.String()), zap.String("userID", userID.String()))
		return domain.Certificate{}, err
	}

	c.logger.Info("course certificate is created",
		zap.String("certificateID", certificate.ID.String()), zap.String("userID", userID.String()))
	return certificate, nil
}
