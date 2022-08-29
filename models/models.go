package models

import (
	"github.com/gofrs/uuid"
	"time"
)

//type CacheInt interface {
//	Set(key string, data interface{}, expiration time.Duration) error
//	Get(key string) ([]byte, error)
//}

type ToDo struct {
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	ExpiryTime time.Time `json:"expiryTime"`
}
type UserId struct {
	ID uuid.UUID `json:"id"`
}

//type Cache struct {
//	Cache map[string]*Result
//}
//
//type Result struct {
//	value []byte
//	err   error
//}
//type Func func() ([]byte, error)

//type Item struct {
//	Object     interface{}
//	Expiration int64
//}
