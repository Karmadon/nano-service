
package active_mq

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	Topic    string
}

func (o Options) ConnectionString() string {
	return o.Host + ":" + o.Port
}
