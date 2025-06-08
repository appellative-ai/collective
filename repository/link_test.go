package repository

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleNewLinkMap() {
	m := newCtorMap[string, messaging.NewAgent]()
	name := ""
	t := m.get("")
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.store(name, func() messaging.Agent { return nil })
	t = m.get(name)
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t != nil)

	//Output:
	//test:  get("") -> <nil>
	//test:  get("common:core:ctor/test") -> true

}
