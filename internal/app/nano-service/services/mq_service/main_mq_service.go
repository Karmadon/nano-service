
package mq_service

import (
	"context"

	"github.com/getsentry/sentry-go"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/message_bus"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/containers"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type MainMQ struct {
	options *Options
	bus     message_bus.Manager
	stop    chan bool
	control chan *object_models.ControlActionMessage
}

func NewMainMQ(options *Options, bus message_bus.Manager) *MainMQ {
	return &MainMQ{options: options, bus: bus, stop: make(chan bool), control: make(chan *object_models.ControlActionMessage)}
}

func (m *MainMQ) Init() error {
	err := m.bus.Init()
	if err != nil {
		return err
	}

	err = m.bus.Subscribe(nil)
	if err != nil {
		return err
	}

	return nil
}

func (m *MainMQ) Start(ctx context.Context, bus *containers.MainBus) {
	defer sentry.Recover()
	bus.RegisterServiceForControl(object_models.ServiceNameMQ, m.control)

}

func (m *MainMQ) Stop() {
	m.bus.Stop()

	m.stop <- true
}

func (m *MainMQ) Reset() error {
	m.bus.Stop()

	return m.Init()
}
