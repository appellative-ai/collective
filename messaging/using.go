package messaging

import "github.com/behavioral-ai/core/messaging"

const (
	ContentTypeUsing = "application/using"
)

type UsingRecord struct {
	Collective string
	Hosts      []string
}

func NewUsingMessage(ur UsingRecord) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	m.SetContent(ContentTypeUsing, ur)
	return m
}

func UsingContent(m *messaging.Message) (UsingRecord, bool) {
	if m.Name != messaging.ConfigEvent || m.ContentType() != ContentTypeUsing {
		return UsingRecord{}, false
	}
	if v, ok := m.Body.(UsingRecord); ok {
		return v, true
	}
	return UsingRecord{}, false
}

// Using - configure collectives in use
/*
func Using(ur messaging.UsingRecord) {
	if ur.Collective == "" || len(ur.Hosts) == 0 {
		return
	}
	agent.Message(messaging.NewUsingMessage(ur))
}


*/
