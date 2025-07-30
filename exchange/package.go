package exchange

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
)

var (
	exchange  = messaging.NewExchange()
	ctor      = std.NewSyncMap[string, messaging.NewAgentFunc]()
	exHandler = std.NewSyncMap[string, rest.ExchangeLink]()
)

// Register - register an agent
func Register(a messaging.Agent) error {
	return exchange.Register(a)
}

// NewAgent - construct a new agent
func NewAgent(name string) messaging.Agent {
	fn := ctor.Load(name)
	if fn != nil {
		return fn()
	}
	return nil
}

// Agent - get an agent
func Agent(name string) messaging.Agent {
	return exchange.Get(name)
	/*
		if agent != nil {
			return agent
		}
		agent = NewAgent(name)
		if agent == nil {
			return nil
		}
		Register(agent)
		return agent

	*/
}

// AgentT - get typed agent
func AgentT[T any](name string) (t T, ok bool) {
	a := Agent(name)
	if a == nil {
		return
	}
	t, ok = any(a).(T)
	return
}

// Exists -
func Exists(name string) bool {
	return ctor.Load(name) != nil
}

// Message - message an agent
func Message(m *messaging.Message) bool {
	return exchange.Message(m)
}

// Broadcast - broadcast a message
func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}

// RegisterConstructor - register a new agent function, used for local agent assignments
func RegisterConstructor(name string, fn messaging.NewAgentFunc) {
	if name == "" || fn == nil {
		return
	}
	ctor.Store(name, fn)
	Register(fn())
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
