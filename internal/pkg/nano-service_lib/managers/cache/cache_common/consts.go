
package cache_common

import "time"

const (
	NoExpiration             time.Duration = -1
	DefaultExpiration        time.Duration = 0
	DefaultExpireTime                      = 1 * time.Minute
	DefaultPurgeTime                       = 1 * time.Minute
)

type TagPrefix string

const (
	TagPrefixNoPrefix      TagPrefix = "NP:"
)

type CacheType string

const (
	CacheTypeMemCache = "memcache"
)

const (
	CachePathConfigType = "cache.type"
)
