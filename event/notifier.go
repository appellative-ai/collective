package event

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

type NotifyItem interface {
	AgentId() string
	Type() string
	Status() string
	Message() string
}

type NotifyFunc func(e NotifyItem)

func Notify(e NotifyItem) {
	fmt.Printf("notify-> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), e.AgentId(), e.Type(), e.Status(), e.Message())
}

type Notifier interface {
	Notify(status *messaging.Status)
}

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
