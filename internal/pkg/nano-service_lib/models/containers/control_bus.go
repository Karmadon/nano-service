
package containers

import (
	"sync"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type ControlBus struct {
	bus map[object_models.ServiceName]chan *object_models.ControlActionMessage
	mu  sync.Mutex
}

func NewControlBus() *ControlBus {
	return &ControlBus{bus: make(map[object_models.ServiceName]chan *object_models.ControlActionMessage)}
}

func (b *ControlBus) ServiceControlChannel(serviceName object_models.ServiceName) chan *object_models.ControlActionMessage {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.bus[serviceName]
}

func (b *ControlBus) SetControlChannel(serviceName object_models.ServiceName, c chan *object_models.ControlActionMessage) {
	b.mu.Lock()
	b.bus[serviceName] = c
	defer b.mu.Unlock()
}
