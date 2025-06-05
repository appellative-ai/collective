package repository

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleNewMessageMap() {
	m := newMessageMap[string, *messaging.Message]()
	name := ""
	t := m.get("")
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.store(name, messaging.NewMessage(messaging.ChannelControl, "test:name"))
	t = m.get(name)
	fmt.Printf("test:  get(\"%v\") -> %v\n", name, t)

	//Output:
	//test:  get("") -> <nil>
	//test:  get("common:core:ctor/test") -> [chan:ctrl] [from:] [to:] [test:name]

}
