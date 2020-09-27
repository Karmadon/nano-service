
package object_models

type ServiceName string

const (
	ServiceNameNotification ServiceName = "notification"
	ServiceNameHandler      ServiceName = "handler"
	ServiceNameErrors       ServiceName = "errors"
	ServiceNameEvents       ServiceName = "events"
	ServiceNameMQ           ServiceName = "mq"
)
