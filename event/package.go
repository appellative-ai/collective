package event

import (
	http2 "github.com/behavioral-ai/core/http"
	"github.com/behavioral-ai/core/messaging"
)

// Agent - content resolution in the real world
var (
	Agent    messaging.Agent
	Exchange http2.Exchange
)

func init() {
	Exchange = http2.Do
	Agent = newAgent(nil)
	Agent.Run()
}

func Override(activity messaging.ActivityFunc, notifier messaging.NotifyFunc, dispatcher messaging.Dispatcher) {
	if agent, ok := any(Agent).(*agentT); ok {
		if activity != nil {
			agent.activity = activity
		}
		if notifier != nil {
			agent.notifier = notifier
		}
		if dispatcher != nil {
			agent.dispatcher = dispatcher
		}
	}
}
