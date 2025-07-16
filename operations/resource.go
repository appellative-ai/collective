package operations

import (
	"github.com/appellative-ai/core/messaging"
)

func representation(method, name, author, contentType string, value []byte) (messaging.Content, *messaging.Status) {
	return messaging.Content{}, messaging.StatusOK()
}

func context(method, name, author, contentType string, value []byte) (messaging.Content, *messaging.Status) {
	return messaging.Content{}, messaging.StatusOK()
}
