package test

import (
	"fmt"
	"github.com/behavioral-ai/collective/content"
)

func ExampleStartup() {
	r := content.Resolver
	Startup()

	name := ResiliencyThreshold
	access, status1 := r.Get(name)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [access:%v]\n", name, status1, access)

	name = ResiliencyInterpret
	access, status1 = r.Get(name)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [access:%v]\n", name, status1, access)

	name = ProfileName
	access, status1 = r.Get(name)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, access)

	//Output:
	//test: Get("behavioral-ai:resiliency:type/operative/threshold") -> [status:OK] [access:vers:  type: text/plain; charset=utf-8 content: true]
	//test: Get("behavioral-ai:resiliency:type/operative/interpret") -> [status:OK] [access:vers:  type: text/plain; charset=utf-8 content: true]
	//test: Get("behavioral-ai:resiliency:type/domain/metrics/profile") -> [status:OK] [buf:vers:  type: text/plain; charset=utf-8 content: true]

}
