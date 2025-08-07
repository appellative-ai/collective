package repository

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func ExampleNewMessageMap() {
	m := std.NewSyncMap[string, *messaging.Message]()
	name := ""
	t, ok := m.Load("")
	fmt.Printf("test:  Load(\"%v\") -> %v [ok:%v]\n", name, t, ok)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, messaging.NewMessage(messaging.ChannelControl, "test:name"))
	t, ok = m.Load(name)
	fmt.Printf("test:  Load(\"%v\") -> %v [ok:%v]\n", name, t, ok)

	//Output:
	//test:  Load("") -> <nil> [ok:false]
	//test:  Load("common:core:ctor/test") -> [chan:ctrl] [from:] [to:[]] [test:name] [ok:true]
	
}
