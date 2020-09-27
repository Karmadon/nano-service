
package registry

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/karmadon/nano-service/internal/app/nano-service/services"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/containers"
	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/models/object_models"
)

type Registry struct {
	Container map[object_models.ServiceName]services.Servicer
	stop      chan bool

	mu sync.Mutex
}

func NewRegistry() *Registry {
	return &Registry{Container: make(map[object_models.ServiceName]services.Servicer), stop: make(chan bool)}
}

func (r *Registry) Service(name object_models.ServiceName) services.Servicer {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.Container[name]
}

func (r *Registry) AddService(name object_models.ServiceName, srv services.Servicer) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Container[name] = srv
}

func (r *Registry) StartAll(ctx context.Context) {
	container := containers.NewMainBus()

	for serviceName, service := range r.Container {
		go service.Start(ctx, container)

		log.Trace("[SRV:" + serviceName + "] started")
	}

	for {
		select {
		case <-ctx.Done():
		case <-r.stop:
			return
		}
	}
}

func (r *Registry) StopAll() {
	for serviceName, service := range r.Container {
		log.Trace("[SRV:" + serviceName + "] stopping...")
		service.Stop()
	}

	r.stop <- true
}

func (r *Registry) InitAll() error {
	for serviceName, service := range r.Container {
		log.Trace("[SRV:" + serviceName + "] start init")
		err := service.Init()
		if err != nil {
			return err
		}
	}

	return nil
}
