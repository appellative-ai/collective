package operations

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/collective/namespace"
	"github.com/appellative-ai/collective/notification"
	"github.com/appellative-ai/collective/resolution"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"sync/atomic"
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
	running  atomic.Bool
	state    *operationsT
	exchange rest.Exchange
	agents   *messaging.Exchange

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
	a.running.Store(false)
	a.exchange = httpx.Do
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
