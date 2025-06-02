package operations

import (
	"github.com/behavioral-ai/collective/namespace"
	"github.com/behavioral-ai/collective/private"
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
	intf   private.Interface
	origin Origin
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

		return
	}
	if m.Name != messaging.ConfigEvent {
		a.agents.Broadcast(m)
	}
}

func (a *agentT) message(m *messaging.Message) {
}

func (a *agentT) advise(m *messaging.Message) {
}

func (a *agentT) subscribe(m *messaging.Message) {
}

func (a *agentT) cancel(m *messaging.Message) {
}

func (a *agentT) trace(name, task, observation, action string) {
}

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		cfg := messaging.ConfigMapContent(m)
		if cfg == nil {
			messaging.Reply(m, messaging.ConfigEmptyMapError(a.Name()), a.Name())
			return
		}
		//a.state.Update(cfg)

	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
}
