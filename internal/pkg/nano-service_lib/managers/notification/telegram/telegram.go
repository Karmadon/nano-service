

package telegram

import (
	"github.com/getsentry/sentry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type Bot struct {
	options *Options
	client  *tgbotapi.BotAPI
}

func (b *Bot) Start() {
	defer sentry.Recover()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.client.GetUpdatesChan(u)
	if err != nil {
		log.Error(err)
		return
	}

	for update := range updates {
		if b.options.Debug {
			log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}
}

func (b *Bot) Stop() {
	b.client.StopReceivingUpdates()
}

func NewBot(options *Options) *Bot {
	return &Bot{options: options, client: nil}
}

func (b *Bot) Init() error {
	bot, err := tgbotapi.NewBotAPI(b.options.BotToken)
	if err != nil {
		return err
	}

	if b.options.Debug {
		log.Infof("Authorized on account %s", bot.Self.UserName)
		//bot.Debug = b.options.Debug
	}

	b.client = bot

	return nil
}

func (b *Bot) Send(notification *object_models.Notification) error {
	return b.sendMessage(notification.NotificationType, notification.Header, notification.Message, notification.Id)
}

func (b *Bot) sendMessage(notificationType object_models.NotificationType, header string, message string, id string) error {
	wholeMessage := "**" + header + "**" + "\n" + message + "\n"

	msg := tgbotapi.NewMessage(b.options.ChatID, wholeMessage)
	msg.ParseMode = "markdown"

	_, err := b.client.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
