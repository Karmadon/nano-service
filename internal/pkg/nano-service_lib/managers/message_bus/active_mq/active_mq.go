
package active_mq

import (
	"crypto/tls"

	"github.com/go-stomp/stomp"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Bus struct {
	options       *Options
	client        *stomp.Conn
	netConnection *tls.Conn
	subscription  *stomp.Subscription
	stop          chan bool
}

func NewBus(options *Options) *Bus {
	return &Bus{options: options, stop: make(chan bool)}
}

func (b *Bus) Init() error {
	var err error
	var options = []func(*stomp.Conn) error{
		stomp.ConnOpt.Login(b.options.User, b.options.Password),
		stomp.ConnOpt.Host("/"),
	}

	b.netConnection, err = tls.Dial("tcp", b.options.ConnectionString(), &tls.Config{})
	if err != nil {
		return errors.Wrap(err, "cannot establish TLS connection with MQ server")
	}

	b.client, err = stomp.Connect(b.netConnection, options...)
	if err != nil {
		return errors.Wrap(err, "cannot connect MQ server")
	}

	log.Trace("[MANAGER MQ] connected")
	return nil
}

func (b *Bus) Subscribe(topic *string) error {
	var err error

	if topic != nil {
		b.options.Topic = *topic
	}

	b.subscription, err = b.client.Subscribe(b.options.Topic, stomp.AckAuto, stomp.SubscribeOpt.Id("nano-service:"+uuid.New().String()))
	if err != nil {
		return errors.Wrap(err, "cannot subscribe to"+b.options.Topic)
	}

	log.Trace("[MANAGER MQ] subscribed to " + b.options.Topic)
	return nil
}

func (b *Bus) SendMessage(message []byte) error {
	log.Trace("[MANAGER MQ] sending message")

	return b.client.Send(b.options.Topic, "application/json", message, stomp.SendOpt.Receipt)
}

func (b *Bus) Stop() {
	b.stop <- true
}

func (b *Bus) ReceiveMessages(channel chan []byte) {
	for {
		if !b.subscription.Active() {
			log.Warn("[MANAGER MQ] Restarting...")
			err := b.Init()
			if err != nil {
				log.Error(err)
			}
			err = b.Subscribe(nil)
			if err != nil {
				log.Error(err)
			}

			continue
		}

		select {
		case <-b.stop:
			log.Trace("[MANAGER MQ] stopping...")
			return
		case msg := <-b.subscription.C:
			if msg != nil {
				log.Trace("[MANAGER MQ] new message")
				channel <- msg.Body
			}
		}
	}
}
