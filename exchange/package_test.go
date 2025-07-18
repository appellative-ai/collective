package exchange

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
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
	//test: RegisterConstructor() [initial]
	//test: RegisterConstructor() [replaced]
	//test: RegisterConstructor() [replaced]
	//test: Exist("one") true
	//test: Exist("invalid") false

}

func ExampleRegisterExchange() {
	fn := func(next rest.Exchange) rest.Exchange {
		fmt.Printf("test: Exchange handler function\n")
		return nil
	}
	name := "test"
	RegisterExchangeHandler(name, fn)

	l := ExchangeHandler(name)
	fmt.Printf("test: ExchangeHandler() -> %v\n", l != nil)

	//Output:
	//test: ExchangeHandler() -> true

}
