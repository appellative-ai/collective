package resource

import (
	"github.com/behavioral-ai/core/messaging"
)

// Resolution - in the real world
// Can only add in current collective. An empty collective is assuming the local vs distributed
// How to handle local vs distributed
type Resolution struct {
	Representation    func(name string) (messaging.Content, *messaging.Status)
	AddRepresentation func(name, author, contentType string, value any) *messaging.Status

	Context    func(name string) (messaging.Content, *messaging.Status)
	AddContext func(name, author, contentType string, value any) *messaging.Status
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Representation: func(name string) (messaging.Content, *messaging.Status) {
			return agent.getRepresentation(name)
		},
		AddRepresentation: func(name, author, contentType string, value any) *messaging.Status {
			return agent.putRepresentation(name, author, contentType, value)
		},
		Context: func(name string) (messaging.Content, *messaging.Status) {
			return messaging.Content{}, messaging.StatusOK()
		},
		AddContext: func(name, author, contentType string, t any) *messaging.Status {
			return messaging.StatusOK()
		},
	}
}()
