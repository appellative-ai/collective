package operations

import (
	"errors"
	m2 "github.com/behavioral-ai/collective/messaging"
	"github.com/behavioral-ai/collective/namespace"
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

const (
	NamespaceName = "collective:agent/operations"
)

var (
	agents = messaging.NewExchange()
	agent  *agentT
)

type agentT struct{}

func init() {
	host.Register(newAgent())
	agents.Register(resource.NewAgent())
	agents.Register(namespace.NewAgent())
}

func newAgent() *agentT {
	agent = new(agentT)
	return agent
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
		if _, ok := m2.UsingContent(m); ok {
			agents.Broadcast(m)
			messaging.Reply(m, messaging.StatusOK(), a.Uri())
		} else {
			messaging.Reply(m, messaging.NewStatus(http.StatusBadRequest, errors.New("invalid Using resource")), a.Uri())
		}
	}
	if m.Event() == messaging.ShutdownEvent {
		agents.Broadcast(m)
	}
}

/*
func (a *agentT) configure(m *messaging.Message) {
	//ur := messaging.messaging.ConfigMapContent(m)
	//if cfg == nil {
	//	messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	//}
	// configure
	//messaging.Reply(m, messaging.StatusOK(), a.Uri())
	agents.Broadcast(m)
}


*/
