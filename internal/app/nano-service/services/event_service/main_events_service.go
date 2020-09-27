package event_service

import (
	"context"

	"github.com/getsentry/sentry-go"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/containers"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type MainEventsService struct {
	options Options
	control chan *object_models.ControlActionMessage
	stop    chan bool
}

func (m *MainEventsService) Reset() error {
	panic("implement me")
}

func NewMainEventsService(options Options) *MainEventsService {
	return &MainEventsService{options: options, stop: make(chan bool), control: make(chan *object_models.ControlActionMessage)}
}

func (m *MainEventsService) Init() error {
	return nil
}

func (m *MainEventsService) Start(ctx context.Context, bus *containers.MainBus) {
	defer sentry.Recover()
	bus.RegisterServiceForControl(object_models.ServiceNameEvents, m.control)

	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				bus.Errors <- object_models.NewError(object_models.ServiceNameEvents, err, "context done")
			}
		case <-m.stop:
			return
		case event := <-bus.Events:
			switch event.EventType {
			case object_models.EventTypeTech:
				bus.Notifications <- object_models.NewNotification(object_models.NotificationTypeTechInfo, "INFO", event.Message, "")
			case object_models.EventTypeError:
				bus.Notifications <- object_models.NewNotification(object_models.NotificationTypeError, "ERROR", event.Message, "")
			}
		}
	}
}

func (m *MainEventsService) Stop() {
	m.stop <- true
}
