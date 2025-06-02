package operations

import (
	"github.com/behavioral-ai/collective/private"
	"github.com/behavioral-ai/core/messaging"
)

func representation(method, name, fragment, author string, value any) (private.Content, *messaging.Status) {
	return private.Content{}, messaging.StatusOK()
}

func context(method, name, author string, ct private.Content) (private.Content, *messaging.Status) {
	return private.Content{}, messaging.StatusOK()
}
