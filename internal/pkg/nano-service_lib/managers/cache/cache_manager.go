
package cache

import "time"

type Manager interface {
	Get(k string) (interface{}, bool)
	GetInto(k string, obj interface{}) bool
	Set(k string, x interface{}, d time.Duration)
	Delete(k string)
	FlushAll()
	Increment(k string, n int64) error
	ItemCount() int
}
