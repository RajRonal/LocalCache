package cacheProvider

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Caches struct {
	Cache map[string]*Result
	lx    sync.Mutex
}

type AppCache struct {
	Client Caches
}

type Result struct {
	value             []byte
	defaultExpiration time.Time
}

func InitCache() *Caches {
	return &Caches{Cache: make(map[string]*Result)}
}

func (r *Caches) Set(key string, data interface{}, expirationTime time.Time) (interface{}, error) {
	var value Result
	var result interface{}
	r.lx.Lock()
	defer r.lx.Unlock()
	bytes, err := json.Marshal(data)
	if err != nil {
		return value, err
	}

	value.value = bytes
	value.defaultExpiration = expirationTime
	r.Cache[key] = &value
	//fmt.Println("before post query", value.defaultExpiration)
	err = json.Unmarshal(value.value, &result)
	if err != nil {
		return value, err
	}

	return result, nil

}

func (r *Caches) Get(key string) ([]byte, error) {
	result, exist := r.Cache[key]
	var err = errors.New("Expired")
	fmt.Println(err)
	if result.defaultExpiration.Before(time.Now()) {
		logrus.Error("Out of time")
		return nil, err
	}

	if !exist {
		return nil, nil
	}

	return result.value, nil
}
