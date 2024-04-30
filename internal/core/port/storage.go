package port

import (
	"context"
	"github.com/paw1a/eschool/internal/core/domain"
)

type IObjectStorage interface {
	SaveFile(ctx context.Context, file domain.File) (domain.Url, error)
}
