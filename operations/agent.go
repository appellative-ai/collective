package operations

import (
	"errors"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/messaging"
	"sync/atomic"
)

const (
	AgentName = "common:core:agent/operations/collective"
)

var (
	agent *agentT
)

type agentT struct {
	running atomic.Bool
	origin  *OriginT
}

func init() {
	exchange.RegisterConstructor(AgentName, func() messaging.Agent {
		return newAgent()
	})
}

func newAgent() *agentT {
	a := new(agentT)
	agent = a
	a.running.Store(false)
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return AgentName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		a.configure(m)
		return
	case messaging.StartupEvent:
		if a.running.Load() {
			return
		}
		a.running.Store(true)
		a.run()
		return
	case messaging.ShutdownEvent:
		if !a.running.Load() {
			return
		}
		a.running.Store(false)
	}
}

// Run - run the agent
func (a *agentT) run() {}

func (a *agentT) startup() error {
	if a.origin == nil {
		return errors.New("origin is required")
	}
	return nil
}

func (a *agentT) configure(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}

}
