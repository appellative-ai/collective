package repository

import "github.com/behavioral-ai/core/messaging"

var (
	exchange = messaging.NewExchange()
	ctorMap  = newMap()
)

// Register - register an agent
func Register(a messaging.Agent) error {
	return exchange.Register(a)
}

// Agent - get an agent from the exchange, constructing the agent if necessary
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

// Message - message an agent, using the message To as the agent name
func Message(m *messaging.Message) {
	if m == nil {
		return
	}
	agent := Agent(m.To())
	if agent != nil {
		agent.Message(m)
	}
	//exchange.Message(m)
}

// Broadcast - broadcast a message to all registered agents
func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}

// RegisterConstructor - register a new agent function
func RegisterConstructor(name string, fn messaging.NewAgent) {
	if name == "" || fn == nil {
		return
	}
	ctorMap.put(name, fn)
}

// Constructor - construct a new agent
func Constructor(name string) messaging.Agent {
	fn := ctorMap.get(name)
	if fn != nil {
		return fn()
	}
	return nil
}
