package operations

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/event"
	"github.com/behavioral-ai/collective/namespace"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
)

const (
	AgentNamespaceName = "resiliency:agent/behavioral-ai/collective/operations"
)

// TODO : need host name
type agentT struct {
	running bool

	agents *messaging.Exchange
}

// New - create a new operations agent
func New() messaging.Agent {
	return newAgent()
}

func newAgent() *agentT {
	a := new(agentT)

	a.agents = messaging.NewExchange()
	a.agents.RegisterMailbox(content.Agent)
	a.agents.RegisterMailbox(event.Agent)
	a.agents.RegisterMailbox(namespace.Agent)
	a.agents.RegisterMailbox(timeseries.Agent)

	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return AgentNamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Event() {
	case messaging.StartupEvent:
		a.agents.Broadcast(m)
	case messaging.ShutdownEvent:
		a.agents.Broadcast(m)
	case messaging.PauseEvent:
		a.agents.Broadcast(m)
	case messaging.ResumeEvent:
		a.agents.Broadcast(m)
	}
}
