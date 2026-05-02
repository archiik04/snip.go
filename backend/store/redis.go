package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(redisURL string) *RedisStore {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	return &RedisStore{client: redis.NewClient(opt)}
}

func (r *RedisStore) Set(ctx context.Context, code, url string, ttl time.Duration) error {
	return r.client.Set(ctx, "url:"+code, url, ttl).Err()
}

func (r *RedisStore) Get(ctx context.Context, code string) (string, error) {
	return r.client.Get(ctx, "url:"+code).Result()
}

func (r *RedisStore) IncrClicks(ctx context.Context, code string) error {
	return r.client.Incr(ctx, "clicks:"+code).Err()
}

func (r *RedisStore) GetClicks(ctx context.Context, code string) (int64, error) {
	return r.client.Get(ctx, "clicks:"+code).Int64()
}
