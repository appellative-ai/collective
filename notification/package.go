package notification

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

type Sender struct {
	Message func(msg *messaging.Message) bool
	Advice  func(msg *messaging.Message) *std.Status
	Status  func(status any)
	Trace   func(name, task, observation, action string)
}

type Receiver struct {
	Message func(name string) *messaging.Message
	Advice  func(name string) *messaging.Message
}

// Send -
var Send = func() *Sender {
	return &Sender{
		Message: func(msg *messaging.Message) bool {
			//agent.message(msg)
			return true
		},
		Advice: func(msg *messaging.Message) *std.Status {
			//agent.advise(msg)
			return std.StatusOK
		},
		Status: func(status any) {
			//agent.message(msg)
		},
		Trace: func(name, task, observation, action string) {
			//agent.trace(name, task, observation, action)
		},
	}
}()

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
