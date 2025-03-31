package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
)

var ErrSessionNotFound error = errors.New("session not found: ")

type SessionRepo struct {
	conn *database.Cache
}

func NewSessionRepo(client *database.Cache) *SessionRepo {
	return &SessionRepo{
		conn: client,
	}
}

func (sr *SessionRepo) SetSession(key string, value any, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sr.conn.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepo) DeleteSesion(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sr.conn.Del(ctx, keys...).Err(); err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepo) GetSession(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stringCmd := sr.conn.Get(ctx, key)
	switch err := stringCmd.Err(); {
	case err == redis.Nil:
		return "", fmt.Errorf("%w: %s", ErrSessionNotFound, key)
	case err != nil:
		return "", err
	}

	return stringCmd.Val(), nil
}
