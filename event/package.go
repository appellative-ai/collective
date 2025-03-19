package event

import (
	http2 "github.com/behavioral-ai/core/http"
	"github.com/behavioral-ai/core/messaging"
)

const (
	ContentTypeNotify   = "application/notify"
	ContentTypeActivity = "application/activity"
	ContentTypeDispatch = "application/dispatch"

	NotifyEvent   = "event:notify"
	ActivityEvent = "event:activity"
	DispatchEvent = "event:dispatch"
)

// Agent - content resolution in the real world
var (
	Agent    messaging.Agent
	Exchange http2.Exchange
)

func init() {
	Exchange = http2.Do
	Agent = newAgent(nil, nil, nil)
	Agent.Run()
}
