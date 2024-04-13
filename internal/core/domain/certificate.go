package domain

import "time"

type CertificateGrade int

const (
	BronzeCertificate CertificateGrade = iota
	SilverCertificate
	GoldCertificate
)

type Certificate struct {
	ID        ID
	CourseID  ID
	UserID    ID
	Name      string
	Score     int
	CreatedAt time.Time
	Grade     CertificateGrade
}
