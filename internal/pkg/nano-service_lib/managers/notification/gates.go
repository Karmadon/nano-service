

package notification

import (
	log "github.com/sirupsen/logrus"
)

type Gates map[GateName]Manager

func NewGates() map[GateName]Manager {
	return make(Gates)
}

func (g *Gates) AddGate(name GateName, manager Manager) {
	(*g)[name] = manager
}

func (g *Gates) Gate(name GateName) Manager {
	return (*g)[name]
}

func (g *Gates) StopAll() {
	for gateName, gate := range *g {
		log.Trace("[GATE:" + gateName + "] stopping...")

		gate.Stop()
	}
}

func (g *Gates) InitAll() error {
	for gateName, gate := range *g {
		log.Trace("[GATE:" + gateName + "] inits")

		err := gate.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Gates) StartAll() {
	for gateName, gate := range *g {
		log.Trace("[GATE:" + gateName + "] starting...")

		go gate.Start()
	}
}
