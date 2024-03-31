package repository

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
)

type CertificateRepository struct {
}

func (c *CertificateRepository) FindAll(ctx context.Context) ([]domain.Certificate, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CertificateRepository) FindByID(ctx context.Context, certificateID int64) (domain.Certificate, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CertificateRepository) FindUserCertificates(ctx context.Context, userID int64) ([]domain.Certificate, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CertificateRepository) CreateCourseCertificate(ctx context.Context, userID, courseID int64) (domain.Certificate, error) {
	//TODO implement me
	panic("implement me")
}
