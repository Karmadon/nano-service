
package notification

import "github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"

type Manager interface {
	Init() error
	Start()
	Stop()
	Send(notification *object_models.Notification) error
}
