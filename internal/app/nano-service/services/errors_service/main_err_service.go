package errors_service

import (
	"context"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/containers"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type MainErrService struct {
	options Options
	control chan *object_models.ControlActionMessage
	stop    chan bool
}

func (m *MainErrService) Reset() error {
	panic("implement me")
}

func NewMainErrService(options Options) *MainErrService {
	return &MainErrService{options: options, stop: make(chan bool), control: make(chan *object_models.ControlActionMessage)}
}

func (m *MainErrService) Init() error {
	return nil
}

func (m *MainErrService) Start(ctx context.Context, bus *containers.MainBus) {
	defer sentry.Recover()
	bus.RegisterServiceForControl(object_models.ServiceNameErrors, m.control)

	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Error(err)
			}
		case <-m.stop:
			return
		case err := <-bus.Errors:
			log.Errorf("[SERVICE ERRORS] %s", err.Error())

			bus.Events <- object_models.NewEvent(object_models.EventTypeError, err.Error())
		}
	}
}

func (m *MainErrService) Stop() {
	m.stop <- true
}
