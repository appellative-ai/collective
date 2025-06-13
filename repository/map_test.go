package repository

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"reflect"
)

func ExampleNewMap() {
	m := newMap[string, messaging.NewAgent]()
	name := ""
	t := m.load("")
	fmt.Printf("test:  load(\"%v\") -> %v\n", name, t)

	//name = "common:core:ctor/test"
	//m.store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.store(name, func() messaging.Agent {
		fmt.Printf("in messaging.Agent\n")
		return nil
	})
	t = m.load(name)
	if t != nil {
		t()
	}
	fmt.Printf("test:  load(\"%v\") -> %v\n", name, reflect.TypeOf(t))

	//Output:
	//test:  load("") -> <nil>
	//test:  load("common:core:ctor/test") -> true

}
