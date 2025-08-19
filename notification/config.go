package notification

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"time"
)

func (a *agentT) configure(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	if e, ok := messaging.ConfigContent[rest.Exchange](m); ok && e != nil {
		if !a.running.Load() {
			a.exchange = e
		}
		return
	}
	if l, ok2 := messaging.ConfigContent[func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)](m); ok2 && l != nil {
		if !a.running.Load() {
			a.logFunc = l
		}
		return
	}
	if h, ok3 := messaging.ConfigContent[[]string](m); ok3 && len(h) > 0 {
		hosts := []string{h[0]}
		if len(h) > 1 {
			hosts = append(hosts, h[1])
		}
		a.hosts.Store(&hosts)
		return
	}
	if c, ok4 := messaging.ConfigContent[string](m); ok4 && c != "" {
		if !a.running.Load() {
			a.collective = c
		}
		return
	}

}
