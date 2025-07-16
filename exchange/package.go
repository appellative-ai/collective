package exchange

import (
	"github.com/appellative-ai/core/host"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
)

var (
	exchange  = messaging.NewExchange()
	ctor      = host.NewSyncMap[string, messaging.NewAgentFunc]()
	exHandler = host.NewSyncMap[string, rest.ExchangeLink]()
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
func RegisterConstructor(name string, fn messaging.NewAgentFunc) {
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
	exHandler.Store(name, fn)
}

// ExchangeHandler - return an exchange handler function
func ExchangeHandler(name string) func(next rest.Exchange) rest.Exchange {
	if name == "" {
		return nil
	}
	return exHandler.Load(name)
}
