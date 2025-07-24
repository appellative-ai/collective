package resolution

import (
	"github.com/appellative-ai/core/std"
)

// Interface - resolution in the real world
// Can only add in current collective. An empty collective is assuming the local vs distributed
// How to handle local vs distributed
type Interface struct {
	Representation func(name string) (std.Content, *std.Status)
	Context        func(name string) (std.Content, *std.Status)
}

// Resolver -
var Resolver = func() *Interface {
	return &Interface{
		Representation: func(name string) (std.Content, *std.Status) {
			return agent.getRepresentation(name)
		},
		Context: func(name string) (std.Content, *std.Status) {
			return std.Content{}, std.StatusOK
		},
	}
}()
