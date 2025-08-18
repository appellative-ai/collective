package namespace

import (
	"fmt"
	"log"
	"time"
)

func ExampleCreateLink() {
	buf, err := createLink("", "", "", "", "")
	fmt.Printf("test: createLink() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createLink("name", "", "", "", "")
	fmt.Printf("test: createLink() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createLink("name", "", "", "", "author")
	fmt.Printf("test: createLink() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createLink("name", "", "thing1", "", "author")
	fmt.Printf("test: createLink() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createLink("name", "", "thing1", "thing2", "author")
	fmt.Printf("test: createLink() -> [buf:%v] err:%v\n", len(buf), err)

	//Output:
	//test: createLink() -> [buf:0] err:empty name [] or author []
	//test: createLink() -> [buf:0] err:empty name [name] or author []
	//test: createLink() -> [buf:0] err:empty thing1 [] or empty thing2 []
	//test: createLink() -> [buf:0] err:empty thing1 [thing1] or empty thing2 []
	//test: createLink() -> [buf:80] err:<nil>

}

func ExampleAddLink() {
	a := newAgent()

	status := a.addLink("name", "", "thing1", "thing2", "author")
	fmt.Printf("test: addLink() -> [status:%v]\n", status)

	a.logExchange = func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
		log.Printf("%v %v %v %v %v %v\n", start, duration, route, req, resp, timeout)
		//logx.LogEgress(nil, start, duration, route, req, resp, timeout)
	}
	status = a.addLink("name", "", "thing1", "thing2", "author")
	fmt.Printf("test: addLink() -> [status:%v]\n", status)

	a.timeout = time.Millisecond * 1000
	status = a.addLink("name", "", "thing1", "thing2", "author")
	fmt.Printf("test: addLink() -> [status:%v]\n", status)

	a.hosts = []string{"localhost:8080"}
	status = a.addLink("name", "", "thing1", "thing2", "author")
	fmt.Printf("test: addLink() -> [status:%v]\n", status)

	//Output:
	//test: addLink() -> [status:Internal Error - Post "https://invalid-host/namespace/request/link": dial tcp: lookup invalid-host: no such host]
	//test: addLink() -> [status:Internal Error - Post "https://invalid-host/namespace/request/link": dial tcp: lookup invalid-host: no such host]
	//test: addLink() -> [status:Internal Error - Post "https://invalid-host/namespace/request/link": context deadline exceeded]
	//test: addLink() -> [status:Internal Error - Post "http://localhost:8080/namespace/request/link": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.]

}
