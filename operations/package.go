package operations

import "github.com/behavioral-ai/core/messaging"

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}
