
package object_models

type ControlAction string

const (
	ControlActionResetService ControlAction = "service_reset"
	ControlActionCleanCache   ControlAction = "cache_flush"
)

type ControlActionMessages []ControlActionMessage

type ControlActionMessage struct {
	Action  ControlAction
	Payload string
}

func NewControlActionMessage(action ControlAction, payload string) *ControlActionMessage {
	return &ControlActionMessage{Action: action, Payload: payload}
}
