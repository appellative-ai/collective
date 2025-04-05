package eventing

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	NamespaceName   = "resiliency:agent/behavioral-ai/collective/eventing"
	defaultDuration = time.Second * 10
)

type agentT struct {
	running  bool
	duration time.Duration

	ticker   *messaging.Ticker
	emissary *messaging.Channel
	master   *messaging.Channel
	notifier NotifyFunc
	activity ActivityFunc
}

func newAgent(notifier NotifyFunc, activity ActivityFunc) *agentT {
	a := new(agentT)
	a.duration = defaultDuration

	a.ticker = messaging.NewTicker(messaging.Emissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	a.notifier = notifier
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	if m.Event() == messaging.ConfigEvent {
		a.configure(m)
		return
	}
	if m.Event() == messaging.StartupEvent {
		a.run()
		return
	}
	if !a.running {
		return
	}
	switch m.Event() {
	case NotifyEvent:
		a.notifier(NotifyContent(m))
		return
	case ActivityEvent:
		a.activity(ActivityContent(m))
		return
	default:
	}
	switch m.Channel() {
	case messaging.Emissary:
		a.emissary.Send(m)
	case messaging.Master:
		a.master.Send(m)
	case messaging.Control:
		a.emissary.Send(m)
		a.master.Send(m)
	default:
		a.emissary.Send(m)
	}
}

func (a *agentT) configure(m *messaging.Message) {
	cfg := messaging.ConfigMapContent(m)
	if cfg == nil {
		messaging.Reply(m, messaging.ConfigEmptyStatusError(a), a.Uri())
	}
	// configure
	messaging.Reply(m, messaging.StatusOK(), a.Uri())
}

// Run - run the agent
func (a *agentT) run() {
	if a.running {
		return
	}
	go masterAttend(a)
	go emissaryAttend(a)
	a.running = true
}

func (a *agentT) addActivity(e ActivityItem) {
	if a.activity != nil {
		a.activity(e)
	} else {
		httpAddActivity("", e.Agent.Uri(), e.Event, e.Source, e.Content)
	}
}

func (a *agentT) notify(e NotifyItem) {
	if a.notifier != nil {
		a.notifier(e)
	} else {
		httpNotify(e)
	}
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}
