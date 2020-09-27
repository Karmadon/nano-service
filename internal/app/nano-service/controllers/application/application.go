

package application

import (
	"context"

	"github.com/karmadon/nano-service/internal/app/nano-service/controllers/registry"
)

type Application struct {
	options  *Options
	registry *registry.Registry
}

func NewApplication(options *Options) (*Application, error) {
	app := &Application{options: options, registry: registry.NewRegistry()}

	return app, nil
}

func (a *Application) Prepare() error {
	return a.registry.InitAll()
}

func (a *Application) Start(c context.Context) {
	a.registry.StartAll(c)
}

func (a *Application) Stop() {
	a.registry.StopAll()
}
