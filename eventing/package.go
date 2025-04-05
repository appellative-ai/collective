package eventing

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
)

const (
	ContentTypeNotify   = "application/notify"
	ContentTypeActivity = "application/activity"

	NotifyEvent   = "eventing:notify"
	ActivityEvent = "eventing:activity"
)

// Agent - content resolution in the real world
var (
	Agent    messaging.Agent
	Exchange httpx.Exchange
)

func init() {
	Exchange = httpx.Do
	Agent = newAgent(nil, nil)
	Agent.Message(messaging.StartupMessage)
}
