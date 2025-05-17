package repository

import "github.com/behavioral-ai/core/messaging"

// NewAgent - agent constructor
type NewAgent func() messaging.Agent

var (
	exchange = messaging.NewExchange()
	cmap     = newMap()
)

func Register(a messaging.Agent) error {
	return exchange.Register(a)
}

func Agent(name string) messaging.Agent {
	agent := exchange.Get(name)
	if agent != nil {
		return agent
	}
	agent = Constructor(name)
	if agent == nil {
		return nil
	}
	Register(agent)
	return agent
}

func Message(m *messaging.Message) {
	exchange.Message(m)
}

func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}

func RegisterConstructor(name string, fn NewAgent) {
	if name == "" || fn == nil {
		return
	}
	cmap.put(name, fn)
}

func Constructor(name string) messaging.Agent {
	fn := cmap.get(name)
	if fn != nil {
		return fn()
	}
	return nil
}
