package yoomoney

import "context"

type PaymentYookassaGateway struct {
}

func NewPaymentGateway() *PaymentYookassaGateway {
	return &PaymentYookassaGateway{}
}

func (g *PaymentYookassaGateway) Pay(ctx context.Context, key string, price int64) error {
	return nil
}
