package event

import "fmt"

func ExampleNewAgent() {
	a := newAgent(nil)

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [resiliency:agent/behavioral-ai/collective/event]

}
