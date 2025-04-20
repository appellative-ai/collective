package namespace

import (
	"github.com/behavioral-ai/core/messaging"
)

// Adder - add
type Adder struct {
	Thing    func(nsName, cName, author string) *messaging.Status
	Relation func(nsName, cName, thing1, thing2, author string) *messaging.Status
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(nsName, cName, author string) *messaging.Status {
			return agent.addThing(nsName, cName, author)
		},
		Relation: func(nsName, cName, nsThing1, nsThing2, author string) *messaging.Status {
			return agent.addRelation(nsName, cName, nsThing1, nsThing2, author)
		},
	}
}()
