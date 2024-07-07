package dto

import (
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
	"time"
)

const (
	BronzeCertificate = "bronze"
	SilverCertificate = "silver"
	GoldCertificate   = "gold"
)

type CertificateDTO struct {
	ID        string
	CourseID  string
	UserID    string
	Name      string
	Score     int
	CreatedAt time.Time
	Grade     string
}

func PrintCertificateDTO(d CertificateDTO) {
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Course ID: %s\n", d.CourseID)
	fmt.Printf("User ID: %s\n", d.UserID)
	fmt.Printf("Name: %s\n", d.Name)
	fmt.Printf("Score: %d\n", d.Score)
	fmt.Printf("Date: %s\n", d.CreatedAt.String())
	fmt.Printf("Grade: %s\n", d.Grade)
}

func NewCertificateDTO(certificate domain.Certificate) CertificateDTO {
	var grade string
	switch certificate.Grade {
	case domain.BronzeCertificate:
		grade = BronzeCertificate
	case domain.SilverCertificate:
		grade = SilverCertificate
	case domain.GoldCertificate:
		grade = GoldCertificate
	}

	return CertificateDTO{
		ID:        certificate.ID.String(),
		CourseID:  certificate.CourseID.String(),
		UserID:    certificate.UserID.String(),
		Name:      certificate.Name,
		Score:     certificate.Score,
		CreatedAt: certificate.CreatedAt,
		Grade:     grade,
	}
}
