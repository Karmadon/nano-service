
package containers

import (
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

const BusChannelSize = 100

type MainBus struct {
	Events         chan *object_models.Event
	Notifications  chan *object_models.Notification
	Errors         chan error
	ControlBus     *ControlBus
	stop           chan bool
}

func NewMainBus() *MainBus {

	return &MainBus{
		Events:         make(chan *object_models.Event, BusChannelSize),
		Notifications:  make(chan *object_models.Notification, BusChannelSize),
		Errors:         make(chan error, BusChannelSize),
		ControlBus:     NewControlBus(),

		stop: make(chan bool),
	}
}

func (b *MainBus) RegisterServiceForControl(serviceName object_models.ServiceName, c chan *object_models.ControlActionMessage) {
	b.ControlBus.SetControlChannel(serviceName, c)
}

func (b *MainBus) ControlChannelForService(serviceName object_models.ServiceName) chan *object_models.ControlActionMessage {
	return b.ControlBus.ServiceControlChannel(serviceName)
}
