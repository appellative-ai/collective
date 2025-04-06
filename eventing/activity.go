package eventing

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

type ActivityEvent struct {
	Agent   messaging.Agent
	Event   string
	Source  string
	Content any
}

func (a ActivityEvent) IsEmpty() bool {
	return a.Agent == nil
}

type ActivityFunc func(a ActivityEvent)

func Activity(a ActivityEvent) {
	uri := "<nil>"
	if a.Agent != nil {
		uri = a.Agent.Uri()
	}
	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), uri, a.Event, a.Source, a.Content)
}

/*
func NewActivityMessage(e ActivityEventm) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, ActivityEvent)
	m.SetContent(ContentTypeActivity, e)
	return m
}

func ActivityContent(msg *messaging.Message) ActivityItem {
	if msg == nil || msg.ContentType() != ContentTypeActivity || msg.Body == nil {
		return ActivityItem{}
	}
	if e, ok := msg.Body.(ActivityItem); ok {
		return e
	}
	return ActivityItem{}
}


*/

func NewActivityConfigMessage(fn ActivityFunc) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, ActivityConfigEvent)
	m.SetContent(ContentTypeActivityConfig, fn)
	return m
}

func ActivityConfigContent(msg *messaging.Message) ActivityFunc {
	if msg == nil || msg.ContentType() != ContentTypeActivityConfig || msg.Body == nil {
		return nil
	}
	if v, ok := msg.Body.(ActivityFunc); ok {
		return v
	}
	return nil
}
