package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository/cache"
)

var (
	ErrMarshallSession   error = errors.New("failed to marshal session data")
	ErrUnMarshallSession error = errors.New("failed to unmarshal session data")
	ErrSetSession        error = errors.New("failed to set session data")
)

type SessionService struct {
	sessionRepo *cache.SessionRepo
}

func NewSessionService(sr *cache.SessionRepo) *SessionService {
	return &SessionService{
		sessionRepo: sr,
	}
}

func (sr *SessionService) SetSession(key string, session models.Session, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sessionStr, err := json.Marshal(session)
	if err != nil {
		return ErrMarshallSession
	}

	return sr.sessionRepo.SetSession(ctx, key, sessionStr, expiration)
}

func (sr *SessionService) DeleteSession(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return sr.sessionRepo.DeleteSesion(ctx, keys...)
}

func (sr *SessionService) GetSession(key string) (models.Session, error) {
	var session models.Session

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sessionStr, err := sr.sessionRepo.GetSession(ctx, key)
	if err != nil {
		return session, err
	}

	if err := json.Unmarshal([]byte(sessionStr), &session); err != nil {
		return session, ErrUnMarshallSession
	}

	return session, nil
}
