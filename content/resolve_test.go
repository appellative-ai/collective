package content

import (
	"fmt"
	"github.com/behavioral-ai/collective/testrsc"
	url2 "net/url"
)

func ExampleResolution_PutValue() {
	urn := "type/test"
	r := NewEphemeralResolver()
	url, _ := url2.Parse(testrsc.ResiliencyTrafficProfile1)

	status := r.PutValue(urn, "author", url, 1)
	fmt.Printf("test: PutValue() -> [status:%v]\n", status)

	buf, status1 := r.GetValue(urn, 1)
	fmt.Printf("test: GetValue() -> [status:%v] [%v]\n", status1, string(buf))

	//Output:
	//fail

}
