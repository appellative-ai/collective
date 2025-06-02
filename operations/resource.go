package operations

import (
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/core/messaging"
)

func representation(method, name, fragment, author string, value any) (exchange.Content, *messaging.Status) {
	return exchange.Content{}, messaging.StatusOK()
}

func context(method, name, author string, ct exchange.Content) (exchange.Content, *messaging.Status) {
	return exchange.Content{}, messaging.StatusOK()
}
