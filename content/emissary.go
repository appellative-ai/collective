package content

import (
	"github.com/behavioral-ai/core/messaging"
)

// emissary attention
func emissaryAttend(agent *agentT) {
	var paused = false
	if paused {
	}
	agent.ticker.Start(-1)
	for {
		select {
		case <-agent.ticker.C():
		default:
		}
		select {
		case msg := <-agent.emissary.C:
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.emissaryFinalize()
				return
			default:
			}
		default:
		}
	}
}
