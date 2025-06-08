package repository

import (
	"fmt"
	"github.com/behavioral-ai/core/rest"
)

func ExampleNewLinkMap() {
	m := newLinkMap[string, rest.ExchangeLink]()
	name := ""
	t := m.load("")
	fmt.Printf("test:  load(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.store(name, func(next rest.Exchange) rest.Exchange { return nil })
	t = m.load(name)
	fmt.Printf("test:  load(\"%v\") -> %v\n", name, t != nil)

	//Output:
	//test:  load("") -> <nil>
	//test:  load("common:core:ctor/test") -> true

}
