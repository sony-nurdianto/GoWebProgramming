package database

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheInterface interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Close() error
}

type Cache struct {
	client *redis.Client
}

func NewCache() (*Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "MySecretPass",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Println("Redis Connected: ", pong)

	return &Cache{client: client}, nil
}

func (c *Cache) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(ctx, key, value, expiration)
}

func (c *Cache) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.client.Get(ctx, key)
}

func (c *Cache) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return c.client.Del(ctx, keys...)
}

func (c *Cache) Close() error {
	return c.client.Close()
}
