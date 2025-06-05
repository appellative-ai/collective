package resourcetest

import (
	"fmt"
	"github.com/behavioral-ai/collective/resource"
	"github.com/behavioral-ai/core/messaging"
)

// NewResolver -
func NewResolver() *resource.Resolution {
	return &resource.Resolution{
		Representation: func(name string) (messaging.Content, *messaging.Status) {
			fmt.Printf("%v  -> %v\n", "representation", name)
			return messaging.Content{}, messaging.StatusOK()
		},
		AddRepresentation: func(name, author, contentType string, value any) *messaging.Status {
			fmt.Printf("%v  -> %v,%v,%v,%v\n", "addRepresentation", name, author, contentType, value)
			return messaging.StatusOK()
		},
		Context: func(name string) (messaging.Content, *messaging.Status) {
			fmt.Printf("%v  -> %v\n", "context", name)
			return messaging.Content{}, messaging.StatusOK()
		},
		AddContext: func(name, author, contentType string, value any) *messaging.Status {
			fmt.Printf("%v  -> %v,%v,%v,%v\n", "addContext", name, author, contentType, value)
			return messaging.StatusOK()
		},
	}
}
