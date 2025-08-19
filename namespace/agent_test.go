package namespace

import (
	"fmt"
	"sync/atomic"
)

func ExampleNewAgent() {
	a := newAgent()

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/namespace/collective]

}

func ExampleHosts() {
	var hosts atomic.Pointer[[]string]

	hosts.Store(&[]string{"host1", "host2"})

	fmt.Printf("test: atomic.Pointer() -> [%v][%v] [%v]\n", hosts.Load(), (*hosts.Load())[0], (*hosts.Load())[1])

	//Output:
	//test: atomic.Pointer() -> [&[host1 host2]][host1] [host2]

}
