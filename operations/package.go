package operations

import "github.com/behavioral-ai/core/messaging"

const (
// ServiceKind = "service"
)

// Origin map and host keys
const (
	RegistryHost1Key = "registry-host1"
	RegistryHost2Key = "registry-host2"
	CollectiveKey    = "collective"
	DomainKey        = "domain"
	RegionKey        = "region"
	ZoneKey          = "zone"
	SubZoneKey       = "sub-zone"
	HostKey          = "host"
	InstanceIdKey    = "instance-id"
)

// Service - in the real world
type Service struct {
	Message func(msg *messaging.Message)
	Advise  func(msg *messaging.Message)
	Trace   func(name, task, observation, action string)

	SubscribeCreate    func(msg *messaging.Message)
	SubscriptionCancel func(msg *messaging.Message)
}

// Serve -
var Serve = func() *Service {
	return &Service{
		Message: func(msg *messaging.Message) {
			agent.message(msg)
		},
		Advise: func(msg *messaging.Message) {
			agent.advise(msg)
		},
		SubscribeCreate: func(msg *messaging.Message) {
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
