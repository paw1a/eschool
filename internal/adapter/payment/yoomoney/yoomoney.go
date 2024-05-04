package yoomoney

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"net/url"
	"strconv"
	"strings"
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
	data := fmt.Sprintf("%s;%s;%d", payload.UserID, payload.CourseID, payload.PaySum)
	dataEncrypted := base64.StdEncoding.EncodeToString([]byte(data))

	formParams := url.Values{
		"sum":           {strconv.FormatInt(payload.PaySum, 10)},
		"receiver":      {g.config.Wallet},
		"quickpay-form": {"donate"},
		"label":         {dataEncrypted},
	}

	return url.URL{
		Scheme:   g.config.Scheme,
		Host:     g.config.Host,
		Path:     g.config.Path,
		RawQuery: formParams.Encode(),
	}, nil
}

func (g *PaymentYookassaGateway) ProcessPayment(ctx context.Context, key string) (domain.PaymentPayload, error) {
	data, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed
	}

	dataSplit := strings.Split(string(data), ";")
	if len(dataSplit) != 3 {
		return domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed
	}

	paySum, err := strconv.ParseInt(dataSplit[2], 10, 64)
	if err != nil {
		return domain.PaymentPayload{}, errs.ErrDecodePaymentKeyFailed
	}

	return domain.PaymentPayload{
		UserID:   domain.ID(dataSplit[0]),
		CourseID: domain.ID(dataSplit[1]),
		PaySum:   paySum,
	}, nil
}
