
package object_models

type EventType string

const (
	EventTypeError EventType = "error"
	EventTypeInfo  EventType = "info"
	EventTypeTech  EventType = "tech"
)

type Event struct {
	EventType EventType
	Message   string
}

func NewEvent(eventType EventType, message string) *Event {
	return &Event{EventType: eventType, Message: message}
}
