package operations

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/eventing"
	"github.com/behavioral-ai/collective/namespace"
	"github.com/behavioral-ai/collective/timeseries"
	"github.com/behavioral-ai/core/messaging"
	"net/url"
	"strconv"
)

var (
	Agent = New()
)

const (
	ContentPath    = "/collective/content"
	ActivityPath   = "/collective/activity"
	NotifyPath     = "/collective/notify"
	NamespacePath  = "/collective/namespace"
	TimeseriesPath = "/collective/timeseries"
	NsNameKey      = "name"
	VersionKey     = "ver"
)

var (
	HostName string
)

func Initialize(hostName string) {
	HostName = hostName
}

func ContentURL(nsName string, version int) string {
	v := make(url.Values)
	v.Set(NsNameKey, nsName)
	v.Set(VersionKey, strconv.Itoa(version))
	return ContentPath + "?" + v.Encode()
}

func ActivityURL() string {
	return ActivityPath
}

func NotifyURL() string {
	return NotifyPath
}

// Configure - configure all agents
func Configure(m *messaging.Message) {
	if m.Event() == messaging.ConfigEvent && m.ContentType() == messaging.ContentTypeMap {
		content.Agent.Message(m)
		eventing.Agent.Message(m)
		namespace.Agent.Message(m)
		timeseries.Agent.Message(m)
	}
}

// Message - operations agent messaging
func Message(event string) error {
	switch event {
	case messaging.StartupEvent:
		if Agent == nil {
			Agent = New()
			Agent.Message(messaging.StartupMessage)
		}
	case messaging.ShutdownEvent:
		if Agent != nil {
			Agent.Message(messaging.ShutdownMessage)
			Agent = nil
		}
	case messaging.PauseEvent:
		if Agent != nil {
			Agent.Message(messaging.PauseMessage)
		}
	case messaging.ResumeEvent:
		if Agent != nil {
			Agent.Message(messaging.ResumeMessage)
		}
	default:
		return errors.New(fmt.Sprintf("operations.Message() -> [%v] [%v]", "error: invalid eventing", event))
	}
	return nil
}
