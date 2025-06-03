package operations

import (
	"github.com/behavioral-ai/core/messaging"
)

func representation(method, name, fragment, author string, value any) (string, string, any, *messaging.Status) {
	return "", "", nil, messaging.StatusOK()
}

func context(method, name, author, ct string, value any) (string, string, any, *messaging.Status) {
	return "", "", nil, messaging.StatusOK()
}
