package namespace

import (
	"github.com/appellative-ai/collective/private"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
	"time"
)

const (
	NamespaceAgentName = "common:core:agent/namespace/collective"
	defaultDuration    = time.Second * 10
)

var (
	agent *agentT
)

type agentT struct {
	running  bool
	duration time.Duration

	ex       rest.Exchange
	ticker   *messaging.Ticker
	emissary *messaging.Channel
}

func NewAgent() messaging.Agent {
	agent = newAgent()
	return agent
}

func newAgent() *agentT {
	a := new(agentT)
	a.duration = defaultDuration
	//a.ex =
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, a.duration)
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
	if !a.running {
		if m.Name == messaging.ConfigEvent {
			a.configure(m)
			return
		}
		if m.Name == messaging.StartupEvent {
			a.run()
			a.running = true
			return
		}
		return
	}
	if m.Name == messaging.ShutdownEvent {
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

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case private.ContentTypeInterface:
		_, status := private.InterfaceContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
		}
		//a.intf = intf
	}
	messaging.Reply(m, std.StatusOK, a.Name())
}

// Run - run the agent
func (a *agentT) run() {
	go emissaryAttend(a)

}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}
