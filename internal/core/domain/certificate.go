package domain

import "time"

type CertificateGrade int

const (
	Bronze CertificateGrade = iota
	Silver
	Gold
)

type Certificate struct {
	ID    int64
	Name  string
	Date  time.Time
	Grade CertificateGrade
}
