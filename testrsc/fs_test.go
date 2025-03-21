package testrsc

import (
	"fmt"
	"github.com/behavioral-ai/core/iox"
)

func ExampleReadFile() {
	name := "file:///f:/files/metrics/traffic-profile-1.json"
	bytes, status := iox.ReadFile(name)
	fmt.Printf("test: ReadFile() -> [buff:%v] [status:%v]\n", len(bytes), status)

	//Output:
	//test: ReadFile() -> [buff:1218] [status:<nil>]

}
