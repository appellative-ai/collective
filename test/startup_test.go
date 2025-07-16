package test

import (
	"fmt"
	"github.com/appellative-ai/collective/resource"
)

func ExampleStartup() {
	r := resource.Resolver
	Startup()

	name := ResiliencyThreshold
	access, status1 := r.Representation(name)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [access:%v]\n", name, status1, access)

	name = ResiliencyInterpret
	access, status1 = r.Representation(name)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [access:%v]\n", name, status1, access)

	name = ProfileName
	access, status1 = r.Representation(name)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, access)

	//Output:
	//test: Get("appellative-ai:resiliency:type/operative/threshold") -> [status:OK] [access:vers:  type: text/plain; charset=utf-8 resource: true]
	//test: Get("appellative-ai:resiliency:type/operative/interpret") -> [status:OK] [access:vers:  type: text/plain; charset=utf-8 resource: true]
	//test: Get("appellative-ai:resiliency:type/domain/metrics/profile") -> [status:OK] [buf:vers:  type: text/plain; charset=utf-8 resource: true]

}
