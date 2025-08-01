package operations

import (
	"github.com/appellative-ai/core/messaging"
	"time"
)

// Origin map and host keys
const (
	RegistryHost1Key = "registry-host1"
	RegistryHost2Key = "registry-host2"
)

// ConfigLogging -
func ConfigLogging(log func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) {
	agent.configureLogging(log)
}

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	//agent.Message(messaging.StartupMessage)
}
