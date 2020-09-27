
package message_bus

type Manager interface {
	Init() error

	Subscribe(topic *string) error

	SendMessage([]byte) error
	ReceiveMessages(chan []byte)
	Stop()
}
