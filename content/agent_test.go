package content

import "fmt"

func ExampleNewAgent() {
	a := newAgent(nil)

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [unn:behavioral-ai.github.com:resiliency:agent/collective/content]

}
