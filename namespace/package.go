package namespace

import (
	"github.com/behavioral-ai/core/messaging"
)

// Adder - add
type Adder struct {
	Thing    func(collective, name, cname, author string) *messaging.Status
	Relation func(collective, name, cname, thing1, thing2, author string) *messaging.Status
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(collective, name, cname, author string) *messaging.Status {
			return agent.addThing(name, cname, "", author)
		},
		Relation: func(collective, name, cname, thing1, thing2, author string) *messaging.Status {
			return agent.addRelation(name, cname, thing1, thing2, "", author)
		},
	}
}()
