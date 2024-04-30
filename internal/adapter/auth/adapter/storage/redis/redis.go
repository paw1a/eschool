package redis

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/paw1a/eschool/internal/adapter/auth/jwt"
	"time"
)

type SessionStorage struct {
	redisClient *redis.Client
}

func NewSessionStorage(redisClient *redis.Client) *SessionStorage {
	return &SessionStorage{redisClient: redisClient}
}

func (s *SessionStorage) Get(refreshToken string) (jwt.AuthSession, error) {
	sessionJson, err := s.redisClient.Get(refreshToken).Bytes()
	if err != nil {
		return jwt.AuthSession{}, err
	}

	var session jwt.AuthSession
	err = json.Unmarshal(sessionJson, &session)
	if err != nil {
		return jwt.AuthSession{}, err
	}

	return session, nil
}

func (s *SessionStorage) Put(refreshToken string, session jwt.AuthSession,
	expireTime time.Duration) error {
	sessionJson, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return s.redisClient.Set(refreshToken, sessionJson, expireTime).Err()
}

func (s *SessionStorage) Delete(refreshToken string) error {
	return s.redisClient.Del(refreshToken).Err()
}
