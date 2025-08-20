package operations

import (
	"errors"
	"github.com/appellative-ai/core/messaging"
	"log"
	"time"
)

func (a *agentT) startup() error {
	if a.origin == nil {
		return errors.New("origin is required")
	}
	s := a.state.Load()
	if len(s.registryHosts) == 0 || s.registryHosts[0] == "" {
		return errors.New("registry hosts are required")
	}

	a.messageExchange(a.logFunc)

	// TODO: request collective host names and collective links.
	//       configure agents hosts and collective for notifications

	return nil
}

func (a *agentT) messageExchange(logFunc func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) {
	if logFunc == nil {
		logFunc = defaultLog
	}
	a.agents.Broadcast(messaging.NewConfigMessage(logFunc))
}

func defaultLog(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	log.Printf("%v %v %v %v %v %v\n", start, duration, route, req, resp, timeout)
}
