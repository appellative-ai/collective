package event

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

type ActivityItem struct {
	Agent   messaging.Agent
	Event   string
	Source  string
	Content any
}

func (a ActivityItem) IsEmpty() bool {
	return a.Agent == nil
}

type ActivityFunc func(a ActivityItem)

func Activity(a ActivityItem) {
	uri := "<nil>"
	if a.Agent != nil {
		uri = a.Agent.Uri()
	}
	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), uri, a.Event, a.Source, a.Content)
}

func NewActivityMessage(e ActivityItem) *messaging.Message {
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
