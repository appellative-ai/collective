package eventing

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

type NotifyEvent interface {
	AgentId() string
	Type() string
	Status() string
	Message() string
	RequestId() string
}

type NotifyFunc func(e NotifyEvent)

func OutputNotify(e NotifyEvent) {
	fmt.Printf("notify-> %v [%v] [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), e.AgentId(), e.Type(), e.RequestId(), e.Status(), e.Message())
}

/*
type Notifier interface {
	Notify(status *messaging.Status)
}

*/

/*
func NewNotifyMessage(e NotifyItem) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, NotifyEvent)
	m.SetContent(ContentTypeNotify, e)
	return m
}

func NotifyContent(msg *messaging.Message) NotifyItem {
	if msg == nil || msg.ContentType() != ContentTypeNotify || msg.Body == nil {
		return nil
	}
	if e, ok := msg.Body.(NotifyItem); ok {
		return e
	}
	return nil
}


*/

func NewNotifyConfigMessage(fn NotifyFunc) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, NotifyConfigEvent)
	m.SetContent(ContentTypeNotifyConfig, fn)
	return m
}

func NotifyConfigContent(m *messaging.Message) NotifyFunc {
	if m.Event() != NotifyConfigEvent || m.ContentType() != ContentTypeNotifyConfig {
		return nil
	}
	if v, ok := m.Body.(NotifyFunc); ok {
		return v
	}
	return nil
}
