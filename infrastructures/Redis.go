package infrastructures

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}

func (r *RedisClient) SetKeyValue(key, value string, exp time.Duration) error {
	_, err := r.Client.Set(context.Background(), key, value, exp).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) GetKeyValue(key string) (string, error) {
	val, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
