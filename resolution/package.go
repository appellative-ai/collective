package resolution

import (
	"github.com/appellative-ai/core/std"
)

// Interface - resolution in the real world
// Can only add in current collective. An empty collective is assuming the local vs distributed
// How to handle local vs distributed
type Interface struct {
	Representation    func(name string) (std.Content, *std.Status)
	AddRepresentation func(name, author, contentType string, value any) *std.Status

	Context    func(name string) (std.Content, *std.Status)
	AddContext func(name, author, contentType string, value any) *std.Status
}

// Resolver -
var Resolver = func() *Interface {
	return &Interface{
		Representation: func(name string) (std.Content, *std.Status) {
			return agent.getRepresentation(name)
		},
		AddRepresentation: func(name, author, contentType string, value any) *std.Status {
			return std.StatusOK
		},
		Context: func(name string) (std.Content, *std.Status) {
			return std.Content{}, std.StatusOK
		},
		AddContext: func(name, author, contentType string, t any) *std.Status {
			return std.StatusOK
		},
	}
}()
