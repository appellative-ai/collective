package repository

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func ExampleNewMessageMap() {
	m := std.NewSyncMap[string, *messaging.Message]()
	name := ""
	t := m.Load("")
	fmt.Printf("test:  Load(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, messaging.NewMessage(messaging.ChannelControl, "test:name"))
	t = m.Load(name)
	fmt.Printf("test:  Load(\"%v\") -> %v\n", name, t)

	//Output:
	//test:  Load("") -> <nil>
	//test:  Load("common:core:ctor/test") -> [chan:ctrl] [from:] [to:[]] [test:name]

}
