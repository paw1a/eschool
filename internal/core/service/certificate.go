package service

import (
	"context"
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
	domainErr "github.com/paw1a/eschool/internal/core/errors"
	"github.com/paw1a/eschool/internal/core/port"
	"time"
)

type CertificateService struct {
	repo       port.ICertificateRepository
	courseRepo port.ICourseRepository
	lessonRepo port.ILessonRepository
	statRepo   port.IStatisticsRepository
}

func NewCertificateService(repo port.ICertificateRepository, courseRepo port.ICourseRepository,
	statRepo port.IStatisticsRepository, lessonRepo port.ILessonRepository) *CertificateService {
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

	var maxMark, mark int
	for _, lesson := range lessons {
		switch lesson.Type {
		case domain.PracticeLesson:
			tests, err := c.lessonRepo.FindLessonTests(ctx, lesson.ID)
			if err != nil {
				return domain.Certificate{}, err
			}

			var maxLessonMark, lessonMark int
			for _, test := range tests {
				maxLessonMark += test.Mark * test.Level
				testStat, err := c.statRepo.FindUserTestStat(ctx, userID, test.ID)
				if err != nil {
					return domain.Certificate{}, err
				}
				lessonMark += testStat.Mark * test.Level
			}
			maxMark += maxLessonMark + lesson.Mark
			mark += lessonMark + lesson.Mark
		case domain.VideoLesson:
			fallthrough
		case domain.TheoryLesson:
			lessonStat, err := c.statRepo.FindUserLessonStat(ctx, userID, lesson.ID)
			if err != nil {
				return domain.Certificate{}, err
			}
			maxMark += lesson.Mark
			mark += lessonStat.Mark
		}
	}

	percentage := float64(mark) / float64(maxMark)
	var grade domain.CertificateGrade
	if percentage > 0.9 {
		grade = domain.GoldCertificate
	} else if percentage > 0.7 {
		grade = domain.SilverCertificate
	} else if percentage > 0.5 {
		grade = domain.BronzeCertificate
	} else {
		return domain.Certificate{}, domainErr.ErrCertificateCourseNotPassed
	}

	return c.repo.Create(ctx, domain.Certificate{
		ID:    domain.NewID(),
		Name:  fmt.Sprintf("Сертификат о прохождении курса \"%s\"", course.Name),
		Date:  time.Now(),
		Grade: grade,
	})
}
