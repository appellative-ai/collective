package notification

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
	"time"
)

type Interface struct {
	Message func(msg *messaging.Message) *std.Status
	Trace   func(name, task, observation, action string)

	Status   func(name string, status any)
	Exchange func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)
}

// Notifier -
var Notifier = func() *Interface {
	return &Interface{
		Message: func(msg *messaging.Message) *std.Status {
			return agent.message(msg)
		},
		Trace: func(name, task, observation, action string) {
			agent.trace(name, task, observation, action)
		},
		Status: func(name string, status any) {
			agent.status(name, status)
		},
		Exchange: func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
			agent.exchangeLog(start, duration, route, req, resp, timeout)
		},
	}
}()

/*
type Receiver struct {
	Message func(name string) *messaging.Message
	Advice  func(name string) *messaging.Message
}

// Receive -
var Receive = func() *Receiver {
	return &Receiver{
		Message: func(name string) *messaging.Message {
			//agent.message(msg)
			return nil
		},
		Advice: func(name string) *messaging.Message {
			//agent.message(msg)
			return nil
		},
	}
}()


*/
