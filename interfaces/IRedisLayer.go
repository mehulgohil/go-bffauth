package interfaces

import "time"

type IRedisLayer interface {
	SetKeyValue(key, value string, exp time.Duration) error
	GetKeyValue(key string) (string, error)
	DeleteKey(key string) error
	HSetKeyValue(key string, value map[string]interface{}, exp time.Duration) error
	HGetKeyValue(key string) (map[string]string, error)
}
