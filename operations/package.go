package operations

import "github.com/behavioral-ai/core/messaging"

const (
	ServiceKind = "service"
)

// Origin map and host keys
const (
	PrimaryHost   = "primary-host"
	SecondaryHost = "secondary-host"
	AppIdKey      = "app-id"
	RegionKey     = "region"
	ZoneKey       = "zone"
	SubZoneKey    = "sub-zone"
	HostKey       = "host"
	InstanceIdKey = "instance-id"
)

// Service - in the real world
type Service struct {
	Message            func(msg *messaging.Message)
	Advise             func(msg *messaging.Message)
	Subscribe          func(msg *messaging.Message)
	CancelSubscription func(msg *messaging.Message)
	Trace              func(name, task, observation, action string)
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
		Subscribe: func(msg *messaging.Message) {
			agent.subscribe(msg)
		},
		CancelSubscription: func(msg *messaging.Message) {
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
