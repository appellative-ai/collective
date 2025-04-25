package fs

import (
	"fmt"
)

func ExampleReadFile() {
	name := "file:///f:/files/metrics/traffic-profile-1.json"
	bytes, status := ReadFile(name)
	fmt.Printf("test: ReadFile() -> [buff:%v] [status:%v]\n", len(bytes), status)

	//Output:
	//test: ReadFile() -> [buff:1218] [status:<nil>]

}
