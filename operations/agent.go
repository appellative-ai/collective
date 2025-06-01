package operations

import (
	"github.com/behavioral-ai/collective/namespace"
	"github.com/behavioral-ai/collective/repository"
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/core/messaging"
)

const (
	NamespaceName = "core:agent/operations/collective"
)

var (
	agent *agentT
)

type agentT struct {
	agents *messaging.Exchange
}

func init() {
	repository.RegisterConstructor(NamespaceName, func() messaging.Agent {
		return newAgent()
	})

}

func newAgent() *agentT {
	a := new(agentT)
	a.agents = messaging.NewExchange()
	a.agents.Register(resource.NewAgent())
	a.agents.Register(namespace.NewAgent())
	agent = a
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Name == messaging.ConfigEvent {
		//if _, ok := m2.UsingContent(m); ok {
		//	agents.Broadcast(m)
		//	messaging.Reply(m, messaging.StatusOK(), a.Name())
		//} else {
		//	messaging.Reply(m, messaging.NewStatus(http.StatusBadRequest, errors.New("invalid Using resource")), a.Name())
		//}
		return
	}
	if m.Name != messaging.ConfigEvent {
		a.agents.Broadcast(m)
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
