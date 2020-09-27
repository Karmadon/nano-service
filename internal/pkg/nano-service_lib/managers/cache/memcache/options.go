
package memcache

type Options struct {
	Host     string
	Port     string
	Debug    bool
	IsShared bool
}

func (o *Options) ConnectionString() string {
	return o.Host + ":" + o.Port
}
