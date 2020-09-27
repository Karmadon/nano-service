

package services

import (
	"context"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/containers"
)

type Servicer interface {
	Init() error
	Start(ctx context.Context, bus *containers.MainBus)
	Stop()
	Reset() error
}
