package repository

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleRegisterConstructor() {
	fn := func() messaging.Agent {
		fmt.Printf("test: RegisterConstructor() [initial]\n")
		return nil
	}
	fn2 := func() messaging.Agent {
		fmt.Printf("test: RegisterConstructor() [replaced]\n")
		return nil
	}
	fn()
	fn2()

	RegisterConstructor("one", fn)
	NewAgent("one")

	RegisterConstructor("one", fn2)
	NewAgent("one")

	name := "one"
	fmt.Printf("test: Exist(\"%v\") %v\n", name, Exists(name))
	name = "invalid"
	fmt.Printf("test: Exist(\"%v\") %v\n", name, Exists(name))

	//Output:
	//test: RegisterConstructor() [initial]
	//test: RegisterConstructor() [replaced]
	//test: RegisterConstructor() [initial]
	//test: RegisterConstructor() [replaced]
	//test: Exist("one") true
	//test: Exist("invalid") false

}
