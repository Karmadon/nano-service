
package memory_cache

import (
	"time"

	internal "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/cache/cache_common"
)

var cacheInstance *InternalCache

type InternalCache struct {
	cache *internal.Cache
	Debug bool
}

func (c *InternalCache) GetInto(k string, obj interface{}) bool {
	panic("implement me")
}

func NewInternalCache(debug bool, cache *internal.Cache) *InternalCache {
	if cacheInstance == nil {
		if cache != nil {
			cacheInstance = &InternalCache{cache: cache, Debug: debug}
		} else {
			cacheInstance = &InternalCache{cache: internal.New(cache_common.DefaultExpireTime, cache_common.DefaultPurgeTime), Debug: debug}
		}
	}
	return cacheInstance
}

func (c *InternalCache) Get(k string) (interface{}, bool) {
	get, b := c.cache.Get(k)
	if c.Debug {
		if b {
			logrus.Tracef("[Cache] hit key: %s", k)
		} else {
			logrus.Tracef("[Cache] miss key: %s", k)
		}
	}

	return get, b
}

func (c *InternalCache) Set(k string, x interface{}, d time.Duration) {
	if c.Debug {
		logrus.Tracef("[Cache] setting key: %s, for %s", k, d.String())
	}
	c.cache.Set(k, x, d)
}

func (c *InternalCache) Delete(k string) {
	c.cache.Delete(k)
	if c.Debug {
		logrus.Tracef("[Cache] deleting key: %s", k)
	}
}

func (c *InternalCache) FlushAll() {
	c.cache.Flush()
	if c.Debug {
		logrus.Tracef("[Cache] flushing all cache")
	}
}

func (c *InternalCache) Increment(k string, n int64) error {
	return c.cache.Increment(k, n)
}

func (c *InternalCache) ItemCount() int {
	return c.cache.ItemCount()
}
