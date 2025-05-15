package config

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
	if m.Event() != messaging.ConfigEvent || m.ContentType() != ContentTypeUsing {
		return UsingRecord{}, false
	}
	if v, ok := m.Body.(UsingRecord); ok {
		return v, true
	}
	return UsingRecord{}, false
}

// Using - configure collectives in use
/*
func Using(ur config.UsingRecord) {
	if ur.Collective == "" || len(ur.Hosts) == 0 {
		return
	}
	agent.Message(config.NewUsingMessage(ur))
}


*/
