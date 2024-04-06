package port

import "context"

type IPaymentGateway interface {
	Pay(ctx context.Context, key string, price int64) error
}
