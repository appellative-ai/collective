package eventtest

import (
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/messaging"
)

type agentT struct {
	notifier eventing.NotifyFunc
	activity eventing.ActivityFunc
}

func New() messaging.Agent {
	return newAgent()
}

func newAgent() *agentT {
	a := new(agentT)
	a.notifier = eventing.OutputNotify
	a.activity = eventing.Activity
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return eventing.NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Event() {
	case eventing.NotifyEvent:
		a.notifier(eventing.NotifyContent(m))
	case eventing.ActivityEvent:
		a.activity(eventing.ActivityContent(m))
	default:
	}
}

// Run - run the agent
func (a *agentT) Run() {}
