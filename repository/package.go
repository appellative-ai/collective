package repository

import (
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/rest"
)

var (
	exchange = messaging.NewExchange()
	ctor     = host.NewSyncMap[string, messaging.NewAgent]()
	link     = host.NewSyncMap[string, rest.ExchangeLink]()
	msg      = host.NewSyncMap[string, *messaging.Message]()
)

// Register - register an agent
func Register(a messaging.Agent) error {
	return exchange.Register(a)
}

// Agent - get an agent
func Agent(name string) messaging.Agent {
	agent := exchange.Get(name)
	if agent != nil {
		return agent
	}
	agent = NewAgent(name)
	if agent == nil {
		return nil
	}
	Register(agent)
	return agent
}

// Message - message an agent
func Message(m *messaging.Message) bool {
	return exchange.Message(m)
}

// Broadcast - broadcast a message
func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}

// RegisterConstructor - register a new agent function
func RegisterConstructor(name string, fn messaging.NewAgent) {
	if name == "" || fn == nil {
		return
	}
	ctor.Store(name, fn)
}

// NewAgent - construct a new agent
func NewAgent(name string) messaging.Agent {
	fn := ctor.Load(name)
	if fn != nil {
		return fn()
	}
	return nil
}

func Exists(name string) bool {
	return ctor.Load(name) != nil
}

// RegisterExchangeHandler - register a new handler function
func RegisterExchangeHandler(name string, fn func(next rest.Exchange) rest.Exchange) {
	if name == "" || fn == nil {
		return
	}
	link.Store(name, fn)
}

// ExchangeHandler - return an exchange handler function
func ExchangeHandler(name string) func(next rest.Exchange) rest.Exchange {
	if name == "" {
		return nil
	}
	return link.Load(name)
}

// GetMessage - get a message
func GetMessage(name string) *messaging.Message {
	return msg.Load(name)
}

// StoreMessage - store a message
func StoreMessage(m *messaging.Message) {
	msg.Store(m.Name, m)
}

/*
// ModifyMessage - modify a message
func ModifyMessage(m *messaging.Message) {
	msg.modify(m)
}

// DeleteMessage - delete a message
func DeleteMessage(name string) {
	msg.delete(name)
}


*/
