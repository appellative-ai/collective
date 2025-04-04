package timeseries

import "github.com/behavioral-ai/core/messaging"

const (
	LoadEvent        = "event:load"
	ContentTypeEvent = "application/event"
)

func NewLoadMessage(e []Event) *messaging.Message {
	m := messaging.NewMessage(messaging.Control, LoadEvent)
	m.SetContent(ContentTypeEvent, e)
	return m
}

func LoadContent(m *messaging.Message) []Event {
	if m.Event() != LoadEvent || m.ContentType() != ContentTypeEvent {
		return nil
	}
	if e, ok := m.Body.([]Event); ok {
		return e
	}
	return nil
}
