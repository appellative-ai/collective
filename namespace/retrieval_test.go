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
	content := std.Content{Value: nil}
	resp := &http.Response{}

	ok, err := createContent(resp, &content)
	fmt.Printf("test: createRetrieval() -> [%v] [ok:%v] err:%v\n", content, ok, err)

	buf, err1 := json.Marshal(&retrieval{
		Name: "name",
		Args: []Arg{
			{"name", "value"},
		},
	})
	if err1 != nil {
		fmt.Printf("test: json.Marshal() -> [err:%v]\n", err)
	}
	content = std.Content{Value: nil}
	h := make(http.Header)
	h.Add("Content-Type", std.ContentTypeJson)
	resp = httpx.NewResponse(http.StatusOK, h, buf)
	ok, err = createContent(resp, &content)
	fmt.Printf("test: createRetrieval() -> [%v] [ok:%v] err:%v\n", content, ok, err)

	//Output:
	//test: createRetrieval() -> [fragment:  type:  value: false] [ok:false] err:<nil>
	//test: createRetrieval() -> [fragment:  type: application/json value: true] [ok:true] err:<nil>

}
