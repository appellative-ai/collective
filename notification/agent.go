package notification

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
	"strings"
	"sync/atomic"
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
	running    atomic.Bool
	timeout    time.Duration
	collective string

	exchange    rest.Exchange
	logExchange func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)
	logStatus   func(status any)

	ticker   *messaging.Ticker
	emissary *messaging.Channel
}

func NewAgent() messaging.Agent {
	return newAgent()
}

func newAgent() *agentT {
	a := new(agentT)
	agent = a
	a.running.Store(false)
	a.timeout = timeout
	a.exchange = httpx.Do

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

func (a *agentT) message(m *messaging.Message) *std.Status {
	if m == nil {
		return std.StatusOK
	}
	recipients := m.To()
	if len(recipients) == 0 {
		status, _, _ := messaging.StatusContent(m)
		if status != nil {
			fmt.Printf("%v\n", status)
		} else {
			fmt.Printf("%v\n", m)
		}
		return std.StatusOK
	}

	var local []string
	var nonLocal []string
	for _, to := range recipients {
		if a.isLocalCollective(to) {
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
	return std.StatusOK
}

func (a *agentT) trace(name, task, observation, action string) {
}

func (a *agentT) status(name string, status any) {
}

func (a *agentT) exchangeLog(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
}

func (a *agentT) isLocalCollective(name string) bool {
	if strings.HasPrefix(name, a.collective+":") {
		return true
	}
	return false
}
