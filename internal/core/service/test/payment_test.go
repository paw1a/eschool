package test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/paw1a/eschool/internal/core/service"
	"github.com/paw1a/eschool/internal/core/service/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"net/url"
	"testing"
)

type PaymentSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (s *PaymentSuite) BeforeEach(t provider.T) {
	loggerBuilder := zap.NewDevelopmentConfig()
	loggerBuilder.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	s.logger, _ = loggerBuilder.Build()
}

// GetCoursePaymentUrl Suite
type PaymentGetCoursePaymentUrlSuite struct {
	PaymentSuite
}

func PaymentGetCoursePaymentUrlSuccessRepositoryMock(gateway *mocks.PaymentGateway,
	courseRepository *mocks.CourseRepository) {
	gateway.
		On("GetPaymentUrl", context.Background(), mock.Anything).
		Return(url.URL{}, nil)
	courseRepository.
		On("FindByID", context.Background(), mock.Anything).
		Return(NewCourseBuilder().Build(), nil)
	courseRepository.
		On("IsCourseStudent", context.Background(), mock.Anything, mock.Anything).
		Return(false, nil)
}

func (s *PaymentGetCoursePaymentUrlSuite) TestGetCoursePaymentUrl_Success(t provider.T) {
	t.Parallel()
	t.Title("Get course payment url success")
	gateway := mocks.NewPaymentGateway(t)
	courseRepository := mocks.NewCourseRepository(t)
	paymentService := service.NewPaymentService(gateway, courseRepository, s.logger)
	PaymentGetCoursePaymentUrlSuccessRepositoryMock(gateway, courseRepository)
	_, err := paymentService.GetCoursePaymentUrl(context.Background(), domain.NewID(), domain.NewID())
	t.Assert().Nil(err)
}

func PaymentGetCoursePaymentUrlFailureRepositoryMock(gateway *mocks.PaymentGateway,
	courseRepository *mocks.CourseRepository) {
	courseRepository.
		On("FindByID", context.Background(), mock.Anything).
		Return(NewCourseBuilder().Build(), nil)
	courseRepository.
		On("IsCourseStudent", context.Background(), mock.Anything, mock.Anything).
		Return(false, errs.ErrUserIsAlreadyCourseStudent)
}

func (s *PaymentGetCoursePaymentUrlSuite) TestGetCoursePaymentUrl_Failure(t provider.T) {
	t.Parallel()
	t.Title("Get course payment url failure")
	gateway := mocks.NewPaymentGateway(t)
	courseRepository := mocks.NewCourseRepository(t)
	paymentService := service.NewPaymentService(gateway, courseRepository, s.logger)
	PaymentGetCoursePaymentUrlFailureRepositoryMock(gateway, courseRepository)
	_, err := paymentService.GetCoursePaymentUrl(context.Background(), domain.NewID(), domain.NewID())
	t.Assert().ErrorIs(err, errs.ErrUserIsAlreadyCourseStudent)
}

func TestPaymentGetCoursePaymentUrlSuite(t *testing.T) {
	suite.RunNamedSuite(t, "GetCoursePaymentUrl", new(PaymentGetCoursePaymentUrlSuite))
}

// ProcessCoursePayment Suite
type PaymentProcessCoursePaymentSuite struct {
	PaymentSuite
}

func PaymentProcessCoursePaymentSuccessRepositoryMock(gateway *mocks.PaymentGateway) {
	gateway.
		On("ProcessPayment", context.Background(), mock.Anything).
		Return(domain.PaymentPayload{}, nil)
}

func (s *PaymentProcessCoursePaymentSuite) TestProcessCoursePayment_Success(t provider.T) {
	t.Parallel()
	t.Title("process payment success")
	gateway := mocks.NewPaymentGateway(t)
	courseRepository := mocks.NewCourseRepository(t)
	paymentService := service.NewPaymentService(gateway, courseRepository, s.logger)
	PaymentProcessCoursePaymentSuccessRepositoryMock(gateway)
	_, err := paymentService.ProcessCoursePayment(context.Background(), "key", 1000)
	t.Assert().Nil(err)
}

func PaymentProcessCoursePaymentFailureRepositoryMock(gateway *mocks.PaymentGateway) {
	gateway.
		On("ProcessPayment", context.Background(), mock.Anything).
		Return(domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed)
}

func (s *PaymentProcessCoursePaymentSuite) TestProcessCoursePayment_Failure(t provider.T) {
	t.Parallel()
	t.Title("Process payment failure")
	gateway := mocks.NewPaymentGateway(t)
	courseRepository := mocks.NewCourseRepository(t)
	paymentService := service.NewPaymentService(gateway, courseRepository, s.logger)
	PaymentProcessCoursePaymentFailureRepositoryMock(gateway)
	_, err := paymentService.ProcessCoursePayment(context.Background(), "key", 1000)
	t.Assert().ErrorIs(err, errs.ErrDecodePaymentKeyFailed)
}

func TestPaymentProcessCoursePaymentSuite(t *testing.T) {
	suite.RunNamedSuite(t, "ProcessCoursePayment", new(PaymentProcessCoursePaymentSuite))
}
