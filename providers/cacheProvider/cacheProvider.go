package cacheProvider

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Caches struct {
	Cache map[string]*Result
	sync.Mutex
}
type AppCache struct {
	Client Caches
}

type Result struct {
	value             []byte
	defaultExpiration time.Duration
}

func InitCache() *Caches {
	return &Caches{Cache: make(map[string]*Result)}
}

func (r *Caches) Set(key string, data interface{}, expirationTime time.Duration) error {
	r.Lock()
	defer r.Unlock()
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var value Result
	value.value = bytes
	value.defaultExpiration = expirationTime
	r.Cache[key] = &value
	return nil

}

func (r *Caches) Get(key string) ([]byte, error) {
	Result, exist := r.Cache[key]
	fmt.Println(r.Cache[key])
	if !exist {
		return nil, nil
	}

	//var data interface{}
	//_ = json.Unmarshal(Result.value, data)

	return Result.value, nil
}
