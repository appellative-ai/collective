package repository

import (
	"fmt"
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleNewCtorMap() {
	m := host.NewSyncMap[string, messaging.NewAgent]()
	name := ""
	t := m.Load("")
	fmt.Printf("test:  Load(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, func() messaging.Agent { return nil })
	t = m.Load(name)
	fmt.Printf("test:  Load(\"%v\") -> %v\n", name, t != nil)

	//Output:
	//test:  Load("") -> <nil>
	//test:  Load("common:core:ctor/test") -> true

}
