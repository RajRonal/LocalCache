package utilities

import (
	"sync"
	"time"
)

//
//type AppCache struct {
//  Client *cache.Cache
////	//	Client  *models.Cache
////	//Client *Caches
////}

type Caches struct {
	Cache map[string]*Result
	sync.Mutex
}

type Result struct {
	value             []byte
	defaultExpiration time.Duration
}

type AppCache struct {
	Client Caches
}

//type Func func() ([]byte, error)

//func InitCache() *Caches {
//	return &Caches{Cache: make(map[string]*Result)}
//}

//
//func (r *AppCache) Set(key string, data interface{}, expiration time.Duration) error {
//	//r.Client
//	//defer r.l.Unlock()
//	//r.Client.
//	//r.Client.
//
//	bytes, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	r.Client.Set(key, bytes, expiration)
//	return nil
//}
//
//func (r *AppCache) Get(key string) ([]byte, error) {
//	Result, exist := r.Client.Get(key)
//	if !exist {
//		return nil, nil
//	}
//
//	resultByte, ok := Result.([]byte)
//	if !ok {
//		return nil, errors.New("Format is not array of bytes")
//	}
//
//	return resultByte, nil
//}

//func NewCache() *Caches {
//	return &Caches{Cache: make(map[string]*Result)}
//}

//func (r *Caches) Get(key string) ([]byte, error) {
//	r.Lock()
//	defer r.Unlock()
//	res, ok := r.Cache[key]
//	if !ok {
//		res = &Result{}
//		res.value, res.err = nil, nil
//		r.Cache[key] = res
//	}
//	return res.value, res.err
//}
