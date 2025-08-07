package exchange

import (
	"fmt"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
)

func ExampleNewLinkMap() {
	m := std.NewSyncMap[string, rest.ExchangeLink]()
	name := ""
	t, ok := m.Load("")
	fmt.Printf("test:  Load(\"%v\") -> %v [ok:%v]\n", name, t, ok)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, func(next rest.Exchange) rest.Exchange { return nil })
	t, ok = m.Load(name)
	fmt.Printf("test:  Load(\"%v\") -> %v [ok:%v]\n", name, t != nil, ok)

	//Output:
	//test:  Load("") -> <nil> [ok:false]
	//test:  Load("common:core:ctor/test") -> true [ok:true]

}
