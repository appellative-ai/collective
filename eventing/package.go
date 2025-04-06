package eventing

import (
	"github.com/behavioral-ai/collective/operations"
	"github.com/behavioral-ai/core/messaging"
)

//NotifyEvent         = "eventing:notify"
//ActivityEvent       = "eventing:activity"
//ContentTypeNotify         = "application/notify"
//ContentTypeActivity       = "application/activity"

const (
	ContentTypeNotifyConfig   = "application/notify-config"
	ContentTypeActivityConfig = "application/activity-config"
	NotifyConfigEvent         = "eventing:notify-config"
	ActivityConfigEvent       = "eventing:activity-config"
)

type Agent interface {
	messaging.Agent
	Notify(e NotifyEvent)
	AddActivity(e ActivityEvent)
}

var (
	Handler Agent
)

func init() {
	Handler = newAgent()
	operations.Register(Handler)
}
