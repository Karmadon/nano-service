
package handler_service

import (
	"context"

	"github.com/getsentry/sentry-go"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/cache"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/containers"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type MainHandler struct {
	options      Options
	cacheManager cache.Manager
	stop         chan bool
	control      chan *object_models.ControlActionMessage
}

func (m *MainHandler) Reset() error {
	panic("implement me")
}

func NewMainHandler(options Options) *MainHandler {
	return &MainHandler{options: options, stop: make(chan bool), control: make(chan *object_models.ControlActionMessage)}
}

func (m *MainHandler) Init() error {
	return nil
}

func (m *MainHandler) Start(ctx context.Context, bus *containers.MainBus) {
	defer sentry.Recover()
	bus.RegisterServiceForControl(object_models.ServiceNameHandler, m.control)

	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				bus.Errors <- object_models.NewError(object_models.ServiceNameHandler, err, "context done")
			}
		case <-m.stop:
			return

		}
	}
}

func (m *MainHandler) Stop() {
	m.stop <- true
}
