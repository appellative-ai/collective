package event

import (
	"github.com/behavioral-ai/core/messaging"
	"time"
)

const (
	AgentNamespaceName = "resiliency:agent/behavioral-ai/collective/event"
	defaultDuration    = time.Second * 10
)

type agentT struct {
	running  bool
	duration time.Duration

	ticker     *messaging.Ticker
	emissary   *messaging.Channel
	master     *messaging.Channel
	notifier   NotifyFunc
	dispatcher Dispatcher
	activity   ActivityFunc
}

func newAgent(notifier NotifyFunc, dispatcher Dispatcher, activity ActivityFunc) *agentT {
	a := new(agentT)
	a.duration = defaultDuration

	a.ticker = messaging.NewTicker(messaging.Emissary, a.duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	a.dispatcher = dispatcher
	a.notifier = notifier
	a.dispatcher = dispatcher
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return AgentNamespaceName }

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
	switch m.Channel() {
	case messaging.Emissary:
		a.emissary.Send(m)
	case messaging.Master:
		a.master.Send(m)
	case messaging.Control:
		if m.ContentType() == ContentTypeNotify {
			a.notify(NotifyContent(m))
			return
		}
		if m.ContentType() == ContentTypeActivity {
			a.addActivity(ActivityContent(m))
			return
		}
		if m.ContentType() == ContentTypeDispatch {
			a.dispatch(DispatchContent(m))
			return
		}
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

func (a *agentT) dispatch(e DispatchItem) {
	if a.dispatcher != nil {
		a.dispatcher.Dispatch(a, e.Channel, e.Event)
	}
}

func (a *agentT) dispatchArgs(channel any, event string) {
	if a.dispatcher != nil {
		a.dispatcher.Dispatch(a, channel, event)
	}
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}
