package private

import (
	"errors"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
	"net/http"
)

const (
	ContentTypeInterface = "application/x-interface"
)

type Representation func(method, name, author, contentType string, value []byte) (std.Content, *std.Status)
type Context func(method, name, author, contentType string, value []byte) (std.Content, *std.Status)

type Thing func(method, name, cname, author string) *std.Status
type Relation func(method, name, cname, thing1, thing2, author string) *std.Status

// Interface -
type Interface struct {
	Representation Representation
	Context        Context
	Thing          Thing
	Relation       Relation
}

func NewInterface() *Interface {
	i := new(Interface)
	i.Representation = func(method, name, author, contentType string, value []byte) (std.Content, *std.Status) {
		if method == http.MethodGet {
			return std.Content{}, std.StatusNotFound
		} else {
			return std.Content{}, std.StatusOK

		}
	}
	i.Context = func(method, name, author, contentType string, value []byte) (std.Content, *std.Status) {
		if method == http.MethodGet {
			return std.Content{}, std.StatusNotFound
		} else {
			return std.Content{}, std.StatusOK

		}
	}
	i.Thing = func(method, name, cname, author string) *std.Status {
		return std.StatusOK
	}
	i.Relation = func(method, name, cname, thing1, thing2, author string) *std.Status {
		return std.StatusOK
	}
	return i
}

func NewInterfaceMessage(i *Interface) *messaging.Message {
	m := messaging.NewMessage(messaging.ChannelControl, messaging.ConfigEvent)
	m.SetContent(ContentTypeInterface, i)
	return m
}

func InterfaceContent(m *messaging.Message) (*Interface, *std.Status) {
	if !messaging.ValidContent(m, messaging.ConfigEvent, ContentTypeInterface) {
		return nil, std.NewStatus(std.StatusInvalidContent, "", errors.New("invalid content"))
	}
	return std.New[*Interface](m.Content)
}
