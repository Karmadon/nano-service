
package notification_service

import (
	"context"

	"github.com/getsentry/sentry-go"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/notification"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/containers"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type MainNotificator struct {
	Gates   notification.Gates
	options Options
	stop    chan bool
	control chan *object_models.ControlActionMessage
}

func (m *MainNotificator) Reset() error {
	panic("implement me")
}

func (m *MainNotificator) Init() error {
	return m.Gates.InitAll()
}

func NewMainNotificator(options Options) *MainNotificator {
	return &MainNotificator{options: options, stop: make(chan bool), Gates: notification.NewGates(), control: make(chan *object_models.ControlActionMessage)}
}

func (m *MainNotificator) Start(ctx context.Context, bus *containers.MainBus) {
	defer sentry.Recover()
	bus.RegisterServiceForControl(object_models.ServiceNameNotification, m.control)

	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				bus.Errors <- object_models.NewError(object_models.ServiceNameNotification, err, "context Done")
			}
		case <-m.stop:
			return
		case notificationMessage := <-bus.Notifications:
			go func() {
				defer sentry.Recover()

				gate := notification.GateNameTechTelegram

				switch notificationMessage.NotificationType {
				case object_models.NotificationTypeError:
				case object_models.NotificationTypeTechInfo:
					gate = notification.GateNameTechTelegram
				}

				err := m.Gates.Gate(gate).Send(notificationMessage)
				if err != nil {
					bus.Errors <- object_models.NewError(object_models.ServiceNameNotification, err, "error in gate:"+string(gate))
				}
			}()
		}
	}
}

func (m *MainNotificator) Stop() {
	m.Gates.StopAll()
	m.stop <- true
}
