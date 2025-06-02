package operations

import (
	"fmt"
	"github.com/behavioral-ai/collective/namespace"
	"github.com/behavioral-ai/collective/private"
	"github.com/behavioral-ai/collective/repository"
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	NamespaceName = "core:agent/operations/collective"
	duration      = time.Second * 30
)

var (
	agent *agentT
)

type agentT struct {
	state  *operationsT
	agents *messaging.Exchange

	ticker   *messaging.Ticker
	emissary *messaging.Channel
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

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.configureAgents()
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
	if !a.state.running {
		if m.Name == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Name == messaging.StartupEvent {
			a.run()
			a.state.running = true
			return
		}
		return
	}
	if m.Name == messaging.ShutdownEvent {
		a.state.running = false
	}
	switch m.Channel() {
	case messaging.ChannelControl, messaging.ChannelEmissary:
		a.emissary.C <- m
	default:
		fmt.Printf("limiter - invalid channel %v\n", m)
	}
}

// Run - run the agent
func (a *agentT) run() {
	go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
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
		var ok bool
		if a.state, ok = initialize(m); !ok {
			return
		}
		// TODO: Need to set linked collective attributes
	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
}

func (a *agentT) configureAgents() {
	a.agents.Broadcast(private.NewInterfaceMessage(private.Interface{
		Rep: representation,
		Ctx: context,
		Th:  thing,
		Rel: relation,
	}))
}
