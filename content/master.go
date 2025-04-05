package content

import (
	"github.com/behavioral-ai/core/messaging"
)

// master attention
func masterAttend(agent *agentT) {
	paused := false
	if paused {
	}

	for {
		select {
		case msg := <-agent.master.C:
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.masterFinalize()
				return
			default:
			}
		default:
		}
	}
}
