package timeseries

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

var (
	Agent messaging.Agent
	agent *agentT
)

func init() {
	agent = newAgent(nil)
	Agent = agent
}

// Interface -
type Interface struct {
	Rollup func(origin Origin) *messaging.Status
	Add    func(events []Event) *messaging.Status
}

// Functions -
var Functions = func() *Interface {
	return &Interface{
		Rollup: func(origin Origin) *messaging.Status {
			return agent.rollup(origin)
		},
		Add: func(events []Event) *messaging.Status {
			if len(events) == 0 {
				return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid argument events are empty")), agent.Uri())
			}
			return agent.addEvents(events)
		},
	}
}()
