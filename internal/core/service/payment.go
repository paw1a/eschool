package service

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/port"
	"go.uber.org/zap"
	"net/url"
)

type PaymentService struct {
	gateway    port.IPaymentGateway
	courseRepo port.ICourseRepository
	logger     *zap.Logger
}

func NewPaymentService(gateway port.IPaymentGateway, courseRepo port.ICourseRepository,
	logger *zap.Logger) *PaymentService {
	return &PaymentService{
		gateway:    gateway,
		courseRepo: courseRepo,
		logger:     logger,
	}
}

func (p *PaymentService) GetCoursePaymentUrl(ctx context.Context, userID, courseID domain.ID) (url.URL, error) {
	course, err := p.courseRepo.FindByID(ctx, courseID)
	if err != nil {
		p.logger.Error("failed to find course by id", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return url.URL{}, err
	}

	isStudent, err := p.courseRepo.IsCourseStudent(ctx, userID, courseID)
	if err != nil {
		p.logger.Error("failed to check if user is a course student", zap.Error(err),
			zap.String("courseID", courseID.String()), zap.String("userID", userID.String()))
		return url.URL{}, err
	}

	if isStudent {
		return url.URL{}, errs.ErrUserIsAlreadyCourseStudent
	}

	link, err := p.gateway.GetPaymentUrl(ctx, domain.PaymentPayload{
		UserID:   userID,
		CourseID: courseID,
		PaySum:   course.Price,
	})
	if err != nil {
		p.logger.Error("failed to get payment link", zap.Error(err),
			zap.String("courseID", courseID.String()))
		return url.URL{}, err
	}

	p.logger.Info("payment link is generated successfully",
		zap.String("url", link.String()), zap.String("userID", userID.String()),
		zap.String("courseID", courseID.String()))
	return link, nil
}

func (p *PaymentService) ProcessCoursePayment(ctx context.Context,
	key string, paid int64) (domain.PaymentPayload, error) {
	payload, err := p.gateway.ProcessPayment(ctx, key)
	if err != nil {
		p.logger.Error("failed to process payment", zap.Error(err),
			zap.String("key", key), zap.Int64("paid sum", paid))
		return domain.PaymentPayload{}, err
	}

	if paid < payload.PaySum {
		p.logger.Error("failed to process payment, payment sum is less then expected",
			zap.Error(err), zap.String("key", key), zap.Int64("paid sum", paid))
		return domain.PaymentPayload{}, errs.ErrInvalidPaymentSum
	}

	p.logger.Info("payment is processed successfully",
		zap.String("key", key), zap.Int64("paid sum", paid),
		zap.String("courseID", payload.CourseID.String()),
		zap.String("userID", payload.UserID.String()))
	return payload, nil
}
