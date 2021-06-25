package cache

import (
	"fmt"
)

type CacheType string

const (
	CacheTypeMemory CacheType = "memory"
)

func New(cacheType CacheType) (Cache, error) {
	var c Cache
	switch cacheType {
	case CacheTypeMemory:
		c = newMemoryCache()
	default:
		return nil, fmt.Errorf("unsupported cache type: %s", string(cacheType))
	}
	return c, nil
}

type Cache interface {
	Set(key string, value []byte) error
	Get(key string) (value []byte, err error)
	Del(key string) error
	GetStat() Stat
}
