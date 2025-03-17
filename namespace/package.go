package namespace

import (
	"github.com/behavioral-ai/core/messaging"
)

var (
	Agent messaging.Agent
	agent *agentT
)

func init() {
	agent = newAgent(nil)
	Agent = agent
}

// Adder - add
type Adder struct {
	Thing    func(nsName, author string) *messaging.Status
	Relation func(nsName1, nsName2, author string) *messaging.Status
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(nsName, author string) *messaging.Status {
			return agent.addThing(nsName, author)
		},
		Relation: func(nsName1, nsName2, author string) *messaging.Status {
			return agent.addRelation(nsName1, nsName2, author)
		},
	}
}()
