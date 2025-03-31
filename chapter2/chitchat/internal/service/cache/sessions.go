package cache

import (
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
	return sr.sessionRepo.SetSession(key, value, expiration)
}

func (sr *SessionService) DeleteSession(keys ...string) error {
	return sr.sessionRepo.DeleteSesion(keys...)
}

func (sr *SessionService) GetSession(key string) (string, error) {
	return sr.sessionRepo.GetSession(key)
}
