package test

import (
	"fmt"
	"github.com/behavioral-ai/collective/content"
)

func ExampleLoadResolver() {
	r := content.Resolver
	status := loadResolver(r)
	fmt.Printf("test: loadResolver() -> [status:%v]\n", status)

	access, status1 := r.Get(ResiliencyThreshold)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", ResiliencyThreshold, status1, access)

	access, status1 = r.Get(ResiliencyInterpret)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] %v]\n", ResiliencyInterpret, status1, access)

	//Output:
	//test: loadResolver() -> [status:OK]
	//test: Get("behavioral-ai:resiliency:type/operative/threshold") -> [status:OK] [vers:  type: text/plain; charset=utf-8 content: true]
	//test: Get("behavioral-ai:resiliency:type/operative/interpret") -> [status:OK] vers:  type: text/plain; charset=utf-8 content: true]

}
