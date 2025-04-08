package exchange

import "github.com/behavioral-ai/core/messaging"

var (
	exchange = messaging.NewExchange()
)

func Register(a messaging.Agent) error {
	return exchange.Register(a)
}

func Agent(uri string) messaging.Agent {
	return exchange.Get(uri)
}

func Send(m *messaging.Message) {
	exchange.Send(m)
}

func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}
