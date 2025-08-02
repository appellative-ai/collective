package namespace

import (
	"encoding/json"
	"fmt"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/std"
	"net/http"
)

func ExampleCreateRetrieval() {
	buf, err := createRetrieval("", nil)
	fmt.Printf("test: createRetrieval() -> [buf:%v] err:%v\n", len(buf), err)

	args := []Arg{
		{Name: "name", Value: "value"},
	}
	buf, err = createRetrieval("name", nil)
	fmt.Printf("test: createRetrieval() -> [buf:%v] err:%v\n", len(buf), err)

	buf, err = createRetrieval("name", args)
	fmt.Printf("test: createRetrieval() -> [buf:%v] err:%v\n", string(buf), err)

	//Output:
	//test: createRetrieval() -> [buf:0] err:empty name
	//test: createRetrieval() -> [buf:0] err:empty args
	//test: createRetrieval() -> [buf:{"name":"name","args":[{"name":"name","value":"value"}]}] err:<nil>

}

func ExampleCreateContent() {
	resp := &http.Response{}
	content, err := createContent(resp)
	fmt.Printf("test: createContent() -> [%v] err:%v\n", content, err)

	buf, err1 := json.Marshal(&retrieval{
		Name: "name",
		Args: []Arg{
			{"name", "value"},
		},
	})
	if err1 != nil {
		fmt.Printf("test: json.Marshal() -> [err:%v]\n", err)
	}
	h := make(http.Header)
	h.Add("Content-Type", std.ContentTypeJson)
	resp = httpx.NewResponse(http.StatusOK, h, buf)
	content, err = createContent(resp)
	fmt.Printf("test: createContent() -> [%v] err:%v\n", content, err)

	//Output:
	//test: createContent() -> [<nil>] err:<nil>
	//test: createContent() -> [fragment:  type: application/json value: true] err:<nil>
	
}
