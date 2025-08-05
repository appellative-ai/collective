package operations

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/collective/namespace"
	"github.com/appellative-ai/collective/notification"
	"github.com/appellative-ai/collective/resolution"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
	"time"
)

const (
	AgentName = "common:core:agent/operations/collective"
	duration  = time.Second * 30
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
	exchange.RegisterConstructor(AgentName, func() messaging.Agent {
		return newAgent()
	})
}

func newAgent() *agentT {
	a := new(agentT)
	agent = a
	a.agents = messaging.NewExchange()
	a.agents.Register(resolution.NewAgent())
	a.agents.Register(namespace.NewAgent())
	a.agents.Register(notification.NewAgent())

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()
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
		if a.state.running {
			return
		}
		return
	case messaging.StartupEvent:
		if a.state.running {
			return
		}
		a.state.running = true
		a.run()
		return
	case messaging.ShutdownEvent:
		if !a.state.running {
			return
		}
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

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		cfg, status := messaging.MapContent(m)
		if !status.OK() {
			//messaging.Reply(m, messaging.EmptyMapError(a.Name()), a.Name())
			return
		}
		a.state = initialize(cfg)
		// Initialize linked collectives
		if std.Origin.Collective != "" {
			// TODO: Initialize linked collectives by reading the configured collective links and then reference the
			//       registry for collective host names
		}
	}
	messaging.Reply(m, std.StatusOK, a.Name())
}

func (a *agentT) configureLogging(log func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) {
	if log != nil {
		a.agents.Broadcast(messaging.NewConfigMessage(log))
	}
}
