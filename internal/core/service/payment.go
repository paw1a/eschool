package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"net/url"
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

func (p *PaymentService) GetCoursePaymentUrl(ctx context.Context, userID, courseID domain.ID) (url.URL, error) {
	course, err := p.courseRepo.FindByID(ctx, courseID)
	if err != nil {
		return url.URL{}, err
	}

	isStudent, err := p.courseRepo.IsCourseStudent(ctx, userID, courseID)
	if err != nil {
		return url.URL{}, err
	}

	if isStudent {
		return url.URL{}, errs.ErrUserIsAlreadyCourseStudent
	}

	return p.gateway.GetPaymentUrl(ctx, domain.PaymentPayload{
		UserID:   userID,
		CourseID: courseID,
		PaySum:   course.Price,
	})
}

func (p *PaymentService) ProcessCoursePayment(ctx context.Context,
	key string, paid int64) (domain.PaymentPayload, error) {
	payload, err := p.gateway.ProcessPayment(ctx, key)
	if err != nil {
		return domain.PaymentPayload{}, err
	}

	if paid < payload.PaySum {
		return domain.PaymentPayload{}, errs.ErrInvalidPaymentSum
	}

	return payload, nil
}
