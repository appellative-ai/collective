package eventtest

import (
	"github.com/behavioral-ai/collective/event"
	"github.com/behavioral-ai/core/messaging"
)

type agentT struct {
	notifier   event.NotifyFunc
	activity   event.ActivityFunc
	dispatcher event.Dispatcher
}

func New(dispatcher event.Dispatcher) messaging.Agent {
	return newAgent(dispatcher)
}

func newAgent(dispatcher event.Dispatcher) *agentT {
	a := new(agentT)
	a.notifier = event.Notify
	a.activity = event.Activity
	if dispatcher == nil {
		a.dispatcher = event.NewTraceDispatcher()
	} else {
		a.dispatcher = dispatcher
	}
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return event.AgentNamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.Control:
		if m.ContentType() == event.ContentTypeNotify {
			a.notifier(event.NotifyContent(m))
			return
		}
		if m.ContentType() == event.ContentTypeActivity {
			a.activity(event.ActivityContent(m))
			return
		}
		if m.ContentType() == event.ContentTypeDispatch {
			e := event.DispatchContent(m)
			a.dispatcher.Dispatch(e.Agent, e.Channel, e.Event)
			return
		}
	default:
	}
}

// Run - run the agent
func (a *agentT) Run() {}
