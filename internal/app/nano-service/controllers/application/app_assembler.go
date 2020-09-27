package application

import (
	"errors"

	config "github.com/spf13/viper"

	"github.com/karmadon/nano-service/internal/app/nano-service/services/errors_service"
	"github.com/karmadon/nano-service/internal/app/nano-service/services/event_service"
	"github.com/karmadon/nano-service/internal/app/nano-service/services/handler_service"
	"github.com/karmadon/nano-service/internal/app/nano-service/services/mq_service"
	"github.com/karmadon/nano-service/internal/app/nano-service/services/notification_service"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/cache"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/message_bus"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/message_bus/active_mq"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/notification"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/notification/telegram"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

func (a *Application) Assemble() error {
	var err error

	mainMQ, err := mqService()
	if err != nil {
		return err
	}
	a.registry.AddService(object_models.ServiceNameMQ, mainMQ)

	mainNotificator, err := notificator()
	if err != nil {
		return err
	}
	a.registry.AddService(object_models.ServiceNameNotification, mainNotificator)

	mainErrService, err := errorService()
	if err != nil {
		return err
	}
	a.registry.AddService(object_models.ServiceNameErrors, mainErrService)

	mainEventsService, err := eventService()
	if err != nil {
		return err
	}
	a.registry.AddService(object_models.ServiceNameEvents, mainEventsService)

	return nil
}

func handler(internalCache cache.Manager) (*handler_service.MainHandler, error) {
	var srv = handler_service.NewMainHandler(handler_service.Options{})

	return srv, nil
}

func errorService() (*errors_service.MainErrService, error) {
	var srv = errors_service.NewMainErrService(errors_service.Options{})

	return srv, nil
}

func eventService() (*event_service.MainEventsService, error) {
	var srv = event_service.NewMainEventsService(event_service.Options{})

	return srv, nil
}

func notificator() (*notification_service.MainNotificator, error) {
	var srv = notification_service.NewMainNotificator(notification_service.Options{})

	mainTlgOptions := &telegram.Options{
		BotToken: config.GetString(notification.MainTelegramConfigPathToken),
		ChatID:   config.GetInt64(notification.MainTelegramConfigPathChatID),
		Debug:    config.GetBool(ConfPathDebugNotificator),
	}
	mainTlgBot := telegram.NewBot(mainTlgOptions)
	srv.Gates.AddGate(notification.GateNameMainTelegram, mainTlgBot)

	techTlgOptions := &telegram.Options{
		BotToken: config.GetString(notification.TechTelegramConfigPathToken),
		ChatID:   config.GetInt64(notification.TechTelegramConfigPathChatID),
		Debug:    config.GetBool(ConfPathDebugNotificator),
	}

	techTlgBot := telegram.NewBot(techTlgOptions)
	srv.Gates.AddGate(notification.GateNameTechTelegram, techTlgBot)

	return srv, nil
}

func mqService() (*mq_service.MainMQ, error) {
	var bus message_bus.Manager

	switch message_bus.MessageBusType(config.GetString(message_bus.MessageBusConfigPathType)) {
	case message_bus.MessageBusTypeActiveMQ:
		activeMQOptions := &active_mq.Options{
			Host:     config.GetString(message_bus.MessageBusConfigPathHost),
			Port:     config.GetString(message_bus.MessageBusConfigPathPort),
			User:     config.GetString(message_bus.MessageBusConfigPathUser),
			Password: config.GetString(message_bus.MessageBusConfigPathPassword),
			Topic:    config.GetString(message_bus.MessageBusConfigPathTopic),
		}
		bus = active_mq.NewBus(activeMQOptions)

	default:
		return nil, errors.New("no message bus found")
	}

	mqOptions := &mq_service.Options{
		Debug: config.GetBool(ConfPathDebugMQ),
	}
	service := mq_service.NewMainMQ(mqOptions, bus)
	err := service.Init()
	if err != nil {
		return nil, err
	}

	return service, nil
}
