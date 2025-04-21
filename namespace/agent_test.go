package namespace

import "fmt"

func ExampleNewAgent() {
	a := newAgent(nil)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/namespace]

}
