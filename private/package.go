package private

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	ContentTypeInterface = "application/x-interface"
)

type Representation func(method, name, fragment, author string, value any) (messaging.Content, *messaging.Status)
type Context func(method, name, author, ct string, value any) (messaging.Content, *messaging.Status)

type Thing func(method, name, cname, author string) *messaging.Status
type Relation func(method, name, cname, thing1, thing2, author string) *messaging.Status

// Interface -
type Interface struct {
	Collective string
	Rep        Representation
	Ctx        Context
	Th         Thing
	Rel        Relation
}

func NewInterfaceMessage(i Interface) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	m.SetContent(ContentTypeInterface, "", i)
	return m
}

func InterfaceContent(m *messaging.Message) Interface {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeInterface) {
		return Interface{}
	}
	if cfg, ok := m.Content.Value.(Interface); ok {
		return cfg
	}
	return Interface{}
}
