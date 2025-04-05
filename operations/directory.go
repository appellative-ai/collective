package operations

import (
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/core/messaging"
)

const (
	contentName    = "resiliency:agent/behavioral-ai/collective/content"
	namespaceName  = "resiliency:agent/behavioral-ai/collective/namespace"
	timeseriesName = "resiliency:agent/behavioral-ai/collective/timeseries"
)

var (
	exchange = messaging.NewExchange()
	agent    = newAgent(eventing.Agent)
)

func validAgent(agent messaging.Agent) bool {
	if agent.Uri() == contentName || agent.Uri() == namespaceName || agent.Uri() == timeseriesName {
		return true
	}
	return false
}
