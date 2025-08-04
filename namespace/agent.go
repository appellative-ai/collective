package namespace

import (
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"strings"
	"time"
)

const (
	NamespaceAgentName = "common:core:agent/namespace/collective"
	duration           = time.Second * 10
	timeout            = time.Second * 3
)

var (
	agent *agentT
)

type agentT struct {
	running bool
	timeout time.Duration
	hosts   []string
	logFunc func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)

	ex       rest.Exchange
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
	a.hosts = []string{"invalid-host"}
	a.ex = httpx.Do
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
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
		messaging.UpdateContent[time.Duration](m, &a.timeout)
		messaging.UpdateContent[[]string](m, &a.hosts)
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

func (a *agentT) log(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	if a.logFunc == nil {
		return
	}
	a.logFunc(start, duration, route, req, resp, timeout)
}

func (a *agentT) url(path string) string {
	scheme := "https"
	i := strings.Index(a.hosts[0], localHost)
	if i >= 0 {
		scheme = "http"
	}
	return scheme + "://" + a.hosts[0] + path
}
