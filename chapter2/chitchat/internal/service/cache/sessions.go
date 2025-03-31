package cache

import (
	"context"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository/cache"
)

type SessionService struct {
	sessionRepo *cache.SessionRepo
}

func NewSessionService(sr *cache.SessionRepo) *SessionService {
	return &SessionService{
		sessionRepo: sr,
	}
}

func (sr *SessionService) SetSession(key string, value any, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return sr.sessionRepo.SetSession(ctx, key, value, expiration)
}

func (sr *SessionService) DeleteSession(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return sr.sessionRepo.DeleteSesion(ctx, keys...)
}

func (sr *SessionService) GetSession(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return sr.sessionRepo.GetSession(ctx, key)
}
