package notification

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
	"time"
)

const (
	AgentName = "common:core:agent/operations/collective"
	duration  = time.Second * 30
	timeout   = time.Second * 2
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
	a.timeout = timeout
	a.ex = httpx.Do

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.configureAgents()
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
		if a.running {
			return
		}
		messaging.UpdateContent[time.Duration](m, &a.timeout)
		messaging.UpdateContent[func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)](m, &a.logFunc)
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
	if m == nil {
		return
	}
	recipients := m.To()
	if len(recipients) == 0 {
		status, _, _ := messaging.StatusContent(m)
		if status != nil {
			fmt.Printf("%v\n", status)
		} else {
			fmt.Printf("%v\n", m)
		}
		return
	}

	var local []string
	var nonLocal []string
	for _, to := range recipients {
		if std.Origin.IsLocalCollective(to) {
			local = append(local, to)
		} else {
			nonLocal = append(nonLocal, to)
		}
	}
	if len(local) > 0 {
		m.DeleteTo()
		m.AddTo(local...)
		exchange.Message(m)
	}
	// TODO : non-local
}

func (a *agentT) advise(m *messaging.Message) {
}

func (a *agentT) trace(name, task, observation, action string) {
}

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		//cfg, status := messaging.MapContent(m)
		//if !status.OK() {
		//	messaging.Reply(m, messaging.EmptyMapError(a.Name()), a.Name())
		//	return
		//}
		//a.state = initialize(cfg)
		// Initialize linked collectives
		if std.Origin.Collective != "" {
			// TODO: Initialize linked collectives by reading the configured collective links and then reference the
			//       registry for collective host names
		}
	}
	messaging.Reply(m, std.StatusOK, a.Name())
}

func (a *agentT) configureAgents() {

}
