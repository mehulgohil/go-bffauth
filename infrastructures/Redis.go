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

func (r *RedisClient) DeleteKey(key string) error {
	_, err := r.Client.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) HSetKeyValue(key string, value map[string]interface{}, exp time.Duration) error {
	_, err := r.Client.HSet(context.Background(), key, value).Result()
	if err != nil {
		return err
	}

	_, err = r.Client.Expire(context.Background(), key, exp).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) HGetKeyValue(key string) (map[string]string, error) {
	result, err := r.Client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}
