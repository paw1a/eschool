package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/port"
)

type PaymentService struct {
	gateway    port.IPaymentGateway
	courseRepo port.ICourseRepository
}

func NewPaymentService(gateway port.IPaymentGateway, courseRepo port.ICourseRepository) *PaymentService {
	return &PaymentService{
		gateway:    gateway,
		courseRepo: courseRepo,
	}
}

func (p *PaymentService) PayCourse(ctx context.Context, userID, courseID domain.ID) error {
	course, err := p.courseRepo.FindByID(ctx, courseID)
	if err != nil {
		return err
	}
	return p.gateway.Pay(ctx, courseID.String(), course.Price)
}
