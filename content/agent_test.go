package content

import "fmt"

func ExampleNewAgent() {
	a := newAgent()

	fmt.Printf("test: newHttpAgent() -> [%v]\n", a)

	//Output:
	//test: newHttpAgent() -> [collective:agent/content]

}
