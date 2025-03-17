package event

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	AgentNamespaceName = "resiliency:agent/behavioral-ai/collective/event"
	agentUri           = AgentNamespaceName
	defaultDuration    = time.Second * 10
)

type agentT struct {
	running  bool
	agentId  string
	duration time.Duration

	ticker     *messaging.Ticker
	emissary   *messaging.Channel
	master     *messaging.Channel
	notifier   messaging.NotifyFunc
	dispatcher messaging.Dispatcher
	activity   messaging.ActivityFunc
}

func newAgent(dispatcher messaging.Dispatcher) *agentT {
	a := new(agentT)
	a.agentId = agentUri
	a.duration = defaultDuration
	a.ticker = messaging.NewTicker(messaging.Emissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	a.dispatcher = dispatcher
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return a.agentId }

// Name - agent name
func (a *agentT) Name() string { return AgentNamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.Emissary:
		a.emissary.Send(m)
	case messaging.Master:
		a.master.Send(m)
	case messaging.Control:
		if m.ContentType() == messaging.ContentTypeNotify {
			a.notify(messaging.NotifyContent(m))
			return
		}
		if m.ContentType() == messaging.ContentTypeActivity {
			a.addActivity(messaging.ActivityContent(m))
			return
		}
		a.emissary.Send(m)
		a.master.Send(m)
	default:
		a.emissary.Send(m)
	}
}

// Run - run the agent
func (a *agentT) Run() {
	if a.running {
		return
	}
	go masterAttend(a)
	go emissaryAttend(a)
	a.running = true
}

// Shutdown - shutdown the agent
func (a *agentT) Shutdown() {
	if !a.emissary.IsClosed() {
		a.emissary.Send(messaging.Shutdown)
	}
	if !a.master.IsClosed() {
		a.master.Send(messaging.Shutdown)
	}
}

func (a *agentT) addActivity(e *messaging.ActivityItem) {
	if a.activity != nil {
		a.activity(*e)
	} else {
		httpAddActivity("", e.Agent.Uri(), e.Event, e.Source, e.Content)
	}
}

func (a *agentT) notify(e messaging.NotifyItem) {
	if e == nil {
		return
	}
	if a.notifier != nil {
		a.notifier(e)
	} else {
		httpNotify(e)
	}
}

func (a *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(a, a.dispatcher, channel, event)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}
