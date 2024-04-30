package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paw1a/eschool/internal/adapter/auth/port"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"
	"time"
)

type Config struct {
	Secret           string
	AccessTokenTime  int64
	RefreshTokenTime int64
}

type AuthProvider struct {
	cfg            *Config
	sessionStorage port.ISessionStorage
}

func NewAuthProvider(cfg *Config, sessionStorage port.ISessionStorage) *AuthProvider {
	return &AuthProvider{
		cfg:            cfg,
		sessionStorage: sessionStorage,
	}
}

func (p *AuthProvider) CreateJWTSession(payload domain.AuthPayload,
	fingerprint string) (domain.AuthDetails, error) {
	accessExpTime := time.Minute * time.Duration(p.cfg.AccessTokenTime)
	accessExp := time.Now().Add(accessExpTime).Unix()
	claims := jwt.MapClaims{
		"exp":    accessExp,
		"userID": payload.UserID.String(),
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := unsignedToken.SignedString([]byte(p.cfg.Secret))
	if err != nil {
		return domain.AuthDetails{}, err
	}

	refreshToken := uuid.NewV4().String()
	refreshExpTime := time.Minute * time.Duration(p.cfg.RefreshTokenTime)
	refreshExp := time.Now().Add(refreshExpTime).Unix()

	session := port.AuthSession{
		RefreshToken: refreshToken,
		RefreshExp:   refreshExp,
		Fingerprint:  fingerprint,
		Payload:      payload,
	}

	err = p.sessionStorage.Put(refreshToken, session, refreshExpTime)
	if err != nil {
		return domain.AuthDetails{}, err
	}

	return domain.AuthDetails{
		AccessToken:  domain.Token(accessToken),
		RefreshToken: domain.Token(refreshToken),
	}, nil
}

func (p *AuthProvider) RefreshJWTSession(refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	session, err := p.sessionStorage.Get(refreshToken.String())
	if err != nil {
		return domain.AuthDetails{}, err
	}

	err = p.sessionStorage.Delete(refreshToken.String())
	if err != nil {
		return domain.AuthDetails{}, err
	}

	if session.Fingerprint != fingerprint {
		return domain.AuthDetails{}, errs.ErrInvalidFingerprint
	}

	return p.CreateJWTSession(session.Payload, fingerprint)
}

func (p *AuthProvider) DeleteJWTSession(refreshToken domain.Token) error {
	return p.sessionStorage.Delete(refreshToken.String())
}

func (p *AuthProvider) VerifyJWTToken(accessToken domain.Token) (domain.AuthPayload, error) {
	token, err := jwt.Parse(accessToken.String(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrInvalidTokenSignMethod
		}
		return []byte(p.cfg.Secret), nil
	})
	if err != nil {
		return domain.AuthPayload{}, errors.Wrap(errs.ErrInvalidToken, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		payload := domain.AuthPayload{
			UserID: domain.ID(claims["userID"].(string)),
		}
		return payload, nil
	}

	return domain.AuthPayload{}, errs.ErrInvalidTokenClaims
}
