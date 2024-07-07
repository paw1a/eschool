package yoomoney

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"github.com/google/uuid"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"net/url"
	"slices"
	"strconv"
)

type Config struct {
	Scheme string
	Host   string
	Path   string
	Wallet string
}

type PaymentYookassaGateway struct {
	config *Config
}

func NewPaymentGateway(config *Config) *PaymentYookassaGateway {
	return &PaymentYookassaGateway{
		config: config,
	}
}

func (g *PaymentYookassaGateway) GetPaymentUrl(ctx context.Context, payload domain.PaymentPayload) (url.URL, error) {
	userUUID, _ := uuid.Parse(payload.UserID.String())
	userUUIDBytes, _ := userUUID.MarshalBinary()
	courseUUID, _ := uuid.Parse(payload.CourseID.String())
	courseUUIDBytes, _ := courseUUID.MarshalBinary()

	paySumBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(paySumBytes, uint64(payload.PaySum))

	dataBytes := slices.Concat(userUUIDBytes, courseUUIDBytes, paySumBytes)
	encodedData := base64.StdEncoding.EncodeToString(dataBytes)
	formParams := url.Values{
		"sum":           {strconv.FormatInt(payload.PaySum, 10)},
		"receiver":      {g.config.Wallet},
		"quickpay-form": {"donate"},
		"label":         {encodedData},
	}

	return url.URL{
		Scheme:   g.config.Scheme,
		Host:     g.config.Host,
		Path:     g.config.Path,
		RawQuery: formParams.Encode(),
	}, nil
}

func (g *PaymentYookassaGateway) ProcessPayment(ctx context.Context, key string) (domain.PaymentPayload, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed
	}

	var userID, courseID uuid.UUID
	err = userID.UnmarshalBinary(dataBytes[:16])
	if err != nil {
		return domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed
	}

	err = courseID.UnmarshalBinary(dataBytes[16:32])
	if err != nil {
		return domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed
	}

	paySum := binary.LittleEndian.Uint64(dataBytes[32:])
	if err != nil {
		return domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed
	}

	return domain.PaymentPayload{
		UserID:   domain.ID(userID.String()),
		CourseID: domain.ID(courseID.String()),
		PaySum:   int64(paySum),
	}, nil
}
