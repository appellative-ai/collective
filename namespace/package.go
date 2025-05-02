package namespace

import (
	"github.com/behavioral-ai/core/messaging"
)

// Adder - add
type Adder struct {
	Thing    func(name, cname, author string) *messaging.Status
	Relation func(name, cname, thing1, thing2, author, authority string) *messaging.Status
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(name, cname, author string) *messaging.Status {
			return agent.addThing(name, cname, "", author)
		},
		Relation: func(name, cname, thing1, thing2, author, authority string) *messaging.Status {
			return agent.addRelation(name, cname, thing1, thing2, authority, author)
		},
	}
}()
