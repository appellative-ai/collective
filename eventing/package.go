package eventing

import (
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/messaging"
)

const (
	ContentTypeNotify   = "application/notify"
	ContentTypeActivity = "application/activity"
	ContentTypeDispatch = "application/dispatch"

	NotifyEvent   = "eventing:notify"
	ActivityEvent = "eventing:activity"
	DispatchEvent = "eventing:dispatch"
)

// Agent - content resolution in the real world
var (
	Agent    messaging.Agent
	Exchange httpx.Exchange
)

func init() {
	Exchange = httpx.Do
	Agent = newAgent(nil, nil, nil)
	Agent.Message(messaging.StartupMessage)
}
