package operations

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

const (
// ServiceKind = "service"
)

// Origin map and host keys
const (
	RegistryHost1Key = "registry-host1"
	RegistryHost2Key = "registry-host2"
)

// Interface - service interface
// TODO: determine which additional namespace requests to support
type Interface struct {
	Ping func() *std.Status
}

// Service -
var Service = func() *Interface {
	return &Interface{
		Ping: func() *std.Status {
			return std.StatusOK
		},
	}
}()

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}
