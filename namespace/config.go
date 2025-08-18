package namespace

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"time"
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

	if e, ok2 := messaging.ConfigContent[func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)](m); ok2 && e != nil {
		if !a.running.Load() {
			a.logExchange = e
			return
		}
	}

}
