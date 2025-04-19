package operations

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	contentName   = "resiliency:agent/behavioral-ai/collective/content"
	namespaceName = "resiliency:agent/behavioral-ai/collective/namespace"
)

var (
	exchange = messaging.NewExchange()
	agent    = newAgent()
)

func validAgent(agent messaging.Agent) bool {
	if agent.Uri() == contentName || agent.Uri() == namespaceName {
		return true
	}
	return false
}
