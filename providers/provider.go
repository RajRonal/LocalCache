package providers

import "time"

type CacheInt interface {
	Set(key string, data interface{}, expirationTime time.Time) (interface{}, error)
	Get(key string) ([]byte, error)
}
