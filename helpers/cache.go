package helpers

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Delete(key string)
}

func NewCache() Cache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &cacheWrapper{c}
}

type cacheWrapper struct {
	*cache.Cache
}

func (c *cacheWrapper) Get(key string) (interface{}, bool) {
	return c.Cache.Get(key)
}

func (c *cacheWrapper) Set(key string, value interface{}) {
	c.Cache.Set(key, value, cache.DefaultExpiration)
}

func (c *cacheWrapper) Delete(key string) {
	c.Cache.Delete(key)
}
