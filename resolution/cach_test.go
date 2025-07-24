package resolution

import (
	"encoding/json"
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"net/http"
)

func ExampleNewCache() {
	name := "test:thing:text"

	c := newCache()
	buf, err := json.Marshal(text{Value: "Hello World!"})
	if err != nil {
		fmt.Printf("test: json.Marshall() -> [err:%v]\n", err)
	} else {

		c.put(name, messaging.Content{Fragment: "", Type: http.DetectContentType(buf), Value: buf})
		//fmt.Printf("test: newCache.put(1) -> [status:%v]\n", status)

		content, err1 := c.get(name)
		fmt.Printf("test: newCache.get(\"%v\") -> [%v] [err:%v]\n", name, content, err1)

		//name = name + "#1.2.4"
		name += "#1.2.4"
		content, err1 = c.get(name)
		fmt.Printf("test: newCache.get(\"%v\") -> [%v] [err:%v]\n", name, content, err1)

		//var v text
		//err = json.Unmarshal(buf, &v)
		//fmt.Printf("test: json.Unmarshal() -> [err:%v] [%v]\n", err, v)
	}

	//Output:
	//test: newCache.get("test:thing:text") -> [fragment:  type: text/plain; charset=utf-8 value: true] [err:<nil>]
	//test: newCache.get("test:thing:text#1.2.4") -> [fragment:  type:  value: false] [err:resolution [test:thing:text#1.2.4] not found]

}
