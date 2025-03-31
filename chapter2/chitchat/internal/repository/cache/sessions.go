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

func (sr *SessionRepo) SetSession(ctx context.Context, key string, value any, expiration time.Duration) error {
	if err := sr.conn.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepo) DeleteSesion(ctx context.Context, keys ...string) error {
	if err := sr.conn.Del(ctx, keys...).Err(); err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepo) GetSession(ctx context.Context, key string) (string, error) {
	stringCmd := sr.conn.Get(ctx, key)
	switch err := stringCmd.Err(); {
	case err == redis.Nil:
		return "", fmt.Errorf("%w: %s", ErrSessionNotFound, key)
	case err != nil:
		return "", err
	}

	return stringCmd.Val(), nil
}
