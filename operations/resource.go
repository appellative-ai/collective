package operations

import (
	"github.com/behavioral-ai/core/messaging"
)

func representation(method, name, fragment, author string, value any) (messaging.Content, *messaging.Status) {
	return messaging.Content{}, messaging.StatusOK()
}

func context(method, name, author, ct string, value any) (messaging.Content, *messaging.Status) {
	return messaging.Content{}, messaging.StatusOK()
}
