
package message_bus

const (
	MessageBusConfigPathHost     = "message_bus.host"
	MessageBusConfigPathPort     = "message_bus.port"
	MessageBusConfigPathType     = "message_bus.type"
	MessageBusConfigPathUser     = "message_bus.user"
	MessageBusConfigPathPassword = "message_bus.pass"
	MessageBusConfigPathTopic    = "message_bus.topic"
)

type MessageBusType string

const (
	MessageBusTypeActiveMQ = "active_mq"
)
