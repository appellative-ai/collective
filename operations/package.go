package operations

import "github.com/appellative-ai/core/messaging"

const (
// ServiceKind = "service"
)

// Origin map and host keys
const (
	RegistryHost1Key = "registry-host1"
	RegistryHost2Key = "registry-host2"
)

// Notification - notification interface
type Notification struct {
	Message func(msg *messaging.Message) bool
	Advise  func(msg *messaging.Message) *messaging.Status
	Trace   func(name, task, observation, action string)
}

// Notifier -
var Notifier = func() *Notification {
	return &Notification{
		Message: func(msg *messaging.Message) bool {
			agent.message(msg)
			return true
		},
		Advise: func(msg *messaging.Message) *messaging.Status {
			agent.advise(msg)
			return messaging.StatusOK()
		},
		Trace: func(name, task, observation, action string) {
			agent.trace(name, task, observation, action)
		},
	}
}()

// Service - servicing add functions
type Service struct {
	AddRepresentation func(name, author, contentType string, value any) *messaging.Status
	AddContext        func(name, author, contentType string, value any) *messaging.Status

	AddThing func(name, cname, author string) *messaging.Status
	AddJoin  func(name, cname, thing1, thing2, author string) *messaging.Status
}

// Serve -
var Serve = func() *Service {
	return &Service{
		AddRepresentation: func(name, author, contentType string, value any) *messaging.Status {
			return messaging.StatusOK()
		},
		AddContext: func(name, author, contentType string, t any) *messaging.Status {
			return messaging.StatusOK()
		},
		AddThing: func(name, cname, author string) *messaging.Status {
			return messaging.StatusOK()
		},
		AddJoin: func(name, cname, thing1, thing2, author string) *messaging.Status {
			return messaging.StatusOK()
		},
		/*
			AddOrderJoin: func(name, cname, thing1, thing2, author string) *messaging.Status {
				return messaging.StatusOK()
			},

		*/
		/*
			SubscriptionCreate: func(msg *messaging.Message) {
				agent.subscribe(msg)
			},
			SubscriptionCancel: func(msg *messaging.Message) {
				agent.cancel(msg)
			},

		*/
	}
}()

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}
