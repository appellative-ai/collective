package operations

import (
	"errors"
	"github.com/appellative-ai/core/messaging"
	"log"
	"time"
)

func (a *agentT) startup(collective string,
	registryHost1 string,
	registryHost2 string,
	exchange func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) error {

	if collective == "" {
		return errors.New("collective is required")
	}
	if registryHost1 == "" && registryHost2 == "" {
		return errors.New("registryHosts are required")
	}
	a.messageExchange(exchange)

	// TODO: request collective host names and collective links

	return nil
}

func (a *agentT) messageExchange(exchange func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) {
	if exchange == nil {
		exchange = logExchange
	}
	a.agents.Broadcast(messaging.NewConfigMessage(exchange))
}

func logExchange(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	log.Printf("%v %v %v %v %v %v\n", start, duration, route, req, resp, timeout)
}
