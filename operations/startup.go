package operations

import (
	"errors"
	"github.com/appellative-ai/core/messaging"
	"time"
)

func (a *agentT) startup(collective string,
	registryHosts []string,
	status func(status any),
	exchange func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) error {

	if collective == "" {
		return errors.New("collective is required")
	}
	if len(registryHosts) == 0 {
		return errors.New("registryHosts are required")
	}
	a.messageStatus(status)
	a.messageExchange(exchange)

	// TODO: request collective host names and collective links

	return nil
}

func (a *agentT) messageStatus(status func(status any)) {
	if status == nil {
		status = logStatus
	}
	a.agents.Broadcast(messaging.NewConfigMessage(status))
}

func (a *agentT) messageExchange(exchange func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) {
	if exchange == nil {
		exchange = logExchange
	}
	a.agents.Broadcast(messaging.NewConfigMessage(exchange))
}
