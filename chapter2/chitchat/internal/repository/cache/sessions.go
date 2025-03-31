package cache

import (
	"context"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
)

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

// func (sr *SessionRepo) GetSession(data string) (string, error) {
//
// }
