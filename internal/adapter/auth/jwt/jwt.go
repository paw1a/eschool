package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/paw1a/eschool/internal/adapter/auth/port"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/twinj/uuid"
	"time"
)

type Config struct {
	Secret           string
	AccessTokenTime  int64
	RefreshTokenTime int64
}

type AuthSession struct {
	RefreshToken string
	RefreshExp   int64
	Fingerprint  string
	Payload      domain.AuthPayload
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

	session := AuthSession{
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
		return domain.AuthDetails{}, errors.New("invalid client fingerprint")
	}

	return p.CreateJWTSession(session.Payload, fingerprint)
}

func (p *AuthProvider) DeleteJWTSession(refreshToken domain.Token) error {
	return p.sessionStorage.Delete(refreshToken.String())
}

func (p *AuthProvider) VerifyJWTToken(accessToken domain.Token) (domain.AuthPayload, error) {
	token, err := jwt.Parse(accessToken.String(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.cfg.Secret), nil
	})
	if err != nil {
		return domain.AuthPayload{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		payload := domain.AuthPayload{
			UserID: domain.ID(claims["userID"].(string)),
		}
		return payload, nil
	}

	return domain.AuthPayload{}, errors.New("token or claims are invalid")
}
