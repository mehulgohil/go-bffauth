package config

import (
	"github.com/mehulgohil/go-bffauth.git/infrastructures"
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	redisObj  *RedisHandler
	redisOnce sync.Once
)

type IRedisHandler interface {
	InitRedisConnection()
}

type RedisHandler struct {
	RedisClient *infrastructures.RedisClient
}

func (r *RedisHandler) InitRedisConnection() {
	r.RedisClient = &infrastructures.RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     EnvVariables.RedisHost,
			Password: EnvVariables.RedisPassword, // no password set
			DB:       0,                          // use default DB
		}),
	}
}

func Redis() IRedisHandler {
	if redisObj == nil {
		redisOnce.Do(func() {
			redisObj = &RedisHandler{}
		})
	}
	return redisObj
}
