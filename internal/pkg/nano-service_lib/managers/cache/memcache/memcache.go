
package memcache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/rainycape/memcache"
	"github.com/sirupsen/logrus"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/cache"
)

var instance *MemCache

type MemCache struct {
	cache   *memcache.Client
	options *Options
}

func NewMemCache(options *Options) *MemCache {
	if instance == nil {
		instance = createInstance(options)
	}

	return instance
}

func createInstance(options *Options) *MemCache {
	m, err := memcache.New(options.ConnectionString())
	if err != nil {
		log.Panic(err)
	}

	options.IsShared = true

	return &MemCache{
		cache:   m,
		options: options,
	}
}

func (c *MemCache) Get(k string) (interface{}, bool) {
	keyHash := cache.Hash(k)

	item, err := c.cache.Get(keyHash)
	if err != nil {
		if c.options.Debug {
			logrus.Warn(err)
		}

		return nil, false
	}

	return item, false
}

func (c *MemCache) GetInto(k string, obj interface{}) bool {
	keyHash := cache.Hash(k)

	item, err := c.cache.Get(keyHash)
	if err != nil {
		if c.options.Debug {
			logrus.Warn(err)
		}

		return false
	}

	err = json.Unmarshal(item.Value, obj)
	if err != nil {
		if c.options.Debug {
			logrus.Warn(err)
		}

		return false
	}

	return true
}

func (c *MemCache) Set(k string, x interface{}, d time.Duration) {
	marshal, err := json.Marshal(x)
	if err != nil {
		if c.options.Debug {
			logrus.Warn(err)
		}
	}

	keyHash := cache.Hash(k)

	item := &memcache.Item{
		Key:        keyHash,
		Value:      marshal,
		Expiration: int32(d.Seconds()),
	}

	err = c.cache.Set(item)
	if err != nil {
		if c.options.Debug {
			logrus.Warn(err)
		}
	}
}

func (c *MemCache) Delete(k string) {
	err := c.cache.Delete(cache.Hash(k))
	if err != nil {
		if c.options.Debug {
			logrus.Warn(err)
		}
	}
}

func (c *MemCache) FlushAll() {
	if !c.options.IsShared {
		err := c.cache.Flush(0)
		if err != nil {
			if c.options.Debug {
				logrus.Warn(err)
			}
		}
	}
}

func (c *MemCache) Increment(k string, n int64) error {
	_, err := c.cache.Increment(cache.Hash(k), uint64(n))
	if err != nil {
		if c.options.Debug {
			logrus.Warn(err)
		}

		return err
	}

	return nil
}

func (c *MemCache) ItemCount() int {
	return 0
}
