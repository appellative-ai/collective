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

// Service - in the real world
type Service struct {
	Message func(msg *messaging.Message) bool
	Advise  func(msg *messaging.Message) *messaging.Status
	Trace   func(name, task, observation, action string)

	SubscriptionCreate func(msg *messaging.Message)
	SubscriptionCancel func(msg *messaging.Message)
}

// Serve -
var Serve = func() *Service {
	return &Service{
		Message: func(msg *messaging.Message) bool {
			agent.message(msg)
			return true
		},
		Advise: func(msg *messaging.Message) *messaging.Status {
			agent.advise(msg)
			return messaging.StatusOK()
		},
		SubscriptionCreate: func(msg *messaging.Message) {
			agent.subscribe(msg)
		},
		SubscriptionCancel: func(msg *messaging.Message) {
			agent.cancel(msg)
		},
		Trace: func(name, task, observation, action string) {
			agent.trace(name, task, observation, action)
		},
	}
}()

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}
