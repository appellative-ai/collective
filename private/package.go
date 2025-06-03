package private

import (
	"github.com/behavioral-ai/core/messaging"
)

const (
	ContentTypeInterface = "application/x-interface"
)

type Representation func(method, name, fragment, author string, value any) (string, string, any, *messaging.Status)
type Context func(method, name, author, ct string, value any) (string, string, any, *messaging.Status)

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
	m.SetContent(ContentTypeInterface, i)
	return m
}

func InterfaceContent(m *messaging.Message) Interface {
	if m.Name != messaging.ConfigEvent || m.ContentType() != ContentTypeInterface {
		return Interface{}
	}
	if cfg, ok := m.Body.(Interface); ok {
		return cfg
	}
	return Interface{}
}
