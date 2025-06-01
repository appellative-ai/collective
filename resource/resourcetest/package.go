package resourcetest

import (
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/core/messaging"
)

var Resolver = func() *resource.Resolution {
	return &resource.Resolution{
		Representation: func(name, fragment string) (resource.Content, *messaging.Status) {
			return resource.Content{}, messaging.StatusOK()
		},
		AddRepresentation: func(name, fragment, author string, value any) *messaging.Status {
			return messaging.StatusOK()
		},
		Context: func(name string) (resource.Content, *messaging.Status) {
			return resource.Content{}, messaging.StatusOK()
		},
		AddContext: func(name, author string, ct resource.Content) *messaging.Status {
			return messaging.StatusOK()
		},
	}
}()
