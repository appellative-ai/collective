package notification

import "github.com/appellative-ai/core/messaging"

// Interface - notification interface
type Interface struct {
	Message func(msg *messaging.Message) bool
	Advise  func(msg *messaging.Message) *messaging.Status
	Trace   func(name, task, observation, action string)
}

// Notifier -
var Notifier = func() *Interface {
	return &Interface{
		Message: func(msg *messaging.Message) bool {
			//agent.message(msg)
			return true
		},
		Advise: func(msg *messaging.Message) *messaging.Status {
			//agent.advise(msg)
			return messaging.StatusOK()
		},
		Trace: func(name, task, observation, action string) {
			//agent.trace(name, task, observation, action)
		},
	}
}()
