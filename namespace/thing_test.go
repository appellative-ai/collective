package namespace

import (
	"fmt"
	"github.com/appellative-ai/core/logx"
	"time"
)

func ExampleCreateThing() {
	buf, err := createThing("", "", "")
	fmt.Printf("test: createThing() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createThing("name", "", "")
	fmt.Printf("test: createThing() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createThing("", "", "author")
	fmt.Printf("test: createThing() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createThing("name", "", "author")
	fmt.Printf("test: createThing() -> [buf:%v] err:%v\n", len(buf), err)

	//Output:
	//test: createThing() -> [buf:0] err:empty name
	//test: createThing() -> [buf:0] err:empty author
	//test: createThing() -> [buf:0] err:empty name
	//test: createThing() -> [buf:68] err:<nil>

}

func ExampleAddThing() {
	a := newAgent()

	a.logFunc = func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
		logx.LogEgress(nil, start, duration, route, req, resp, timeout)
	}
	status := a.addThing("name", "", "author")
	fmt.Printf("test: addThing() -> [status:%v]\n", status)

	//Output:
	//test: addThing() -> [status:Internal Error - Post "https://invalid-host/namespace/request/thing": dial tcp: lookup invalid-host: no such host]

}
