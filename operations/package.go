package operations

import "github.com/behavioral-ai/core/messaging"

// Register - used for collective agents
func Register(agent messaging.Agent) {
	if validAgent(agent) {
		exchange.Register(agent)
	}
}

// Configure - configure collective agents
// TODO - configuration error handling
func Configure(m *messaging.Message) {
	exchange.Broadcast(m)
	exchange.Broadcast(messaging.StartupMessage)
}

// Shutdown - shutdown collective agents
func Shutdown() {
	exchange.Broadcast(messaging.ShutdownMessage)
}
