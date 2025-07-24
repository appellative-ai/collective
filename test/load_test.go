package test

import (
	"fmt"
	"github.com/appellative-ai/collective/resolution"
)

func ExampleLoadResolver() {
	r := resolution.Resolver
	status := loadResolver(r)
	fmt.Printf("test: loadResolver() -> [status:%v]\n", status)

	access, status1 := r.Representation(ResiliencyThreshold)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [%v]\n", ResiliencyThreshold, status1, access)

	access, status1 = r.Representation(ResiliencyInterpret)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] %v]\n", ResiliencyInterpret, status1, access)

	//Output:
	//test: loadResolver() -> [status:OK]
	//test: Get("appellative-ai:resiliency:type/operative/threshold") -> [status:OK] [vers:  type: text/plain; charset=utf-8 resolution: true]
	//test: Get("appellative-ai:resiliency:type/operative/interpret") -> [status:OK] vers:  type: text/plain; charset=utf-8 resolution: true]

}
