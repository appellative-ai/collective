package namespace

import "fmt"

func ExampleNewAgent() {
	a := newAgent(nil)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [resiliency:agent/behavioral-ai/collective/namespace]

}
