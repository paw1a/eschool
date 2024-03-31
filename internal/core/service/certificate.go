package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type CertificateService struct {
	repo port.ICertificateRepository
}

func NewCertificateService(repo port.ICertificateRepository) *CertificateService {
	return &CertificateService{
		repo: repo,
	}
}

func (c *CertificateService) FindAll(ctx context.Context) ([]domain.Certificate, error) {
	return c.repo.FindAll(ctx)
}

func (c *CertificateService) FindByID(ctx context.Context, certificateID int64) (domain.Certificate, error) {
	return c.repo.FindByID(ctx, certificateID)
}

func (c *CertificateService) FindUserCertificates(ctx context.Context,
	userID int64) ([]domain.Certificate, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CertificateService) CreateCourseCertificate(ctx context.Context,
	userID, courseID int64) (domain.Certificate, error) {
	//TODO implement me
	panic("implement me")
}
