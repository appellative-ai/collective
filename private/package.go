package private

import (
	"errors"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

const (
	ContentTypeInterface = "application/x-interface"
)

type Representation func(method, name, author, contentType string, value []byte) (messaging.Content, *messaging.Status)
type Context func(method, name, author, contentType string, value []byte) (messaging.Content, *messaging.Status)

type Thing func(method, name, cname, author string) *messaging.Status
type Relation func(method, name, cname, thing1, thing2, author string) *messaging.Status

// Interface -
type Interface struct {
	Representation Representation
	Context        Context
	Thing          Thing
	Relation       Relation
}

func NewInterface() *Interface {
	i := new(Interface)
	i.Representation = func(method, name, author, contentType string, value []byte) (messaging.Content, *messaging.Status) {
		if method == http.MethodGet {
			return messaging.Content{}, messaging.StatusNotFound()
		} else {
			return messaging.Content{}, messaging.StatusOK()

		}
	}
	i.Context = func(method, name, author, contentType string, value []byte) (messaging.Content, *messaging.Status) {
		if method == http.MethodGet {
			return messaging.Content{}, messaging.StatusNotFound()
		} else {
			return messaging.Content{}, messaging.StatusOK()

		}
	}
	i.Thing = func(method, name, cname, author string) *messaging.Status {
		return messaging.StatusOK()
	}
	i.Relation = func(method, name, cname, thing1, thing2, author string) *messaging.Status {
		return messaging.StatusOK()
	}
	return i
}

func NewInterfaceMessage(i *Interface) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	m.SetContent(ContentTypeInterface, i)
	return m
}

func InterfaceContent(m *messaging.Message) (*Interface, *messaging.Status) {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeInterface) {
		return nil, messaging.NewStatus(messaging.StatusInvalidContent, errors.New("invalid content"))
	}
	return messaging.New[*Interface](m.Content)
}
