package operations

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "resiliency:agent/behavioral-ai/collective/operations"
)

type agentT struct{}

func newAgent() *agentT {
	a := new(agentT)
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Event() == messaging.ConfigEvent {
		a.configure(m)
	}
}

func (a *agentT) configure(m *messaging.Message) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	}
	// configure
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}
