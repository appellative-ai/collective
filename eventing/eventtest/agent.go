package eventtest

import (
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/messaging"
)

type agentT struct {
	notifier   eventing.NotifyFunc
	activity   eventing.ActivityFunc
	dispatcher eventing.Dispatcher
}

func New(dispatcher eventing.Dispatcher) messaging.Agent {
	return newAgent(dispatcher)
}

func newAgent(dispatcher eventing.Dispatcher) *agentT {
	a := new(agentT)
	a.notifier = eventing.OutputNotify
	a.activity = eventing.Activity
	if dispatcher == nil {
		a.dispatcher = eventing.NewTraceDispatcher()
	} else {
		a.dispatcher = dispatcher
	}
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return eventing.AgentNamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.Control:
		if m.ContentType() == eventing.ContentTypeNotify {
			a.notifier(eventing.NotifyContent(m))
			return
		}
		if m.ContentType() == eventing.ContentTypeActivity {
			a.activity(eventing.ActivityContent(m))
			return
		}
		if m.ContentType() == eventing.ContentTypeDispatch {
			e := eventing.DispatchContent(m)
			a.dispatcher.Dispatch(e.Agent, e.Channel, e.Event)
			return
		}
	default:
	}
}

// Run - run the agent
func (a *agentT) Run() {}
