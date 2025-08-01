package resolutiontest

import (
	"fmt"
	"github.com/appellative-ai/collective/resolution"
	"github.com/appellative-ai/core/std"
)

// NewResolver -
func NewResolver() *resolution.Interface {
	return &resolution.Interface{
		Representation: func(name string) (std.Content, *std.Status) {
			fmt.Printf("%v  -> %v\n", "representation", name)
			return std.Content{}, std.StatusOK
		},
		AddRepresentation: func(name, author, contentType string, value any) *std.Status {
			fmt.Printf("%v  -> %v,%v,%v,%v\n", "addRepresentation", name, author, contentType, value)
			return std.StatusOK
		},
		Context: func(name string) (std.Content, *std.Status) {
			fmt.Printf("%v  -> %v\n", "context", name)
			return std.Content{}, std.StatusOK
		},
		AddContext: func(name, author, contentType string, value any) *std.Status {
			fmt.Printf("%v  -> %v,%v,%v,%v\n", "addContext", name, author, contentType, value)
			return std.StatusOK
		},
	}
}
