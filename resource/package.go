package resource

import (
	"github.com/appellative-ai/core/messaging"
)

// Resolution - in the real world
// Can only add in current collective. An empty collective is assuming the local vs distributed
// How to handle local vs distributed
type Resolution struct {
	Representation func(name string) (messaging.Content, *messaging.Status)
	Context        func(name string) (messaging.Content, *messaging.Status)
}

// Resolver -
var Resolver = func() *Resolution {
	return &Resolution{
		Representation: func(name string) (messaging.Content, *messaging.Status) {
			return agent.getRepresentation(name)
		},
		Context: func(name string) (messaging.Content, *messaging.Status) {
			return messaging.Content{}, messaging.StatusOK()
		},
	}
}()
