package exchange

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func ExampleNewCtorMap() {
	m := std.NewSyncMap[string, messaging.NewAgentFunc]()
	name := ""
	t, ok := m.Load("")
	fmt.Printf("test:  Load(\"%v\") -> %v [ok:%v]\n", name, t, ok)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, func() messaging.Agent { return nil })
	t, ok = m.Load(name)
	fmt.Printf("test:  Load(\"%v\") -> %v [ok:%v]\n", name, t != nil, ok)

	//Output:
	//test:  Load("") -> <nil> [ok:false]
	//test:  Load("common:core:ctor/test") -> true [ok:true]

}
