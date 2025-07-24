package resolution

import "fmt"

func ExampleNewAgent() {
	a := newAgent()

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [common:core:agent/resolution/collective]

}
