package namespace

import (
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"time"
)

const (
	NamespaceAgentName = "common:core:agent/namespace/collective"
	defaultDuration    = time.Second * 10
	defaultTimeout     = time.Second * 3
)

var (
	agent *agentT
)

type agentT struct {
	running bool
	timeout time.Duration

	ex       rest.Exchange
	logFunc  func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)
	ticker   *messaging.Ticker
	emissary *messaging.Channel
}

func NewAgent() messaging.Agent {
	return newAgent()
}

func newAgent() *agentT {
	a := new(agentT)
	agent = a
	a.timeout = defaultTimeout
	a.ex = httpx.Do
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, defaultDuration)
	a.emissary = messaging.NewEmissaryChannel()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent name
func (a *agentT) Name() string { return NamespaceAgentName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		if a.running {
			return
		}
		messaging.UpdateContent[time.Duration](&a.timeout, m)
		messaging.UpdateContent[func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)](&a.logFunc, m)
		return
	case messaging.StartupEvent:
		if a.running {
			return
		}
		a.running = true
		a.run()
		return
	case messaging.ShutdownEvent:
		if !a.running {
			return
		}
		a.running = false
	}
	switch m.Channel() {
	case messaging.ChannelEmissary:
		a.emissary.Send(m)
	case messaging.ChannelControl:
		a.emissary.Send(m)
	default:
		a.emissary.Send(m)
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
