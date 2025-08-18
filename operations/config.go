package operations

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
)

func (a *agentT) configure(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	if ex, ok := messaging.ConfigContent[rest.Exchange](m); ok && ex != nil {
		if !a.running.Load() {
			a.exchange = ex
			return
		}
	}

}
