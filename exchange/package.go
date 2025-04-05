package exchange

import "github.com/behavioral-ai/core/messaging"

var (
	exchange = messaging.NewExchange()
)

func Message(m *messaging.Message) {
	exchange.Send(m)
}

func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}

func Register(a messaging.Agent) error {
	return exchange.Register(a)
}
