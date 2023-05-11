package interfaces

import "time"

type IRedisLayer interface {
	SetKeyValue(key, value string, exp time.Duration) error
	GetKeyValue(key string) (string, error)
}
