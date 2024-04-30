package port

import (
	"github.com/paw1a/eschool/internal/adapter/auth/jwt"
	"time"
)

type ISessionStorage interface {
	Get(refreshToken string) (jwt.AuthSession, error)
	Put(refreshToken string, session jwt.AuthSession, expireTime time.Duration) error
	Delete(refreshToken string) error
}
